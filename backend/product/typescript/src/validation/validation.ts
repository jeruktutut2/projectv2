import { ZodType } from "zod";
// import { CreateProductErrorMessage } from "../features/create-product/models/create-product-error-message";
import { ResponseException } from "../exceptions/response-exception";
import { ErrorMessage } from "../helpers/error-message";

export class Validation {
    static validate<T>(schema: ZodType, data: T): T {
        const result = schema.safeParse(data)
        if (result.error) {
            // console.log("result.error:", result.error.errors);
            const errorMessages: ErrorMessage[] = []
            for (let i = 0; i < result.error.errors.length; i++) {
                const errorMessage: ErrorMessage = {
                    field: result.error.errors[i].path[0].toString(),
                    message: result.error.errors[i].message
                }
                // const field = result.error.errors[i].path[0]
                // const message = result.error.errors[i].message
                errorMessages.push(errorMessage)
            }
            throw new ResponseException(400, errorMessages, "validation error")
        }
        
        // return schema.parse(data)
        return result.data
    }
}