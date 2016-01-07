require 'minitest/autorun'

class TestDig < Minitest::Test
  def setup
    @nested_hash = {
      data: {
        name: "bob",
        address: {
          pref: "Tokyo",
          city: "Setagaya",
          street: nil,
        }
      },
      success: true,
    }
    @nested_array = [
      ["Bob", "Programmer"],
      ["Maru", "Ph.D"]
    ]

    @data = [
      { name: "bob", hobbies: [ "programming", "cycling", "running"] },
      { name: "Maru", hobbies: [ "programming", "trecking"] },
      nil
    ]
  end

  def test_basic_for_hash
    assert_equal true, @nested_hash.dig(:success)
    assert_equal "bob", @nested_hash.dig(:data, :name)
    assert_equal "Tokyo", @nested_hash.dig(:data, :address, :pref)
  end

  def test_return_nil_for_hash
    # 存在しないkeyを指定した場合
    assert_equal nil, @nested_hash.dig(:hoge)
    assert_equal nil, @nested_hash.dig(:hoge, :fuga)
    assert_equal nil, @nested_hash.dig(:data, :address, :street, :hoge)
  end

  def test_raise_for_hash
    assert_raises { @nested_hash.dig(:data, :name, :hoge)}
  end

  def test_basic_for_array
    assert_equal ["Maru", "Ph.D"], @nested_array.dig(1)
    assert_equal "Ph.D", @nested_array.dig(1, 1)
  end

  def test_nil_for_array
    assert_equal nil, @nested_array.dig(3)
    assert_equal nil, @nested_array.dig(3, 3)
    assert_equal nil, @nested_array.dig(3, 3, 3)
    assert_equal nil, @nested_array.dig(2, 3)
    assert_equal nil, @nested_array.dig(2, :hoge)
  end

  def test_raise_for_array
    assert_raises { @nested_array.dig(0, 0, 3) }
  end

  def test_basic
    assert_equal "bob", @data.dig(0, :name)
    assert_equal "trecking", @data.dig(1, :hobbies, 1)
  end

  def test_return_nil
    # 存在しないkey
    assert_equal nil, @data.dig(3)
    # 存在しないkeyの戻り値がhashの場合に、それに対して存在しないkeyを指定した場合
    assert_equal nil, @data.dig(0, :hoge)
    assert_equal nil, @data.dig(0, 0)
    # 存在しないkeyの戻り値に対して、keyを指定した場合
    assert_equal nil, @data.dig(0, :hoge, :fuga)
    assert_equal nil, @data.dig(0, :hoge, 0)
    assert_equal nil, @data.dig(3, :hoge)

  end
end
