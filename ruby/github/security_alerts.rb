require "pry"
require "dotenv"
require "octokit"
require "json"

Dotenv.load

client = Octokit::Client.new(access_token: ENV["TOKEN"])

result = client.get("/repos/#{ENV { "OWNER" }}/#{ENV["REPO"]}/secret-scanning/alerts")
binding.pry
