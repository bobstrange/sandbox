import { LocalDataSource } from './data/local_data_source'

async function displayData(): Promise<string> {
  let dataSource = new LocalDataSource()
  let allProducts = await dataSource.getProducts("name")
  let categories = await dataSource.getCategories()
  let bicycleProducts = await dataSource.getProducts("name", "Bicycle")

  let result = ""

  allProducts.forEach(product => result += `Product: ${product.name}, ${product.category}\n`)
  categories.forEach(category => result += `Category: ${category}\n`)
  bicycleProducts.forEach(product => dataSource.order.addProduct(product, 1))
  result += `Order total: $${dataSource.order.total.toFixed(2)}`
  return result
}

displayData().then(res => console.log(res))
