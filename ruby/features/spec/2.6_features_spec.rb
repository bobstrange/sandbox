RSpec.describe "Ruby 2.6 features" do
  example "Enumerable#to_h" do
    expect((1..5).to_h { |x| [x, x ** 2] }).to eq(1 => 1, 2 => 4, 3 => 9, 4 => 16, 5 => 25)
  end

  example "Enumerator#filter" do
    expect([1, 2, 3, 4, 5].filter(&:even?)).to eq([2, 4])
  end

  example "Enumerator#chain" do
    expect([1, 2, 3].chain([4, 5], (6..9)).to_a).to eq([1, 2, 3, 4, 5, 6, 7, 8, 9])
  end

  example "Array#union / difference" do
    expect(
      [1, 2, 3].union(
        [2, 4, 6],
        [3, 5, 7]
      )
    ).to eq([1, 2, 3, 4, 6, 5, 7])

    expect(
      [1, 2, 3, 4, 5].difference(
        [2, 3, 4],
        [4, 5]
      )
    ).to eq([1])
  end

  example "Hash#merge" do
    expect(
      { foo: 1, bar: 2 }.merge(
        { baz: 3, qux: 4 },
        hoge: 1,
      )
    ).to eq(
      foo: 1,
      bar: 2,
      baz: 3,
      qux: 4,
      hoge: 1,
    )
  end

  example "procs composition" do
    capitalize = :capitalize.to_proc
    add_header = ->(value) { "Title: #{value}" }

    # 右から順番に実行される
    format_as_title = add_header << capitalize << :strip.to_proc

    expect("hello world   ".then(&format_as_title)).to eq("Title: Hello world")

    # 左から順番に実行される
    expect("hello world   ".then(&(add_header >> capitalize >> :strip.to_proc))).to eq("Title: hello world")
  end
end
