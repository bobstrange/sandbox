export class DatabaseConnectionError extends Error {
  reason = 'Failed to connect database'

  constructor() {
    super()

    Object.setPrototypeOf(this, DatabaseConnectionError.prototype)
  }
}
