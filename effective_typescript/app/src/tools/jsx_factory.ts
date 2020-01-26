export function createElement(tag: any, props: Object, ...children: Object[]): HTMLElement {
  function addChild(element: HTMLElement, child: any) {
    element.appendChild(child instanceof Node ? child : document.createTextNode(child.toString()))
  }

  if (typeof tag === 'function') {
    // @ts-ignore
    return Object.assign(new tag(), { props: props || {}}.getContent())
  }

  const element = Object.assign(document.createElement(tag), props || {})
  children.forEach(child => Array.isArray(child) ? child.forEach(c => addChild(element, c)) : addChild(element, child))
}

declare global {
  namespace JSX {
    interface ElementAttributesProperty { props }
  }
}
