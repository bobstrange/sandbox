import express, { Request, Response } from 'express'
import { body, validationResult } from 'express-validator'

const router = express.Router()

const validations = [
  body('email').isEmail().withMessage('Email must be valid'),
  body('password').trim().isLength({ min: 4, max: 20 }).withMessage('Password must be between 4 and 20 characters')
]

router.post('/api/users/signup', validations, (req: Request, res: Response) => {
  const errors = validationResult(req)

  if (!errors.isEmpty()) {
    throw new Error('Invalid email ror password')
  }

  const { email, password } = req.body
})

export { router as signupRouter }
