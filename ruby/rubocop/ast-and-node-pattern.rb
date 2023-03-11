require 'rubocop'

code = '!something.empty?'

source = RuboCop::ProcessedSource.new(code, RUBY_VERSION.to_f)
node = source.ast

puts node
puts "node.type: #{node.type}"
puts "node.children: #{node.children}"
puts "node.source: #{node.source}"

puts "\n" * 2
puts "RuboCop::NodePattern.new('send').match(node) => #{RuboCop::NodePattern.new('send').match(node)}"
puts "\n"
puts "RuboCop::NodePattern.new('(send ...)').match(node) => #{RuboCop::NodePattern.new('(send ...)').match(node)}"
puts "\n"
puts "RuboCop::NodePattern.new('(send (send ...) :!)').match(node) => #{RuboCop::NodePattern.new('(send (send ...) :!)').match(node)}"
puts "\n"
puts "RuboCop::NodePattern.new('(send (send (send ...) :empty?) :!)').match(node) => #{RuboCop::NodePattern.new('(send (send (send ...) :empty?) :!)').match(node)}"

# $... でマッチした部分のリテラル表現を取得できる
puts RuboCop::NodePattern.new('(send (send (send $...) empty?) :!)').match(node)

puts RuboCop::ProcessedSource.new('!"something".empty?', RUBY_VERSION.to_f).ast
puts RuboCop::ProcessedSource.new('![].empty?', RUBY_VERSION.to_f).ast
