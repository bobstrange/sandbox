#!/usr/bin/env ruby

require './writer'

writer = Writer.new('sample.txt')
writer.write_line('Sample text')
writer.close
