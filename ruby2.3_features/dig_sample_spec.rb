describe '#dig' do
  subject(:response) {
    {
      code: 200,
      data: {
        users: [
          { name: 'Foo Bar',   age: 20, hobbies: ['running', 'swimming'], tel: '08011112222' },
          { name: 'Hoge Fuga', age: 25, hobbies: []                     , tel: '09033334444' },
          { name: 'Piyo',      age: 30                                                       }
        ]
      }
    }
  }

  it { expect(response.dig(:data, :users, 0, :name)).to eq 'Foo Bar' }
  it { expect(response.dig(:data, :users, 0, :hobbies, 3)).to be_nil }
  it { expect(response.dig(:data, :users, 1, :hobbies)).to eq [] }
  it { expect(response.dig(:data, :users, 2, :tel)).to be_nil }

end
