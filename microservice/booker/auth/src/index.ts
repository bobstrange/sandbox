import express from 'express'
import 'express-async-errors'
import { json } from 'body-parser'

import { signinRouter } from './routes/signin'
import { signoutRouter } from './routes/signout'
import { signupRouter } from './routes/signup'
import { currentUserRouter } from './routes/current-user'

import { errorHandler } from './middlewares/error-handler'
import { NotFoundError } from './errors/not-found-error'

const app = express()
app.use(json())

app.use(signinRouter)
app.use(signoutRouter)
app.use(signupRouter)
app.use(currentUserRouter)

app.all('*', async () => {
  throw new NotFoundError()
})

app.use(errorHandler)

app.listen(8080, () => {
  console.log('Listening on port 8080')
  console.log('hi there')
})
