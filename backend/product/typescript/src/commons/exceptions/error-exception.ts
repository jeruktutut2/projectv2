import { setErrorMessages, setInternalServerErrorErrorMessages } from "../helpers/error-message";
import { setResponse } from "../helpers/response-message";
import { ResponseException } from "./response-exception";
import { Response } from "express";

export function errorHandler(error: unknown, requestId: string) {
    if (error instanceof Error) {
        console.log(JSON.stringify({logTime: (new Date()).toISOString(), app: "backend-product-typescript", requestId: requestId, error: error.stack}));
        if (error instanceof ResponseException) {
            throw new ResponseException(error.status, error.errorMessages, error.message)
        } else {
            throw new ResponseException(500, setInternalServerErrorErrorMessages(), "internal server error")
        }
    } else {
        throw new ResponseException(500, setInternalServerErrorErrorMessages(), "internal server error")
    }
}

export function errorHandlerResponse(res: Response, error: unknown, requestId: string): Response {
    let status: number = 0
    let response: any = null
    if (error instanceof Error) {
        if (error instanceof ResponseException) {
            status = error.status
            response = setResponse(null, error.errorMessages)
        } else {
            status = 500
            response = setResponse(null, setInternalServerErrorErrorMessages())
        }
    } else {
        status = 500
        response = setResponse(null, setInternalServerErrorErrorMessages())
    }
    return res.status(status).json(response)
}