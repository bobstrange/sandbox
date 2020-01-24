import { TodoItem } from './todo_item'
import { TodoCollection } from './todo_collection'

let todos = [
  new TodoItem(1, 'Read Effective Typescript'),
  new TodoItem(2, 'Running'),
  new TodoItem(3, 'Learn k8s')
]

let collection = new TodoCollection('Bob', todos)

console.clear()
console.info(`${collection.userName}'s Todo list`)


let newId = collection.addTodo('Workout')
let todoItem = collection.getTodoById(newId)

console.info(JSON.stringify(todoItem))
collection.markComplete(4, true)

collection.removeComplete()
collection.getTodoItems(true).forEach(item => item.printDetails())
