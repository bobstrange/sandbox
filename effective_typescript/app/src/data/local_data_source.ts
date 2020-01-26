import { AbstractDataSource } from './abstract_data_source'
import { Product } from './entities'

export class LocalDataSource extends AbstractDataSource {
  async loadProducts(): Promise<Product[]> {
    // return Promise.resolve([
    //   { id: 1, name: 'Dummy1', category: 'Watersports', description: 'Dummy1 desc', price: 100 },
    //   { id: 2, name: 'Dummy2', category: 'Watersports', description: 'Dummy2 desc', price: 200 },
    //   { id: 3, name: 'Dummy3', category: 'Running', description: 'Dummy3 desc', price: 300 },
    //   { id: 4, name: 'Dummy4', category: 'Bicycle', description: 'Dummy4 desc', price: 400 },
    //   { id: 5, name: 'Dummy5', category: 'Bicycle', description: 'Dummy5 desc', price: 500 }
    // ])
    return await ([
      { id: 1, name: 'Dummy1', category: 'Watersports', description: 'Dummy1 desc', price: 100 },
      { id: 2, name: 'Dummy2', category: 'Watersports', description: 'Dummy2 desc', price: 200 },
      { id: 3, name: 'Dummy3', category: 'Running', description: 'Dummy3 desc', price: 300 },
      { id: 4, name: 'Dummy4', category: 'Bicycle', description: 'Dummy4 desc', price: 400 },
      { id: 5, name: 'Dummy5', category: 'Bicycle', description: 'Dummy5 desc', price: 500 }
    ])
  }

  storeOrder(): Promise<number> {
    console.log('Store Order')
    console.log(JSON.stringify(this.order))
    return Promise.resolve(1)
  }

}
