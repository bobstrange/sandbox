import express, { Request, Response } from 'express'

const app = express()

app.use(express.json())
app.get('/echo', (req: Request, res: Response) => {
  res.send('Hi there')
})

app.listen(8080, () => {
  console.log('Running on port 8080')
})
