#!/usr/bin/env ruby

require './numbering_writer'
require './writer.rb'

writer = NumberingWriter.new(Writer.new('sample_2.txt'))
writer.write_line('Sample line')
writer.close

