def array_or_nil(is_array:)
  is_array ? ["foo", "bar"] : nil
end

RSpec.describe "Ruby 2.4 features" do
  example "multiple assignment in conditional expression" do
    foo, bar = array_or_nil(is_array: true)
    expect(foo).to eq("foo")
    expect(bar).to eq("bar")

    foo, bar = array_or_nil(is_array: false)
    expect(foo).to be_nil
    expect(bar).to be_nil

    expect(array_or_nil(is_array: true)).to be_truthy
    expect(array_or_nil(is_array: false)).to be_falsey
  end

  example "#clamp" do
    expect(1.clamp(0, 2)).to eq(1)
    expect(1.clamp(2, 3)).to eq(2)
    expect(1.clamp(-1, 0)).to eq(0)

    expect("b".clamp("a", "c")).to eq("b")
    expect("b".clamp("c", "d")).to eq("c")
    expect("d".clamp("a", "b")).to eq("b")
  end

  example "#digits" do
    expect(987.digits).to eq([7, 8, 9])
    expect(789.digits).to eq([9, 8, 7])
  end

  example "Regexp#match?" do
    expect(/\d{4}-\d{2}-\d{2}/.match?("2022-06-01")).to be_truthy
  end

  example "MatchData#named_captures" do
    match_data = /(?<year>\d{4})-(?<month>\d{2})-(?<day>\d{2})/.match("2022-06-01")
    expect(match_data.named_captures).to eq("year" => "2022", "month" => "06", "day" => "01")
  end
end
