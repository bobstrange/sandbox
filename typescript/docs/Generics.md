# Generics memo
## When are the generics bound

The place where you declare a generic type doesn't just scope the type, but also
dictates when TypeScript will bind a concrete type to your generic.

```
type Filter = {
    <T>(array: T[], f: (item: T) => boolean): T[]
}
let filter: Filter = (array, f) => {
    // ...
}
```

Because we declared <T> as part of a call signature, TypeScript will bind a concrete
type to `T` when we actually call a function of type `Filter`.

If we'd instead scoped `T` to the type alias `Filter`, TypeScript would have required
us to bind a type explicitly when we used `Filter`:

```
type Filter<T> = {
    (array: T[], f: (item: T) => boolean): T[]
}
let filter: Filter = (array, f) => { // Error Generic type 'Filter' requires 1 type argument
}

let filter: Filter<number> = (array, f) => { // You need to pass type argument
  // ...
}
```

## Where can you declare generics?

As T is scoped to a single signature, TypeScript will bind the T
in this signature to a concrete type when you **call** a function of type `filter`.

```
type Filter = {
    <T>(array: T[], f: (item: T) => boolean): T[]
}
type Filter = <T>(array: T[], f: (item: T) => boolean) => T[]
```

As T is declared as part of `Filter`'s type, TypeScript will bind the T
when you **declare** a function of type `Filter`

```
type Filter<T> = {
    (array: T[], f: (item: T) => boolean): T[]
}
type Filter<T> = (array: T[], f: (item: T) => boolean) => T[]
```

## Generic type inference
When you call the `map` function, TypeScript infers that `T` is `string` and `U` is `boolean`:

```
function map<T, U>(array: T[], f: (item: T) => U): U[] {
  let result = []
  for (let i = 0; i < array.length; i++) {
    result[i] = f(array[i])
  }
  return result
}
map(
    ['a', 'b', 'c'], // An array of T
    _ => _ === 'a'   // A function that returns a U
)
```

You can explicitly annotate your generics:

```
map<string, boolean>(
    ['a', 'b', 'c'],
    _ => _ === 'a'
)
```

Since TypeScript infers concrete types for your generics from the arguments you pass into
your generic function, sometimes you'll hit a case like:

```
let promise = new Promise(resolve => resolve(45))
promise.then(result => result * 4) // Error: The left-hand side of an arithmetic operation must be of type 'any', 'number', 'bigint', or an enum type.
```

TypeScript infers result as unknown
As we didn't give it enough information to work with.
To fix this, we have to explicitly annotate `Promise`s generic type parameter.

```
let promise = new Promise<number>(resolve => resolve(45))
promise.then(result => result * 4) // result will be inferred as number
```

## Generic type aliases

Define a `MyEvent` type that describes a `DOM` event, like a `click` or a `mousedown`.

```
type MyEvent<T> = {
    target: T
    type: string
}
```

Note that this is the only valid place to declare a generic type in a type alias:
right after the type alias's name, before its assignment.

When you use a generic type like `MyEvent`, you have to explicitly bind its type
parameters when you use the type; they won't be inferred for you:

```
let myEvent: MyEvent<HTMLButtonElement | null> = {
    target: document.querySelector('#myButton'),
    type: 'click'
}
```

You can use `MyEvent` to build another type --- say, `TimedEvent`.
When the generic `T` in `TimedEvent` is bound,　TypeScript will also bind it to `MyEvent`.

```
type TimedEvent<T> = {
    event: MyEvent<T>
    from: Date
    to: Date
}
```

You can use a generic type alias in a function's signature, too.

```
function triggerEvent<T>(event: MyEvent<T>) {
    // ...
}
triggerEvent({ // T is Element | null
    target: document.querySelector('#myButton'),
    type: 'mouseover'
})
