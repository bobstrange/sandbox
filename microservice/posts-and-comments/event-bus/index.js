const express = require('express')
const bodyParser = require('body-parser')
const axios = require('axios')

const app = express()
app.use(bodyParser.json())

const events = []

app.post('/events', (req, res) => {
  const event = req.body

  events.push(event)

  axios.post("http://posts-cluster-ip-srv:8080/events", event);
  axios.post('http://comments-cluster-ip-srv:8081/events', event)
  axios.post('http://query-cluster-ip-srv:8082/events', event)
  axios.post('http://moderation-cluster-ip-srv:8083/events', event)

  res.send({ status: 'OK' })
})

app.get('/events', (req, res) => {
  res.send(events)
})

app.listen(8085, () => {
  console.log('Listening on 8085')
})
