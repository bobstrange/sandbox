/** Spread */

const hat = { name: 'Hat', price: 100 }
const boots = { name: 'Boots', price: 200 }

const otherHat = { ...hat }
console.log(`Other Hat name: ${otherHat.name}, price: ${otherHat.price}`)

const additionalProperties = { ...hat, discounted: false }
console.log(`Other Hat name: ${additionalProperties.name}, price: ${additionalProperties.price} discounted: ${additionalProperties.discounted}`)

const replaceProperties = { ...hat, price: 500 }
console.log(`Other Hat name: ${replaceProperties.name}, price: ${replaceProperties.price}`)

const prependReplaceProperties = { price: 500, ...hat }
console.log(`Other Hat name: ${replaceProperties.name}, price: ${replaceProperties.price}`)

const { price, ...restProperties } = hat
console.log(`Price: ${price} Rest properties ${JSON.stringify(restProperties)}`)

/** Map */

class Product {
  constructor(name, price) {
    this.name = name
    this.price = price
  }

  toString() {
    return `Name: ${this.name} Price: ${this.price}`
  }
}

const dataStore = new Map()
dataStore.set('hat', new Product('Hat', 1000))
dataStore.set('boot', new Product('Boot', 2000))

// keys()
console.log(`Map.keys() ${dataStore.keys()}`) // Map.keys() [object Map Iterator]
for (const key of dataStore.keys()) {
  console.log(`key: ${key}`)
}

// values()
console.log(`Map.values() ${dataStore.values()}`) // Map.values() [object Map Iterator]
for (const value of dataStore.values()) {
  console.log(`value: ${value}`)
}

// entries()
console.log(`Map.entries() ${dataStore.entries()}`) // Map.entries() [object Map Iterator
for (const entry of dataStore.entries()) {
  console.log(`entry: ${entry}`)
}
// key-value pair
// entry: hat,Name: Hat Price: 1000
// entry: boot,Name: Boot Price: 2000

for (const [key, value] of dataStore) {
  console.log(`key: ${key} value: ${value}`)
}
const data = new Map(
  [
    ['js', { name: 'JavaScript' }],
    ['ts', { name: 'TypeScript' }],
  ]
)
for (const [key, value] of data) {
  console.log(`key: ${key} value: ${JSON.stringify(value)}`)
}

// Map.keys()/Map.values()でキーのみ、値のみのIterator
for (const key of data.keys()) {
  console.log(`key: ${key}`)
}
for (const value of data.values()) {
  console.log(`value: ${JSON.stringify(value)}`)
}

