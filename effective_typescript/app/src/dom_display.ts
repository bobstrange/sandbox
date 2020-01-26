import { Product, Order } from './data/entities'

export class DomDisplay {
  props: {
    products: Product[],
    order: Order
  }

  getContent(): HTMLElement {
    let element = document.createElement('h3')
    element.innerText = this.getElementText()
    element.classList.add('bg-primary', 'text-center', 'text-white', 'p-2')
    return element
  }

  getElementText() {
    return `${this.props.products.length} Products,`
      + `Order total: $${this.props.order.total}`
  }
}
