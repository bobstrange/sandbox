# js

[`nodemon`](https://github.com/remy/nodemon)

Monitor for any changes in yoru node.js application and automatically restart
the server.

- Spread operators
- private properties 
    - `_` prefix
    - (proposal) `#` prefix
- Map

## Map
Objectは文字列のみしかキーにできないという制約がある。
キーと値での組み合わせでデータを表現する`Map`がある。

```js
const data = new Map(
  [
    ['js', { name: 'JavaScript' }],
    ['ts', { name: 'TypeScript' }],
  ]
)
// Map自身もIteratable
for (const [key, value] of data) {
  console.log(`key: ${key} value: ${JSON.stringify(value)}`)
}

// Map.entries()はMapのDefaultのIterator
for (const [key, value] of data.entries()) {
  console.log(`key: ${key} value: ${JSON.stringify(value)}`)
}

// Map.keys()/Map.values()でキーのみ、値のみのIterator
for (const key of data.keys()) {
  console.log(`key: ${key}`)
}
for (const value of data.values()) {
  console.log(`value: ${JSON.stringify(value)}`)
}
```

Mapはどんな値もキーとして使うことができる。
特に、`Symbol`は、UniqueでImmutableなので、キーとして適切。

