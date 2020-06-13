import express from 'express'
import { json } from 'body-parser'
import { currentUserRouter } from './routes/current-user'

const app = express()
app.use(json())

app.use(currentUserRouter)

app.listen(8080, () => {
  console.log('Listening on port 8080')
  console.log('hi there')
})

