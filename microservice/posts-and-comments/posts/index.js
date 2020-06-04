const express = require('express')
const { randomBytes } = require('crypto')
const bodyParser = require('body-parser')

const app = express()
app.use(bodyParser.json())

const posts = {}

app.get('/posts', (req, res) => {
  res.send(posts)
})

app.post('/posts', (req, res) => {
  const id = randomBytes(4).toString('hex')
  const { title } = req.body
  posts[id] = { id, title }

  res.status(201).send(posts[id])
})

app.listen(8080, () => {
  console.log('Listening on port 8080')
})
