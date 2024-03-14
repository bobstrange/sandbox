describe('array matcher') do
  it do
    users = [
      { 'name' => 'John', 'age' => 25 },
      { 'name' => 'Jane', 'age' => 22 },
    ]

    expect(users).to contain_exactly(
      { 'name' => 'John', 'age' => 25 },
      { 'name' => 'Jane', 'age' => 22 },
    )
  end

  describe 'include は特定の要素が含まれているかどうかのみをチェックする' do
    it '特定の要素が全て含まれている' do
      users = [
        { 'name' => 'John', 'age' => 25 },
        { 'name' => 'Jane', 'age' => 22 },
      ].shuffle

      expect(users).to include(
        { 'name' => 'John', 'age' => 25 },
        { 'name' => 'Jane', 'age' => 22 },
      )
    end

    it '特定の要素が一部含まれている' do
      users = [
        { 'name' => 'John', 'age' => 25 },
        { 'name' => 'Jane', 'age' => 22 },
      ].shuffle

      expect(users).to include(
        { 'name' => 'John', 'age' => 25 },
      )
    end

    it '特定の要素が含まれていない (Fail)' do
      users = [
        { 'name' => 'John', 'age' => 25 },
        { 'name' => 'Jane', 'age' => 22 },
      ].shuffle

      expect(users).to include(
        { 'name' => 'Sam', 'age' => 41 },
      )
    end
  end
end
