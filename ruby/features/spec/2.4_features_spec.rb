def array_or_nil(is_array:)
  is_array ? ["foo", "bar"] : nil
end

RSpec.describe "Ruby 2.4 features" do
  it "multiple assignment in conditional expression" do
    foo, bar = array_or_nil(is_array: true)
    expect(foo).to eq("foo")
    expect(bar).to eq("bar")

    foo, bar = array_or_nil(is_array: false)
    expect(foo).to be_nil
    expect(bar).to be_nil

    expect(array_or_nil(is_array: true)).to be_truthy
    expect(array_or_nil(is_array: false)).to be_falsey
  end
end
