# frozen_string_literal: true

require_relative "../simplify_not_empty_with_any"

RSpec.describe RuboCop::Cop::Style::SimplifyNotEmptyWithAny do
  subject(:cop) { described_class.new(config) }

  let(:config) { RuboCop::Config.new }

  # For example
  it 'registers an offense when using  !a.empty?' do
    expect_offense(<<~RUBY)
      !array.empty?
    RUBY
  end

  it 'does not register an offense when using `.any?` or `.empty?' do
    expect_no_offenses(<<~RUBY)
      array.any?
      array.empty?
    RUBY
  end
end
