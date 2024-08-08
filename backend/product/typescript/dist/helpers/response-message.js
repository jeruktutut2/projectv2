"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.setResponse = void 0;
const setResponse = (data, errors) => {
    const response = {
        data: data,
        errors: errors
    };
    return response;
};
exports.setResponse = setResponse;
