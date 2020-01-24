# Ex
## Which parts of a function's type signature does TypeScript infer: the parameters the return type, or both ?
both

## Is JavaScript's arguments object typesafe ? If not , what can you use instead ?
No.
Use rest parameters.

```ts
// Use this
function sum(...numbers: number[]): number {
    return numbers.reduce((total, n) => total + n, 0)
}
// Don't use this
function sum(): number {
    return Array.from(arguments).reduce((total, n) => total + n, 0) // (parameter) n: any
}
```

## You want the ability to book a vacation that starts immediately. Update the overloaded
`reserve` function with a third call signature that takes just a destination, without an explicit start date.
Update `reserve`'s implementation to support this new overloaded signature.

```ts
type Reserve = {
  (from: Date, to: Date, destination: string): Reservation
  (from: Date, destination: string): Reservation
  (destination: string): Reservation
}

type Reservation = {
  destination: string
  from?: Date
  to?: Date
}
let reserve: Reserve = (
  fromOrDestination: Date | string,
  toOrDestination?: Date | string,
  destination?: string
): Reservation {
  if (typeof fromOrDestination === 'string') {
    return { destination: fromOrDestination }
  } else if (typeof toOrDestination == 'string') {
    return { destination: toOrDestination, from: fromOrDestination }
  } else {
    return { destination, from: fromOrDestination, to: toOrDestination }
  }
}
```

## Update our call implementation to only work for functions whose second argument is  a string.

```ts
function call(
    f: (...args: unknown[]) => unknown,
    ...args: unknown[]
): unknown {
    return f(...args)
}
```

## Implement a small typesafe assertion library, `is`.

```
is('string', 'otherstring') // false
is(true, false) // false
is(42, 42) // true
is(10, 'foo') // Comparing two different types should give a compile-time error
is([1], [1, 2], [1, 2, 3]) // false
```
