# Chapter5 Classes and interfaces memo
## Classes an inheritance
Chess game

```ts
class Game {}

class Piece {}

class Position {}
```

There are six types of pieces:

```ts
class King extends Piece {}
class Queen extends Piece {}
class Bishop extends Piece {}
class Knight extends Piece {}
class Rook extends Piece {}
class Pawn extends Piece {}
```

Every piece has a color and a current position.

```ts
type Color = 'Black' | 'White'

/* This will let us squeeze out some extra safety by constraining these types' domains from all strings and all numbers to a handful of very specific strings and numbers. */
type File = 'A' | 'B' | 'C' | 'D' | 'E' | 'F' | 'G' | 'H'
type Rank = 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8

class Position {
  constructor(
    /* The `private` access modifier in the constructor automatically assigns the parameter to `this` */
    private file: File,
    private rank: Rank
  ) {}
}

class Piece {
  protected position: Position
  constructor(
    private readonly color: Color,
    file: File,
    rank: Rank
  ) {
    this.position = new Position(file, rank)
  }
}
```

We've defined a `Piece` class, but we don't want users to instantiate a new `Piece` directly.
We can use the type system to enforce that for us, using the `abstract` keyword.

```ts
abstract class Piece {
  // ...
  // You can define some methods on it.

  moveTo(position: Position) {
    this.position = position
  }
  abstract canMoveTo(position: position): boolean
}
```

If a class extends `Piece` but forgets to implement the abstract `canMoveTo`  method,
that's a type error at compile time.

Implement `King`

```ts
class Position {
  // ...
  distanceFrom(position: Position) {
    return {
      rank: Math.abs(position.rank - this.rank),
      file: Math.abs(position.file.charCodeAt(0) - this.file.charCodeAt(0))
    }
  }
}

class King extends Piece {
  canMoveTo(position: Position) {
    let distance = this.position.distanceFrom(position)
    return distance.rank < 2 && distance.file < 2
  }
}
```

When we make a new game, we'll automatically create a board and some pieces:

```ts
class Game {
  private pieces = Game.makePieces()

  private static makePieces() {
    return [
      // Kings
      new King('White', 'E', 1),
      new King('Black', 'E', 8),

      // Queens
      // ...
    ]
  }
}
```

If we had entered another letter (like 'J') or an out-of-range number (like 9), TypeScript
 would have given us a compile time error.

```ts
new King('White', 'J', 1) // Error
new King('White', 'E', 9) // Error
```

## Super
The child instance can make a `super` call to its parent's version of the method.
There are two kinds of super calls:

- Method calls, like `super.take`
- Constructor calls, which have the special form `super()` and can only be called from
a constructor function. If your child class has a constructor function, you must
call `super()` from the child's constructor to correctly wire up the class.

## Using this as a return type

Just like you can use `this` as a value, you can also use it as a type.
When working with classes, the `this` type can be useful for annotating methods'
return types.

```ts
let set = new Set
set.add(1).add(2).add(3)
set.has(2) // true
set.has(4) // false

class Set {
  has(value: number): boolean {
    // ...
  }
  add(value: number): Set {
    // ...
  }
}
```

What happens when we try to subclass `Set` ?

```ts
class MutableSet extends Set {
  delete(value: number): boolean {
    // ...
  }
  // We'll need to override with MutableSet...
  add(value: number): MutableSet {
    // ...
  }
}
```

Instead you can use `this` as a return type annotation.

```ts
class Set {
  add(value: number): this {
    // ...
  }
}

// You can remove the add override from MutableSet
class MutableSet extends Set {
  // ...
}
```

**This is a really convenient feature for working with chained AP(I)S, like we do in
"Builder Pattern"**

## Interfaces

Like type aliases, interfaces are a way to name a type so you don't have to define it inline.
Type aliases an interfaces are mostly two syntaxes for the same thing, but there are
 a few small differences.

```ts
type Sushi = {
  calories: number
  salty: boolean
  tasty: boolean
}

interface Sushi {
  calories: number
  salty: boolean
  tasty: boolean
}
```

```ts
type Food = {
  calories: number
  tasty: boolean
}

type Sushi = Food & {
  salty: boolean
}

type Cake = Food & {
  sweet: boolean
}
```

You can do this interface too.

```ts
interface Food {
  calories: number
  tasty: boolean
}

interface Sushi extends Food {
  salty: boolean
}

interface Cake extends Food {
  sweet: boolean
}
```

There are three differences between types and interfaces.

The first is that type aliases are more general, in that their righthand side can be
any type expression.

The second is that when you extend an interface, TypeScript will make sure that
the interface you're extending is assignable to your extension.

```ts
interface A {
  good(x: number): string
  bad(x: number): string
}

interface B extends A {
  good(x: string | number): string // Error Type 'number' is not assignable to type 'string'
  bad(x: string): string
}
```

The third is that multiple interfaces with the same name in the same scope are
automatically merged; multiple type aliases with the same name in the same scope
will throw a compile-time error.

## Declaration merging

```ts
interface User {
  name: string
}

// User now has two fields, name and age
interface User {
  age: number
}

let user: User = {
  name: 'Bob',
  age: 36
}
```

## Implementations
When you declare a class, you can use the `implements` keyword to say that it satisfies a
particular interface.

```ts
interface Animal {
  eat(food: string): void
  sleep(hours: number): void
}

class Cat implements Animal {
  eat(food: string) {
    console.info(`Ate some ${food}. Mmm!`)
  }
  sleep(hours: number) {
    console.info(`Slept for ${hours} hours`)
  }
}
```

Cat has to implement every method that `Animal` declares.


## Implementing interfaces vs extending abstract classes
The difference is that interfaces are more general and lightweight, and abstract
classes are more special-purpose and feature-rich.

An interface is a way to model a shape. At the value level, that means an object, array, function,
class or class instance. Interfaces do not emit JavaScript code, and only exist at compile time.

## Classes are structurally typed
Like every other type in TypeScript, TypeScript compares classes by their structure, not by their name.

It means that if you have a function that takes a `Zebra` and you give it a `Poodle`,
TypeScript might not mind:

```ts
class Zebra {
  trot() {
    // ...
  }
}

class Poodle {
  trot() {
    // ...
  }
}

function ambleAround(animal: Zebra) {
  animal.trot()
}

let zebra = new Zebra
let poodle = new Poodle
ambleAround(zebra) // Ok
ambleAround(poodle) // Ok!!!
```

As long as `Poodle` is assignable to `Zebra`, TypeScript is OK with it, because
from our function's point of view, the two are interchangeable.

## Classes declare both values and types

```ts
type State = {
  [key: string]: string
}

class StringDatabase {
  state: State = {}
  get(key: string): string | null {
    return key in this.state ? this.state[key] : null
  }
  set(key: string, value: string): void {
    this.state[key] = value
  }
  static from(state: State) {
    let db = new StringDatabase
    for (let key in state) {
      db.set(key, state[key])
    }
    return db
  }
}
```

The instance type `StringDatabase`:

```ts
interface StringDatabase {
  state: State
  get(key: string): string | null
  set(key: string, value: string): void
}
```

The constructor type `typeof StringDatabase`:

```ts
interface StringDatabaseConstructor {
  new(): StringDatabase
  from(state: State): StringDatabase
}
```

Because TypeScript is structurally typed, that's the best we can do to describe
what a class is; **a class is anything that can be `new-ed`**

So, not only does a class declaration generate terms at the value and type levels, but
it generates two terms at the type level: one representing an instance of the class;
one representing the class constructor itself.

## Mixins

JavaScript and TypeScript don't have trait or mixin keywords.
Both are ways to simulate multiple inheritance and do role-oriented programming, a
style of programming where you don't say thing like "this thing is a Shape" but
instead describe properties of a thing, like "it can be measured" or "it has four
sides".
Instead of "is-a" relationships, you describe "can" and "has-a" relationships.

Let's design a debugging library for TypeScript classes.

```ts
class User {
  // ...
}
User.debug()
```

```ts
type ClassConstructor = new(...args: any[]) => {}
function withEZDebug(C extends ClassConstructor)(Class: C) {
  return class extends Class {
    // constructor(...args: any[]) {
    //   super(...args)
    // }
    debug() {
      let Name = Class.constructor.name
      let value = this.getDebugValue()
      return `${Name}(${JSON.stringify(value)})`
    }
  }
}
```

Instead of accepting any old class, we use a generic type to make sure the class
passed into `withEEZDebug` defines  a `.getDebugValue` method:

```ts
type ClassConstructor<T> = new(...args: any[]) => T

function withEZDebug<C extends ClassConstructor<{
  getDebugValue(): object
}>>(Class: C) {
  // ...
}
```

Usage:

```ts
class HardToDebugUser {
  constructor(
    private id: number,
    private firstName: string,
    private lastName: string
  ) {}
  getDebugValue() {
    return {
      id: this.id,
      name: this.firstName + ' ' + this.lastName
    }
  }
}

let User = withEZDebug(HardToDebugUser)
let user = new User(3, 'Emma', 'Gluzman')
user.debug()
```

## Decorators

Decorators gives us a clean syntax for metaprogramming with classes, class methods, properties and method parameters.

Example:
```ts
@serializable
class APIPayload {
  getValue(): Payload {
    // ...
  }
}
```

The `@serializable` decorator wraps `APIPayload` class, and optionally returns a
new class that replaces it.

For each type of decorator, TypeScript requires that you have a function in scope
with the given name and required signature for that type of decorator.

The implementation for each kind of decorator is a regular function that satisfies
a specific signature, depending on what it's decorating.

```ts
type ClassConstructor<T> = new(...args: any[]) => T
function serializable<
  T extends ClassConstructor <{
    getValue(): Payload
  }>
>(Constructor: T) {
  return class extends Constructor {
    serialize() {
      return this.getValue().toString()
    }
  }
}
```

## Simulating final classes

```ts
class MessageQueue {
  private constructor(private messages: string[]) {}
}

class BadQueue extends MessageQueue {} // Error
new MessageQueue([]) // Error
```

Make static method to make a instance:

```ts
class MessageQueue {
  private constructor(private messages: string[]) {}
  static create(messages: string[]) {
    return new MessageQueue(messages)
  }
}
```
