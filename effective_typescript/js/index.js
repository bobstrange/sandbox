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

