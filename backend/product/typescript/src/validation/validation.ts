import { ZodType } from "zod";
import { ResponseException } from "../exceptions/response-exception";
import { ErrorMessage } from "../helpers/error-message";

export class Validation {
    static validate<T>(schema: ZodType, data: T): T {
        const result = schema.safeParse(data)
        if (result.error) {
            const errorMessages: ErrorMessage[] = []
            for (let i = 0; i < result.error.errors.length; i++) {
                const errorMessage: ErrorMessage = {
                    field: result.error.errors[i].path[0].toString(),
                    message: result.error.errors[i].message
                }
                errorMessages.push(errorMessage)
            }
            throw new ResponseException(400, errorMessages, "validation error")
        }

        return result.data
    }
}