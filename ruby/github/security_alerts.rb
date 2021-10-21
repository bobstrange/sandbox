require "dotenv"
require "octokit"

Dotenv.load

client = Octokit::Client.new(access_token: ENV["TOKEN"])
