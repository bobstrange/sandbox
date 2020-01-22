function squareOf(n: number) {
  return n * n
}

squareOf(2)

/** Unknown */
let a: unknown = 100
let b = a === 123

if (typeof a === 'number') {
  let c = a + 10
}

/** BigInt */
let d = 1234n

/** Symbol */
let e = Symbol('e')
let f: symbol = Symbol('f')
let g = e === f

let newE = Symbol('e')
console.log(e === newE) /** false !!!*/

/** object */
let obj = {
  key1: 'x'
}

obj.key1

const something: { b: number } = { b: 12 }
const something2 = { b: 12 }

let i
let j = i * 3

/** Union and Intersection */
type Cat = { name: string, purrs: boolean }
type Dog = { name: string, barks: boolean, wags: boolean }
type CatOrDogOrBoth = Cat | Dog // Union
type CatAndDog = Cat & Dog // Intersection

let creature: CatAndDog = {
  name: 'Domito',
  barks: true,
  purrs: true,
  wags: true
}

/** null, undefined, void and never */

function returnNumberOrNull(x: number) {
  if (x > 10) {
    return x
  }
  return null
}

function returnUndefined() {
  return undefined
}

function returnVoid() {
  let a = 2 + 2
  let b = a * a
}

// Inferred as void ???
function returnNever() {
  throw TypeError('I always error')
}

// Inferred as void as well ???
function returnNever2() {
  while (true) {
    console.log('Endless loop')
  }
}

/** enum */
const enum Flippable {
    Burger = 'Burger',
    Chair = 'Chair'
}

function flip(f: Flippable) {
    return 'Flipped it'
}

flip(Flippable.Burger)
flip(Flippable.Chair)
flip(100) // Error
flip('Hat') // Error

