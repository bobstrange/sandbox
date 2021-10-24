require "pry"
require "dotenv"
require "octokit"
require "json"

Dotenv.load

client = Octokit::Client.new(access_token: ENV["TOKEN"])

def get(client, page_number, per_page)
  puts "page_number: #{page_number}"
  begin
    result = client.get(
      "/repos/#{ENV["OWNER"]}/#{ENV["REPO"]}/secret-scanning/alerts",
      header: "application/vnd.github.v3+json",
      page: page_number,
      per_page: per_page,
    )
    result
  rescue Octokit::InternalServerError => e
    puts "error page_number: #{page_number}"
    error = JSON.parse(e.response_body)
    match = error["message"].match(/Unable to fetch secret for alert number (\d+)\.$/)

    File.open("error_numbers.txt", "a") { |f|
      f.write("#{match[1]}\n")
    }
  end
end

# Remove the error_numbers.txt file if it exists
File.exist?("error_numbers.txt") && File.delete("error_numbers.txt")

per_page = 5

result = (0..(955 / per_page)).collect { |page_number|
  data = get(client, page_number, per_page)
  sleep 1
  data
}.flatten.map(&:to_h) # Result is Array<Sawyer::Resource> so we need to convert to Hash

File.write("result.json", result.to_json)
