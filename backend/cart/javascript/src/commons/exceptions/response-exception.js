export class ResponseException extends Error {
    constructor(status, errorMessages, message) {
        super(message)
        this.status = status
        this.errorMessages = errorMessages
    }
}