import { NextFunction, Request, Response } from "express";
import { RedisUtil } from "../utils/redis-util";
import { Redis } from "ioredis";
import { setResponse } from "../helpers/response-message";
import { setErrorMessages } from "../helpers/error-message";
import { errorHandlerResponse } from "../exceptions/error-exception";

export const authenticate = async (req: Request, res: Response, next: NextFunction) => {
    const requestId: string = req.get("X-REQUEST-ID") ?? ""
    try {
        const sessionIdUser = req.get("X-SESSION-USER-ID") ?? ""
        if (!sessionIdUser) {
            const errorMessages = setErrorMessages("unauthorized")
            const response = setResponse(null, errorMessages)
            return res.status(401).json(response)
        }
        const user = await RedisUtil.getClient().get(sessionIdUser)
        if (!user) {
            const errorMessages = setErrorMessages("unauthorized")
            const response = setResponse(null, errorMessages)
            return res.status(401).json(response)
        }
        next()
    } catch(e: unknown) {
        return errorHandlerResponse(res, e, requestId)
    }
}