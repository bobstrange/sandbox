# Exercise
## What are differences between a class and an interface ?
An interface just defines a type.
A class defines both a type and a value.

## When you mark a class's constructor as `private`, that means you can't instantiate or extend the class. What happens when you mark it as `protected` instead ?

Still it cannot use new.

## Extend the implementation to make it safer. Update the implementation so that a consumer knows at compile time that calling `Shoe.create('boot') returns a `Boot` and `Shoe.create('balletFlat')` returns a `BalletFlat`

```ts
interface Shoe {
  purpose: string
}

class BalletFlat implements Shoe {
  purpose = 'dancing'
}

class Boot implements Shoe {
  purpose = 'woodcutting'
}

class Sneaker implements Shoe {
  purpose = 'walking'
}

type ShoeFactory = {
  create(type: 'balletflat'): BalletFlat
  create(type: 'boot'): Boot
  create(type: 'sneaker'): Sneaker
}

let Shoe: ShoeFactory = {
  create(type: 'balletflat' | 'boot' | 'sneaker'): Shoe {
    switch(type) {
      case 'balletflat': return new BalletFlat
      case 'boot': return new Boot
      case 'sneaker': return new Sneaker
    }
  }
}
```

## Extend the Builder pattern

```ts
class RequestBuilder {
  private data: object | null = null
  private method: 'get' | 'post' | null = null
  private url: string | null = null

  setURL(url: string): this {
    this.url = url
    return this
  }

  setMethod(method: 'get' | 'post'): this {
    this.method = method
    return this
  }

  setData(data: object): this {
    this.data = data
    return this
  }

  send() {
  }
}

new RequestBuilder().setURL('/users/').send() // Should throw compile-time error
```
