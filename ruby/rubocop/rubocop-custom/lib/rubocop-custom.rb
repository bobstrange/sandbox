# frozen_string_literal: true

require 'rubocop'

require_relative 'rubocop/custom'
require_relative 'rubocop/custom/version'
require_relative 'rubocop/custom/inject'

RuboCop::Custom::Inject.defaults!

require_relative 'rubocop/cop/custom_cops'
