/** Difference with types and interfaces */

interface A {
  good(x: number): string
  bad(x: number): string
}

interface B extends A { // Error Type 'number' is not assignable to type 'string'
  good(x: string | number): string
  bad(x: string): string
}

type C = {
  good(x: number): string
  bad(x: number): string
}

type D = C & { // No error
  good(x: number | string): string
  bad(x: number): string
}

/** Classes are structurally typed */
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

/** Classes declare both values and types */

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
