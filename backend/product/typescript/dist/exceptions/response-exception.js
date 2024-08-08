"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ResponseException = void 0;
class ResponseException extends Error {
    constructor(status, errorMessages, message) {
        super(message);
        this.status = status;
        this.errorMessages = errorMessages;
        this.message = message;
    }
}
exports.ResponseException = ResponseException;
