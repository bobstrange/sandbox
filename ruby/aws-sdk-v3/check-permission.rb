require 'bundler/setup'
require 'aws-sdk-s3'

client = Aws::S3::Client.new
result = client.list_buckets
puts result
