import { ResponseException } from "../exceptions/response-exception.js";

export const validate = (schema, data) => {
    const result = schema.validate(data, {
        abortEarly: false,
        allowUnknown: false
    })

    if (result.error) {
        const errorMessage = []
        for (let i = 0; i < result.error.details.length; i++) {
            const field = result.error.details[i].path[0]
            const fieldMessage = result.error.details[i].message.replaceAll('"', '')
            const message = {field: field, message: fieldMessage}
            errorMessage.push(message)
        }
        throw new ResponseException(400, errorMessage, "validation error")
    } else {
        return result.value
    }
}