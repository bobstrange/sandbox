def list
  [1, 2, 3].shuffle
end

describe('Using not appropriate matcher') do
  it do
    expect(list.sort).to eq([1, 2, 3])
  end

  it do
    expect(list).to contain_exactly(1, 2, 3)
  end
end
