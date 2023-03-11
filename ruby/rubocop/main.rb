require 'rubocop'

code = '!something.empty?'

source = RuboCop::ProcessedSource.new(code, RUBY_VERSION.to_f)
node = source.ast

puts node
puts "node.type: #{node.type}"
puts "node.children: #{node.children}"
puts "node.source: #{node.source}"
