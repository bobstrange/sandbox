#!/usr/bin/env ruby

class Decorator
  def initialize(writer)
    @writer = writer
  end

  def write_line(line)
    @writer.write_line(line)
  end

  def position
    @writer.position
  end

  def rewind
    @writer.rewind
  end

  def close
    @writer.close
  end
end
