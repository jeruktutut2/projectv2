import { ResponseException } from "./response-exception.js";
import responseMessageHelper from "../helpers/response-message-helper.js";
export const errorHandler = (error, requestId) => {
    console.log(JSON.stringify({logTime: (new Date()).toISOString(), app: "backend-cart-nodejs", requestId: requestId, error: error.stack}));
    if (error instanceof ResponseException) {
        throw new ResponseException(error.status, error.ErrorMessages, error.message)
    } else {
        throw new ResponseException(500, setInternalServerErrorMessage(), "internal server error")
    }
}

export const setErrorMessages = (message) => {
    return [{field: "message", message: message}]
}

export const setInternalServerErrorMessage = () => {
    return [{field: "message", message: "internal server error"}]
}

export const errorHandlerResponse = (res, error) => {
    let status
    let response
    if (error instanceof ResponseException) {
        status = error.status
        response = responseMessageHelper.setResponse(null, error.ErrorMessages)
    } else {
        status = 500
        response = responseMessageHelper.setResponse(null, setInternalServerErrorMessage())
    }
    return res.status(status).json(response)
}