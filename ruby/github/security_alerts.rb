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
    match = error["message"] =~ /Unable to fetch secret for alert number (\d+)\.$/

    File.open("error_numbers", "a") { |f|
      f.write("#{alert_number = match[1]}\n")
    }
  end
end

per_page = 5

result = (0..(955 / per_page)).collect do |page_number|
  get(client, page_number, per_page)
  sleep 1
end

binding.pry

File.write("result.json", result.to_json)
