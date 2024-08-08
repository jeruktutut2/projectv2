"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Validation = void 0;
const response_exception_1 = require("../exceptions/response-exception");
class Validation {
    static validate(schema, data) {
        const result = schema.safeParse(data);
        if (result.error) {
            const errorMessages = [];
            for (let i = 0; i < result.error.errors.length; i++) {
                const errorMessage = {
                    field: result.error.errors[i].path[0].toString(),
                    message: result.error.errors[i].message
                };
                errorMessages.push(errorMessage);
            }
            throw new response_exception_1.ResponseException(400, errorMessages, "validation error");
        }
        return result.data;
    }
}
exports.Validation = Validation;
