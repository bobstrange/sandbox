def list
  [1, 2, 3].shuffle
end

describe('Using not appropriate matcher') do
  it do
    expect(list.include?(1)).to be_truthy
    expect(list.include?(2)).to be_truthy
    expect(list.include?(3)).to be_truthy
  end

  it do
    expect(list).to contain_exactly(1, 2, 3)
  end
end
