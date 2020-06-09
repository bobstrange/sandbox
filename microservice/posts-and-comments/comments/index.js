const express = require('express')
const bodyParser = require('body-parser')
const { randomBytes } = require('crypto')
const cors = require('cors')
const axios = require('axios')

const app = express()
app.use(bodyParser.json())
app.use(cors())

const commentsByPostId = {}

app.get('/posts/:id/comments', (req, res) => {
  res.send(commentsByPostId[req.params.id] || [])
})

app.post('/posts/:id/comments', async (req, res) => {
  const commentId = randomBytes(4).toString('hex')
  const { content } = req.body

  const comments = commentsByPostId[req.params.id] || []
  const comment = {
    id: commentId,
    content,
    status: 'pending'
  }

  comments.push(comment)
  commentsByPostId[req.params.id] = comments

  await axios.post('http://event-bus-cluster-ip-srv:8085/events', {
    type: 'CommentCreated',
    data: { ...comment, postId: req.params.id }
  })

  res.status(201).send(comments)
})

app.post('/events', (req, res) => {
  const { type, data } = req.body
  console.log(`Received event: ${type}`)

  if (type === 'CommentModerated') {
    const { id, postId, status, content } = data
    const comments = commentsByPostId[postId]
    const comment = comments.find(_ => {
      return _.id === id
    })
    comment.status = status

    axios.post('http://event-bus-cluster-ip-srv:8085/events', {
      type: 'CommentUpdated',
      data: {
        id,
        content,
        postId,
        status
      }
    })
  }

  res.send({})
})


app.listen(8081, () => {
  console.log('Listening on port 8081')
})
