import request from 'supertest'
import { app } from '../../app'
import mongoose from 'mongoose'

describe('updateTicketRoute', () => {
  it('returns a 404 if the provided id does not exist', async () => {
    const id = new mongoose.Types.ObjectId().toHexString
    await request(app)
      .put(`/api/tickets/${id}`)
      .set('Cookie', global.signin())
      .send({
        title: 'test1',
        price: 10,
      })
      .expect(404)
  })

  it('returns a 401 if the user is not authenticated', async () => {
    const id = new mongoose.Types.ObjectId().toHexString
    const response = await request(app).put(`/api/tickets/${id}`).send({
      title: 'test1',
      price: 10,
    })
    console.log(response.body)
    expect(response.status).toEqual(401)
  })
  it('returns a 401 if the user does not own the ticket', async () => {})
  it('returns a 400 if the user provides an invalid title or price', async () => {})
  it('updates the ticket with given title and price', async () => {})
})
