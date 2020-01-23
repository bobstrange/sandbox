# ProgrammingTypeScriptMemo

## How typechecker works
1. TypeScriptSource -> TypeScriptAST
2. AST is checked by Typechecker
3. TypeScriptAST -> JavaScript source

TypeScript is a gradually typed language.
It could compile without any type information.

## Basic setup.
tsconfig.json
.eslintrc.js
.eslintignore

npm install --save-dev typescript @typescript-eslint/parser @typescript-eslint/eslint-plugin

## noImplicitAny
To get TypeScript complain about implicit anys.
This is part of the strict family of TSC flags.

## unknown
For the few cases you have a type whose type your really don't know ahead of time, don't use any,
and instead reach for unknown.
TypeScript won't let you use an `unknown` type until you refine it by checking what it is.

## bigint
While the number type can only represent whole numbers up to 2 ** 53, bigint can represent integers bigger than that too.

## symbol
Symbols are used as an alternative to string keys in objects and maps, in places where you want to be
extra sure that people are using the right well-known key and didn't accidentally set the key.

Symbol is unique and will not be equal even if you create a second symbol with the same exact name.

## object
Typescript  favors that style of programming over a nominally typed style.
(see: [Nominal type system](https://en.wikipedia.org/wiki/Nominal_type_system)

### Structural typing
A style of programming where you just care that an object has certain properties and not what its name is.

Unlike the primitive types we've looked at sor far, declaring an object with const won't hint to TypeScript to infer its type mor narrowly.

```
const something = { key: 10 } /* TypeScript infers { key: number }, not { key: 10 }
```

### Definite Assignment

```
let i: number
let j = i * 3 // Error TS2454: Variable 'i' is used before being assigned.
```

```
let i
let j = i * 3 // Error TS2532: Object is possibly 'undefined'
```

### Index signatures

```
let a: {
  b: number
  c?: string // optional property
  [key: number]: boolean
}
// what types objects we can assign to a
a = { b: 1 }
a = { b: 1, c: undefined }
a = { b: 1, c: 'd' }
a = { b: 1, 100: true }
a = { b: 1, 50: true, 100: false }
a = { b: 1, 40: 'red' } // Error
```

The `[key: T]: U` syntax is called an index signature.
This is the way you tell TypeScript that the given object might contain more keys.

**There is one rule to keep in mind for index signatures: the index signature key's type `(T)` must be assignable to either number or string.**

You can use any word for the index signature key's name (it doesn't have to be `key`).

```
let airplaneSeatingAssignments: {
  [seatNumber: string]: string
} = {
  '34D': 'Bob',
  '34E': 'Jane'
}
```

## Type aliases
Just liike you can use variable decclarations to declare variable that aliases a value,
you can declare a type alias that points to a type.

```
type Age = number
type Person = {
  name: string
  age: Age
}
```

Type aliases are block-scoped.
Type aliases are useful for DRYing up repeated complex types, and for making it clear what a variable is used for.

## Union an intersection types

Union A or B (Logical or)
Intersection A and B (Logical conjunction)

```
type Cat = { name: string, purrs: boolean }
type Dog = { name: string, barks: boolean, wags: boolean }
type CatOrDogOrBoth = Cat | Dog // Union
type CatAndDog = Cat & Dog // Intersection
```

## Arrays
The general rule of thumb is to keep arrays homogeneous.
Try to design your programs so that every element of you array has the same type.

```
// Bad example

let d = [1, 'a']
d.map(_ => {
  if (typeof _ === 'number') { // You need to query the type of each item with typeof
    return _ * 3
  }
  return _.toUpperCase()
})
```

## Tuples
Tuples are subtype of array.
They're a special way to type arrays that have fixed lengths, where the values at each index have specific,
known types.
Unlike most other types, tuples have to be explicitly typed when you declare them.

Optional elements

```
let trainFares: [number, number?][] = [
    [3.75],
    [8.25, 7.70],
    [10.50]
]
```

Rest elements

```
let friends: [string, ...string[]] = ['Sara', 'Tali', 'Chloe', 'Claire']
let friends: [number, boolean, ...string[]] = [1, false, 'a', 'b', 'c']
```

## Read-only arrays and tuples

```
let as: readonly number[] = [1, 2, 3]
let bs: readonly number[] = as.concat(4)
```

https://www.npmjs.com/package/immutable

## null, undefined, void and never
`undefined` means that something hasn't been defined yet.
`null` means an absence of value. (like if you tried to compute a value, but ran into an error along the way)

`unknown` is the subtype of every other type, then `never` is the subtype of every other type.

## Enums
A way to enumerate the possible values for a type.
They are unordered data structures that map keys to values.

TypeScript lets you access enums both by value and by key for convenience, but this can
get unsafe quickly:

```
enum Color {
    Red = '#c10000',
    Blue = '#007ac1',
    White = 255
}

let a = Color.Red
let b = Color.Green // Error Property 'Green' does not exist
let c = Color[0] // string
let d = Color[6] // string (!!! Unsafe behavior)
```

We can ask TypeScript to prevent this kind of unsafe access by opting into a safer subset of enum
behavior with const enum instead.

```
const enum Color {
    ...
}

let a = Color.Red
let b = Color.Green // Error
let c = Color[0] // Error A const enum member can only be accessed using a string literal
```

We should use string-valued enums:

```
const enum Flippable {
    Burger,
    Chair
}

function flip(f: Flippable) {
    return 'Flipped it'
}

flip(Flippable.Burger)
flip(Flippable.Chair)
flip(100) // 'Flipped it' (!!! Unexpected behavior)
```

Use string-valued enums:

```
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
```

## Optional and Default Parameters

When declaring your function's parameters, required parameters have to come first.

```
function log(message: string, userId?: string) {
    let time = new Date().toLocaleTimeString()
    console.log(time, message, userId || 'Not signed in')
}
log('Page loaded')
log('User signed in', 'da763be')
```

You can provide default values for optional parameters.

```
function log(message: string, userId = 'Not signed in') {
    let time = new Date().toLocaleTimeString()
    console.log(time, message, userId)
}
log('Page loaded')
log('User signed in', 'da763be')
```

## Rest parameters

```
function sumVariadic(...numbers: number[]): number {
    return numbers.reduce((total, n) => total + n, 0)
}
```

## call, apply and bind

In addition to invoking a function with parentheses, JavaScript supports at least two other ways to call a function.

```
function add(a: number, b: number): number {
    return a + b
}

add(10, 20)
add.apply(null, [10, 20])
add.call(null, 10, 20)
add.bind(null, [10, 20])
```

`apply` binds a value to `this` within your function, and spreads its second argument over your
function's parameters.
`call` does the same, but applies its arguments in order instead of spreading.
`bind` is similar, in that it binds `this` argument and a list of arguments to your function.
The difference is that bind odes not invoke your function.

## Generator Functions
Generator functions are a convenient way to  generate a bunch of values.
Because they're lazy, they only compute the next value when a consumer asks for it, they can do things
that can be hard to do otherwise, like generate infinite lists.

```
function* createFibonacciGenerator() {
    let a = 0
    let b = 1
    while (true) {
        yield a
        [a, b] = [b, a + b]
    }
}

let fibonacciGenerator = createFibonacciGenerator()
fibonacciGenerator.next() // { value: 0, done: false }
fibonacciGenerator.next() // { value: 1, done: false }
fibonacciGenerator.next() // { value: 1, done: false }
fibonacciGenerator.next() // { value: 2, done: false }
fibonacciGenerator.next() // { value: 3, done: false }
fibonacciGenerator.next() // { value: 5, done: false }
```

The asterisk before a function's name makes that function a generator.
Calling a generator returns an iterable iterator.

```
function* createNumbers(): IterableIterator<number> {
    let n = 0
    while (true) {
        yield n
        n += 1
    }
}
let numbersGenerator = createNumbers()
numbersGenerator.next()
numbersGenerator.next()
numbersGenerator.next()
```

## Iterators
Iterators are flip side to generators: while generators a way to produce a stream of values,
iterators are a way to consume those values.

Iterable:
Any object that contains a property called `SYmbol.iterator`, whose value is a function that
returns an iterator.

Iterator:
Any object that defines a method called next, which returns an object with the
properties `value` and `done`.

```
let numbers = {
    *[Symbol.iterator]() {
        for (let n = 1; n <= 10; n++) {
            yield n
        }
    }
}
```

## Call signatures

```
function add(a: number, b: number): number {
    return a + b
}
```

What is the type of `sum` ? -> `Function`
How else can we type `sum` ? -> `(a: number, b: number) => number`

This is TypeScript's syntax for a function's type, or `call signature`(also called a `type signature`)

```
type Greet = (name: string) => string
type Log = (message: string, userId?: string) => void
type SumVariadic = (...numbers: number[]) => number
```

The function's call signatures look remarkably similar to their implementations.

```
type Log = (message: string, userId?: string) => void
let log: Log = (
    message, // We don't need to annotate argument
    userId = 'Not signed in' // We need to pass default value as it is "value" nor "type"
) => { // We don't need to annotate return value
    let time = new Date().toISOString()
    console.log(time, message, userId)
}
```

### Type Level and Value Level Code
People use the terms "type-level" and "value-level" a lot when talking about programming
with static types, and it helps to have a common vocabulary.

## Contextual Typing

```
function times(
    f: (index: number) => void,
    n: number
) {
    for (let i = 0; i < n; i++) {
        f(i)
    }
}
```

## Overloaded Function Types

```
// Shorthand call signature
type Log = (message: string, userId?: string) => void

// Full call signature
type Log = {
    (message: string, userId?: string): void
}
```

For simple cases like our `Log` function, you should prefer the shorthand, but
for more complicated functions, there are a few good use cases for full signatures.

The first of these if `overloading` a function type.

Overloaded function:
A function with multiple call signatures.

```
type Reserve = {
    (from: Date, to: Date, destination: string): Reservation
}
let reserve: Reserve = (from, to, destination) => {
    // ...
}
```

We might re-purpose our API to support one-way trips too:

```
type Reserve = {
    (from: Date, to: Date, destination: string): Reservation
    (from: Date, destination: string): Reservation
}
```

Overloads come up naturally in browser DOM APIs.

```
type CreateElement = {
    (tag: 'a'): HTMLAnchorElement
    (tag: 'canvas'): HTMLCanvasElement
    (tag: 'table'): HTMLTableElement
    (tag: string): HTMLElement // catchall case
}
```

```
function warnUser(warning) {
    if (warnUser.wasCalled) {
        return
    }
    warnUser.wasCalled = true
    alert(warning)
}
warnUser.wasCalled = false

type WarnUser = {
    (warning: string): void
    wasCalled: boolean
}
```

## Polymorphism
Sometimes you don't know what type to expect beforehand, and you don't want to
restrict your function's behavior to a specific type.

```
function filter(array, f) {
    let result = []
    for (let i = 0; i < array.length; i++) {
        let item = array[i]
        if (f(item)) {
            result.push(item)
        }
    }
    return result
}

filter([1, 2, 3, 4], _ => _ < 3) // evaluates to [1, 2]

type Filter = {
    (array: number[], f: (item: number) => boolean): number[]
    (array: string[], f: (item: string) => boolean): string[]
    (array: object[], f: (item: object) => boolean): object[]
}
```

Rewrite it with a generic type parameter T:

```
type Filter = {
    <T>(array: T[], f: (item: T) => boolean): T[]
}
```
### Generic type parameter
A placeholder type used to enforce a type-level constraint in multiple places.
Also known as polymorphic type parameter.

T is just a type name. By convention, people use uppercase single-letter names starting
with the letter T and continuing U, V, W and so on depending on how many generics they need.
