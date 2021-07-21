import { Response } from "express"

const unauthenticatedResponse = (res: Response, message: string) => {
  res.status(401).json({
    message
  });
}

export default unauthenticatedResponse;