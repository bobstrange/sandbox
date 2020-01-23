/** Generic type bound */
type Filter<T> = {
  (array: T[], f: (item: T) => boolean): T[]
}

/**
let filter: Filter = (array, f) => {  // Error Generic type 'Filter' requires 1 type argument

}
*/

/**
let filter: Filter<number> = (array, f) => { // You need to pass type argument
  // ...
}
*/

/** Where can you declare generics */
/**
function map(array: unknown[], f: (item: unknown) => unknown): unknown[] {
  let result = []
  for (let i = 0; i < array.length; i++) {
    result[i] = f(array[i])
  }
  return result
}
*/

function map<T, U>(array: T[], f: (item: T) => U): U[] {
  let result = []
  for (let i = 0; i < array.length; i++) {
    result[i] = f(array[i])
  }
  return result
}

/** Generic type inference */
/**
let promise = new Promise(resolve => resolve(45))
promise.then(result => result * 4) // Error: The left-hand side of an arithmetic operation must be of type 'any', 'number', 'bigint', or an enum type.
// TypeScript infers result as unknown
*/

let promise = new Promise<number>(resolve => resolve(45))
promise.then(result => result * 4) // result will be inferred as number


