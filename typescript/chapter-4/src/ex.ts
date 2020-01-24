// 3.
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

// 4.

// function call(
//     f: (...args: unknown[]) => unknown,
//     ...args: unknown[]
// ): unknown {
//     return f(...args)
// }

// function call(
//     f: (arg: string, ...args: unknown[]) => unknown,
//     arg: string,
//     ...args: unknown[]
// ): unknown {
//     return f(arg, ...args)
// }
// call(_ => _, 0, 1, 2, 3) // Should fail on compile time
// call(_ => _, 'foo', 1, 2, 3)
// call(_ => _, 'foo')

function call<T extends [unknown, string, ...unknown[]], R>(
  f: (...args: T) => R,
  ...args: T
): R {
  return f(...args)
}

function fill(length: number, value: string): string[] {
  return Array.from({length}, () => value)
}

call(fill, 10, 'something')

// 5.

// function is<T>(a: T, b: T): boolean {
//   return a === b
// }

// function is<T>(a: T, b: T, ...args: T[]): boolean {
//   if (!args) {
//     return a === b
//   }
//   return a === b && args.every(_ => _ === a)
// }

// Refactor b can be concate with args...
function is<T>(a: T, ...args: [T, ...T[]]): boolean {
  return args.every(_ => _ === a)
}

is('string', 'otherstring')
is(true, false)
is(42, 42)
is(10, 'foo')
is([1], [1, 2], [1, 2, 3])
