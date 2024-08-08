"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.setInternalServerErrorErrorMessages = exports.setErrorMessages = void 0;
const setErrorMessages = (message) => {
    const errorMessages = [{
            field: "message",
            message: message
        }];
    return errorMessages;
};
exports.setErrorMessages = setErrorMessages;
const setInternalServerErrorErrorMessages = () => {
    const errorMessages = [{
            field: "message",
            message: "internal server error"
        }];
    return errorMessages;
};
exports.setInternalServerErrorErrorMessages = setInternalServerErrorErrorMessages;
