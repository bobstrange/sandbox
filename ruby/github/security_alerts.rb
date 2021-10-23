require "pry"
require "dotenv"
require "octokit"
require "json"

Dotenv.load

client = Octokit::Client.new(access_token: ENV["TOKEN"])

# query = <<~GRAPHQL
# query {
# repository(owner:"#{ENV["OWNER"]}", name:"#{ENV["REPO"]}") {
# id
# name
# description
# vulnerabilityAlerts(first: 100) {
# totalCount
# nodes {
# createdAt
# dismisser {
# id
# name
# }
# dismissedAt
# dismissReason
# repository {
# id
# nameWithOwner
# }
# id
# securityVulnerability {
# advisory {
# id
# ghsaId
# description
# summary
# severity
# updatedAt
# }
# severity
# }
# vulnerableManifestFilename
# vulnerableManifestPath
# vulnerableRequirements
# }
# }
# }
# }
# GRAPHQL
# result = client.post("/graphql", { query: query }.to_json).to_hash
# result.dig(:data, :repository, :vulnerabilityAlerts, :nodes)

result = client.get("/repos/#{ENV { "OWNER" }}/#{ENV["REPO"]}/secret-scanning/alerts")
binding.pry
