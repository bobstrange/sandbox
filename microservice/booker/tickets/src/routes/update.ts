import express, { Request, Response } from 'express'
import { body } from 'express-validator'
import {
  validateRequest,
  NotFoundError,
  UnAuthorizedError,
  requireAuth,
} from '@bsbooker/common'
import { Ticket } from '../models/ticket'

const router = express.Router()

router.put(
  '/api/tickets/:id',
  requireAuth,
  [
    body('title').not().isEmpty().withMessage('Title must be provided'),
    body('price')
      .isFloat({ gt: 0 })
      .withMessage('Price must be provided and must be greater than 0'),
  ],
  validateRequest,
  async (req: Request, res: Response) => {
    const ticket = await Ticket.findById(req.params.id)

    if (!ticket) {
      throw new NotFoundError()
    }

    if (ticket.userId !== req.currentUser!.id) {
      throw new UnAuthorizedError()
    }

    const { title, price } = req.body
    ticket.title = title
    ticket.price = price
    await ticket.save()

    res.send(ticket)
  }
)

export { router as updateTicketRouter }
