export interface ErrorMessage {
    field: string
    message: string
}

export const setErrorMessages = (message: string): ErrorMessage[] => {
    const errorMessages: ErrorMessage[] = [{
        field: "message",
        message: message
    }]
    return errorMessages
}

export const setInternalServerErrorErrorMessages = (): ErrorMessage[] => {
    const errorMessages: ErrorMessage[] = [{
        field: "message",
        message: "internal server error"
    }]
    return errorMessages
}