require 'minitest/autorun'

class TestHashToProc < Minitest::Test
  def setup
    @hash = {
      name: "Bob",
      hobbies: [ "cycling", "programming", "running" ]
    }
  end

  def test_basic
    keys = [:name, :hobbies]
    assert_equal ["Bob", [ "cycling", "programming", "running" ]], keys.map(&@hash)

    keys.map do |k|
      @hash[k]
    end
  end
end
