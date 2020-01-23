/** Generator */
function* createFibonacciGenerator() {
    let a = 0
    let b = 1
    while (true) {
        yield a

        let tmp
        tmp = a
        a = b
        b = tmp + b
    }
}

let fibonacciGenerator = createFibonacciGenerator()
console.log(fibonacciGenerator.next())
console.log(fibonacciGenerator.next())
console.log(fibonacciGenerator.next())
console.log(fibonacciGenerator.next())
console.log(fibonacciGenerator.next())
console.log(fibonacciGenerator.next())

function* createNumbers(): IterableIterator<number> {
    let n = 0
    while (true) {
        yield n
        n += 1
    }
}
let numbersGenerator = createNumbers()
console.log(numbersGenerator.next())
console.log(numbersGenerator.next())
console.log(numbersGenerator.next())

/** Iterator */
let numbers = {
    *[Symbol.iterator]() {
        for (let n = 1; n <= 10; n++) {
            yield n
        }
    }
}
let iterator = numbers[Symbol.iterator]()
console.log("--- iterator ---")
console.log(iterator.next()) // 1
console.log(iterator.next()) // 2
console.log("--- iterator end ---")

console.log("--- for of start ---")
for (let a of numbers) {
  console.log(`a: ${a}`)
}
console.log("--- for of end ---")

let allNumbers = [...numbers]

console.log('spread')
console.log(allNumbers)

/** Contextual typing */
function times(
    f: (index: number) => void,
    n: number
) {
    for (let i = 0; i < n; i++) {
        f(i)
    }
}
times(n => console.log(n), 4) // You don't have to explicitly annotate the function you pass to `times`
function f(n) { // no implicit any (you need to declare f inline)
  console.log(n)
}
times(f, 4)

/** Overloaded Function Types */
type Reservation = {
  from: Date,
  destination: string
  to?: Date,
}

type Reserve = {
    (from: Date, to: Date, destination: string): Reservation
    (from: Date, destination: string): Reservation
}

/** You'll get error in the following implementation */
/* let reserve: Reserve = (from, to, destination) => {
    // ...
} */

let reserve: Reserve = (
  from: Date,
  toOrDestination: Date | string,
  destination?: string
): Reservation => {
  if (toOrDestination instanceof Date && destination !== undefined) {
    // Book a one-way trip
    return { from, destination, to: toOrDestination }
  } else {
    // Book a round trip
    return { from, destination: toOrDestination}
  }
}

/** Polymorphism */
// type Filter = {
//     (array: number[], f: (item: number) => boolean): number[]
//     (array: string[], f: (item: string) => boolean): string[]
//     (array: object[], f: (item: object) => boolean): object[]
// }

type Filter = {
    <T>(array: T[], f: (item: T) => boolean): T[]
}


let filter: Filter = (array, f) => {
    let result = []
    for (let i = 0; i < array.length; i++) {
        let item = array[i]
        if (f(item)) {
            result.push(item)
        }
    }
    return result
}

filter([1, 2, 3], _ => _ > 2) // T is bound to number
filter(['a', 'b'], _ => _ !== 'b') // T is bound to string

let names = [
  { firstName: 'beth' },
  { firstName: 'jane' },
  { firstName: 'xin' }
]
filter(names, _ => _.firstName.startsWith('b')) // T is bound to { firstName: string }
