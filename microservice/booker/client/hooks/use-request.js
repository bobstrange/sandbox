import axios from "axios"
import { useState } from "react"

export default ({ url, method, body }) => {
  const [errors, setErrors] = useState(null)

  const executeRequest = async () => {
    try {
      const response = await axios[method](url, body)
      return response.data
    } catch (error) {
      console.log(error)
      const { errors } = error.response.data
      setErrors(
        <div className="alert alert-danger">
          <ul className="my-0">
            {errors.map((error) => (
              <li key={error.message}>{error.message}</li>
            ))}
          </ul>
        </div>
      )
    }
  }

  return { executeRequest, errors }
}
