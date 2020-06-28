import express, { Request, Response } from 'express'
import { requireAuth } from '@bsbooker/common'

const router = express.Router()

router.post('/api/tickets', requireAuth, (req: Request, res: Response) => {
  res.status(200).send({})
})

export { router as createTicketRouter }
