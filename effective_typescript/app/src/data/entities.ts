export type Product = {
  id: number,
  name: string,
  description: string,
  category: string,
  price: number
}

export class OrderLine {
  constructor(public product: Product, public quantity: number) {}
  get total(): number {
    return this.product.price * this.quantity
  }
}

export class Order {
  private _orderLines = new Map<number, OrderLine>()
  constructor(orderLines?: OrderLine[]) {
    if (orderLines) {
      orderLines.forEach(orderLine => {
        this._orderLines.set(orderLine.product.id, orderLine)
      })
    }
  }

  public addProduct(product: Product, quantity: number) {
    if (this._orderLines.has(product.id)) {
      if (quantity === 0) {
        this.removeProduct(product.id)
      } else {
        this._orderLines.get(product.id)!.quantity += quantity
      }
    } else {
      this._orderLines.set(product.id, new OrderLine(product, quantity))
    }
  }

  public removeProduct(id: number) {
    this._orderLines.delete(id)
  }

  get orderLines(): OrderLine[] {
    return [...this._orderLines.values()]
  }

  get productCount(): number {
    return this.orderLines.reduce((total, orderLine) => {
      return total + orderLine.quantity
    }, 0)
  }

  get total(): number {
    return this.orderLines.reduce((total, orderLine) => {
      return total + orderLine.total
    }, 0)
  }
}
