require 'minitest/autorun'

class TestSafeNavigationOperator < Minitest::Test
  Book = Struct.new(:title, :author)
  Author = Struct.new(:name)

  def test_basic
    author = Author.new("Taro")
    book = Book.new("Awesome book", author)

    assert_equal author, book&.author
    assert_equal "Taro", book&.author&.name
  end

  def test_nil
    book = Book.new("Awesome book", nil)

    assert_equal nil, nil&.hoge
    assert_equal nil, book&.author&.name
  end

  def test_exception
    assert_raises { book&.hoge }
    assert_raises { false&.hoge }
    assert_raises { ""&.hoge }
  end

end
