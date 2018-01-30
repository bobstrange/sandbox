#!/usr/bin/env ruby

class Writer
  def initialize(path)
    @file = File.open(path, 'w')
  end

  def write_line(line)
    @file.print(line)
    @file.print("\n")
  end

  def position
    @file.pos
  end

  def rewind
    @file.rewind
  end

  def close
    @file.close
  end
end
