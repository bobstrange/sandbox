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
