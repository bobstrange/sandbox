require 'aws-sdk-s3'
require './bootstrap'

resource = Aws::S3::Resource.new

buckets = resource.buckets

# Show bucket names
buckets.map(&:name)

bucket = buckets.find do |bucket|
  bucket.name == ENV['BUCKET_NAME']
end

# List object keys
keys = bucket.objects.map(&:key)
pp "All keys: #{keys}"
puts

# List objects filtered by prefix
filtered_keys = bucket.objects(prefix: ENV['EXAMPLE_PREFIX']).map(&:key)

pp "Filtered keys: #{filtered_keys}"
puts

# You could filter by yourself but it traverses all objects
filtered_keys = bucket.objects.select { |obj| obj.key =~ /^#{ENV['EXAMPLE_PREFIX']}/ }.map(&:key)
pp "Filtered keys: #{filtered_keys}"

# Fetch object
data = bucket.object(filtered_keys.first).get.body.read

pp "Data: #{data}"
