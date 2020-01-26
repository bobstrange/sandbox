import { LocalDataSource } from './data/local_data_source'
import { HtmlDisplay } from './html_display'
import 'bootstrap/dist/css/bootstrap.css'

let dataSource = new LocalDataSource()

async function displayData(): Promise<HTMLElement> {
  let display = new HtmlDisplay()
  display.props = {
    products: await dataSource.getProducts("name"),
    order: await dataSource.order
  }
  return display.getContent()
}

document.onreadystatechange = async () => {
  if (document.readyState === 'complete') {
    const element = await displayData()
    const rootElement = document.getElementById('app')
    rootElement.innerHTML = ''
    rootElement.appendChild(element)
  }
}
