const express = require('express');
const redis = require('redis');

const app = express();
const redisClient = redis.createClient({
  host: 'redis-server',
  port: 6379
});

redisClient.set('visits', 0);

app.get('/', (req, res) => {
  redisClient.get('visits', (err, visits) => {
    res.send(`Number of visits is ${visits}`);
    redisClient.set('visits', parseInt(visits) + 1)
  })
})

app.listen(8080, () => {
  console.log(`Listening on port 8080`);
})
