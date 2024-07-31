export interface Response {
    data: any
    errors: any
}

export const setResponse = (data: any, errors: any): Response => {
    const response: Response = {
        data: data,
        errors: errors
    }
    return response
}