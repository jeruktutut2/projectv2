"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.errorHandler = errorHandler;
exports.errorHandlerResponse = errorHandlerResponse;
const error_message_1 = require("../helpers/error-message");
const response_message_1 = require("../helpers/response-message");
const response_exception_1 = require("./response-exception");
function errorHandler(error, requestId) {
    if (error instanceof Error) {
        console.log(JSON.stringify({ logTime: (new Date()).toISOString(), app: "backend-product-typescript", requestId: requestId, error: error.stack }));
        if (error instanceof response_exception_1.ResponseException) {
            throw new response_exception_1.ResponseException(error.status, error.errorMessages, error.message);
        }
        else {
            throw new response_exception_1.ResponseException(500, (0, error_message_1.setInternalServerErrorErrorMessages)(), "internal server error");
        }
    }
    else {
        throw new response_exception_1.ResponseException(500, (0, error_message_1.setInternalServerErrorErrorMessages)(), "internal server error");
    }
}
function errorHandlerResponse(res, error, requestId) {
    let status = 0;
    let response = null;
    if (error instanceof Error) {
        if (error instanceof response_exception_1.ResponseException) {
            status = error.status;
            response = (0, response_message_1.setResponse)(null, error.errorMessages);
        }
        else {
            status = 500;
            response = (0, response_message_1.setResponse)(null, (0, error_message_1.setInternalServerErrorErrorMessages)());
        }
    }
    else {
        status = 500;
        response = (0, response_message_1.setResponse)(null, (0, error_message_1.setInternalServerErrorErrorMessages)());
    }
    return res.status(status).json(response);
}
