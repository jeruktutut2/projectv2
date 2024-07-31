import { ErrorMessage } from "../helpers/error-message";

export class ResponseException extends Error {
    constructor(public status: number, public errorMessages: ErrorMessage[], public message: string) {
        super(message)
    }
}