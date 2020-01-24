# Handling Errors

Most common patterns:
- Returning `null`
- Throwing exceptions
- Returning exceptions
- The `Option` type

## Returning null

We're going to write a program that asks a user for their birthday, which we will
then parse into a `Date` object.

```ts
function ask() {
  return prompt('When is your birthday?')
}

function parse(birthday: string): Date {
  return new Date(birthday)
}

let date = parse(ask())
console.info(`Date is ${date.toISOString()}`)
```

We should probably validate the date the user entered.

```ts
function parse(birthday: string): Date | null {
  let date = new Date(birthday)
  if (!isValid(date)) {
    return null
  }
  return date
}
// Checks if the given date is valid
function isValid(date: Date) {
  return Object.prototype.toString.call(date) === '[object Date]' &&
    !Number.isNaN(date.getTime())
}
```

When we consume this, we're forced to first check if the result is null before we can use it:

```ts
let date = parse(ask())
if (date) {
  console.info(`Date is ${date.toISOString()}`)
} else {
  console.error('Error parsing date for some reason')
}
```

Returning `null` is the most lightweight way to handle errors in a typesafe way.
We lose some information doing it this way `parse` doesn't tell us why the operation
failed.
Returning `null` is also difficult to compose: having to check for `null` after
every operation can become verbose as you start to nest and chain operations.

## Throwing exceptions

Let's throw an exception instead of returning `null`.

```ts
function parse(birthday: string): Date {
  let date = new Date(birthday)
  if (!isValid(date)) {
    throw new RangeError('Enter a date in the form YYYY/MM/DD')
  }
  return date
}

try {
  let date = parse(ask())
  console.info(`Date is ${date.toISOString()}`)
} catch(e) {
  console.error(e.message)
}
```

We probably want to be careful to rethrow other exceptions, so we don't silently swallow
every possible error.

```ts
try {
  let date = parse(ask())
  console.info(`Date is ${date.toISOString()}`)
} catch(e) {
  if (e instanceof RangeError) {
    console.error(e.message)
  }
  throw e
}
```

We might want to subclass the error for something more specific.

```ts
class InvalidDateFormatError extends RangeError {}
class DateIsInTheFutureError extends RangeError {}

function parse(birthday: string): Date {
  let date = new Date(birthday)
  if (!isValid(date)) {
    throw new InvalidDateFormatError('Enter a date in the form YYYY/MM/DD')
  }
  if (date.getTime() > Date.now()) {
    throw DateIsInTheFutureError('Are you a timeload?')
  }
  return date
}

try {

  let date = parse(ask())
  console.info(`Date is ${date.toISOString()}`)
} catch(e) {
  if (e instanceof InvalidDateFormatError) {
    console.error(e.message)
  } else if (e instanceof DateIsInTheFutureError) {
    console.info(e.message)
  }
  throw e
}
```

TypeScript doesn't encode exceptions as part of a function's signature.

## Returning exceptions

```ts
function parse(
  birthday: string
): Date | InvalidDateFormatError | DateIsInTheFutureError {
  let date = new Date(birthday)
  if (!isValid(date)) {
    return new InvalidDateFormatError('Enter a date in the form YYYY/MM/DD')
  }
  if (date.getTime() > Date.now()) {
    return DateIsInTheFutureError('Are you a timeload?')
  }
  return date
}
```

Now consumer is forced to handle all three cases or they'll get a `TypeError` at compile time.

```ts
let result = parse(ask())
if (result instancof InvalidDateFormatError) {
  console.error(result.message)
} else if (result instanceof DateIsInTheFutureError) {
  console.info(result.message)
} else {
  console.info(`Date is ${result.toISOString()}`)
}
```

A downside is that chaining and nesting error-giving operations can quickly get verbose.
If a function returns `T | Error1`, then any function that consumes that function
has two options:

1. Explicitly handle `Error1`
2. Handle `T` and pass `Error1` through to its consumers to handle.

## The option type

You can also describe exceptions using special-purpose data types.
THree of the most popular options ar the `Try`, `Option` and `Either` types.

The idea is that instead of returning a value, your return a `container` that may
chain operations even though there may not actually be a value inside.

The container can be pretty much any data structure, so long as it can hold a value.
For example you could use an array as the container:

```ts
function parse(birthday: string): Date[] {
  let date = new Date(birthday)
  if (!isValid(date)) {
    return []
  }
  return [date]
}

let date = parse(ask())
date.map(_ => _.toISOString()).forEach(_ => console.info(`Date is ${_}`))
```

It's hard to understand so let's wrap what we're doing.

```ts
ask()
  .flatMap(parse)
  .flatMap(date => new Some(date.toISOString()))
  .flatMap(date => new Some(`Date is ${date}`))
  .getOrElse('Error parsing date for some reason')
```

We'll define our `Option` type like this:

- `Option` is an interface that is implemented by two classes: `Some<T>` and `None`
- `Option` is both a type and a function. Its type is an interface that simply serves as the supertype of `Some` and `None`. Its function is the way to create a new value type `Option`.

```ts
interface Option<T> {
  flatMap<U>(f: (value: T) => None): None
  flatMap<U>(f: (value: T) => Option<U>): Option<U>
  getOrElse(value: T): T
}
class Some<T> implements Option<T> {
  constructor(private value: T) {}

  flatMap<U>(f: (value: T) => None): None
  flatMap<U>(f: (value: T) => Some<U>): Some<U>
  flatMap<U>(f: (value: T) => Option<U>): Option<U> {
    return f(this.value)
  }
  getOrElse(): T {
    return this.value
  }
}

class None implements Option<never> {
  flatMap(): None {
    return this
  }
  getOrElse<U>(value: U): U {
    return value
  }
}
```
