RSpec.describe "Ruby 2.5 features" do
  example "ERB#result_with_hash" do
    expect(ERB.new("<%= a + b %>").result_with_hash(a: 1, b: 2)).to eq("3")
  end

  example "Kernel#yield_self" do
    expect(2.yield_self { |x| x + 2 }.yield_self { |x| x ** x }.to_s).to eq("256")
    # Ruby 2.6 then
    expect(2.then { |x| x + 2 }.then { |x| x ** x }.to_s).to eq("256")
  end

  example "Array#prepend / Array#append" do
    expect(["c", "d"].prepend("a", "b")).to eq(["a", "b", "c", "d"])
    expect(["a", "b"].append("c", "d")).to eq(["a", "b", "c", "d"])

    expect(["c", "d"].unshift("a", "b")).to eq(["a", "b", "c", "d"])
    expect(["a", "b"].push("c", "d")).to eq(["a", "b", "c", "d"])
  end

  example "Hash#transform_keys" do
    expect({
      foo: "bar",
      baz: "qux",
    }.transform_keys(&:to_s)).to eq("foo" => "bar", "baz" => "qux")

    expect({
      "foo" => "bar",
      "baz" => "qux",
    }.transform_keys(&:to_sym)).to eq(foo: "bar", baz: "qux")
  end

  example "String#start_with?" do
    expect("Hello world".start_with?(/\w+/)).to be_truthy
  end

  example "Hash#slice" do
    expect({
      foo: "bar",
      bar: "baz",
      qux: "qux",
    }.slice(:foo, :qux)).to eq(foo: "bar", qux: "qux")
  end
end
