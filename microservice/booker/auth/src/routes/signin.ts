import express, { Request, Response } from "express";
import { body, validationResult } from 'express-validator'
import { RequestValidationError } from "../errors/request-validation-error";
import { validateRequest } from "../middlewares/validate-request";

const router = express.Router();

const validations = [
  body('email').isEmail().withMessage('Email must be valid'),
  body('password').trim().notEmpty().withMessage('You must supply a password')
]

router.post(
  '/api/users/signin',
  validations,
  validateRequest,
  (req: Request, res: Response) => {
  }
);

export { router as signinRouter };
