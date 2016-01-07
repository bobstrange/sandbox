require 'minitest/autorun'

class TestGrepV < Minitest::Test
  def setup
    @target = [
      "Line1: OK",
      "Line2: Invalid character",
      "Line3: Invalid number",
      "Line4: OK",
      "Test: OK"
    ]
  end

  def test_grep_v
    assert_equal ["Line1: OK", "Line4: OK", "Test: OK"], @target.grep(/OK$/)
    assert_equal ["Line2: Invalid character", "Line3: Invalid number"], @target.grep_v(/OK$/)
    assert_equal ["Test: OK"], @target.grep(/^Test/)
    assert_equal [], @target.grep_v(/:/)
  end
end
