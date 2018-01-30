#!/usr/bin/env ruby

require './decorator'

class NumberingWriter < Decorator
  def initialize(writer)
    super(writer)
    @line_number = 1
  end

  def write_line(line)
    @writer.write_line("#{@line_number}: #{line}")
  end
end
