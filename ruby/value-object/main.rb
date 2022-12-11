# frozen_string_literal: true

class Money
  attr_reader :amount

  def initialize(amount)
    @amount = amount
  end

  # @param [Money] value
  def ==(other)
    amount == other.amount
  end

  def hash
    [self.class, amount].hash
  end
end
