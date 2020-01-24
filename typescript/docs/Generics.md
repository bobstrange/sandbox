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
```

## Bounded Polymorphism
Let's say we're implementing a binary tree, and have three types of nodes:

1. Regular Tree Nodes
2. LeafNodes, which are TreeNodes that don't have children.
3. InnerNodes, which are TreeNodes that do have children.

```
type TreeNode = {
    value: string
}

type LeafNode = TreeNode & {
    isLeaf: true
}

type InnerNode = TreeNode & {
    children: [TreeNode] | [TreeNode, TreeNode]
}
```

Let's write a mapNode function that takes a TreeNode and maps over its value, returning
a new TreeNode.

```
let a: TreeNode = { value: 'a' }
let b: LeafNode = { value: 'b', isLeaf: true }
let c: InnerNode = { value: 'c', children: [b] }

let a1 = mapNode(a, _ => _.toUpperCase())
let b1 = mapNode(b, _ => _.toUpperCase())
let c1 = mapNode(c, _ => _.toUpperCase())

function mapNode<T extends TreeNode>( // T can be either a TreeNode or a subtype of TreeNode
    node: T,
    f: (value: string) => string
): T {
    return {
        ...node,
        value: f(node.value)
    }
}
```

If we had typed T as just T(leaving off `extends TreeNode`), then `mapNode` would
have thrown a compile-time error.

If we had left off the T entirely and declared mapNode as (node: TreeNode, f: (value: string) => string) => TreeNode, then we would have lost information after mapping a node: a1, b1, and c1 would all just be TreeNode.

### Bounded polymorphism with multiple constraints

```
type HasSides = { numberOfSides: number }
type SidesHaveLength = { sideLength: number }

function logPerimeter<
    Shape extends HasSIdes & SidesHaveLength
>(s: Shape): Shape {
    console.log(s.numberOfSides * s.sideLength)
    return s
}

type Square = HasSides & SidesHaveLength
let square = {
    numberOfSides: 4,
    sideLength: 3
}

logPerimeter(square)
```

### Using bounded polymorphism to model arity

```
function call(
    f: (...args: unknown[]) => unknown,
    ...args: unknown[]
): unknown {
    return f(...args)
}

function fill(length: number, value: string): string[] {
    return Array.from({ length }, () => value)
}

call(fill, 10, 'a')
```

Let's fill the unknowns. The constraints we want to express are:

- f should be a function that takes some set of arguments T, and returns some type R.
We don't know how many arguments it'll have ahead of time.
- call takes f, along with the same set of arguments T that f itself takes.
- call returns the same type R that f returns

```
function call<T extends unknown[]>(
    f: (...args: T) => R,
    ...args: T
): R {
    return f(...args)
}
```

## Generic Type Defaults

Just like you can give function parameters default values, you ca give generic type parameters
default types.

```
type MyEvent<T = HTMLElement> = {
    target: T
    type: string
}
```

We can also add a bound to T, to make sure T is an HTML element.

```
type MyEvent<T extends HTMLElement = HTMLElement> = {
    target: T
    type: string
}
```

## Type driven development

Type-driven development:
A style of programming where your sketch out type signatures first, and fill in values later.

```
function map<T, U>(array: T[], f: (item: T) => U): U[] {
    // ...
}
```

Just looking at that signature you should have some intuition for what map does !!!
