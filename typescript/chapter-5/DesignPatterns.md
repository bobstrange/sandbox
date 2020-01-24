# Design Patterns
## Factory

The factory pattern is a way to create objects of some type, leaving the decision of
which concrete object to create to the specific factory that creates that object.

```ts
type Shoe = {
  purpose: string
}

class BalletFlat implements Shoe {
  purpose = 'dancing'
}

class Boot implements Shoe {
  purpose = 'woodcutting'
}

class Sneaker implements Shoe {
  purpose = 'walking
}
```

`Shoe` factory:

```ts
let Shoe = {
  create(type: 'balletFlat' | 'boot' | 'sneaker'): Shoe { // Using a union type
    switch(type) {
      case 'balletFlat': return new BalletFlat
      case 'boot': return new Boot
      case 'sneaker': return new Sneaker
    }
  }
}
```

In this example we use the companion object pattern to declare a type `Shoe` and
a value `Shoe` with the same name (TypeScript has separate namespaces for values and for types)
, as a way to signal that the value provides methods for operating on the type.

## Builder

The builder pattern is a way to separate the construction of an object from the way that
object is actually implemented.

Example:

```ts
new RequestBuilder()
  .setURL('/users')
  .setMethod('get')
  .setData({firstName: 'Anna'})
  .send()
```

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
    // ...
  }
}
```
