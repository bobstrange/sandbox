# Ruby で Value Object を実装する

参考 [Value Object Semantics in Ruby](https://thoughtbot.com/blog/value-object-semantics-in-ruby#semantics)

2 つの Value Object が同じ値を表すということ

- `#==` での比較が `true` になる
- `#equal?` は、オブジェクトの同一性の比較なので `false` になる
- それぞれの `#hash` が同じ値になる
- `#eql?` での比較が `true` になる
