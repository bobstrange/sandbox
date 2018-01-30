#!/usr/bin/env ruby

require 'forwardable'

class WriterDecorator
  extend Forwardable

  def_delegators :@writer, :wirte_line, :position, :rewind, :close

  def initialize(writer)
    @writer = writer
  end
end
