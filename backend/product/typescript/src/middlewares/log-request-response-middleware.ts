import { Request, Response, NextFunction } from "express";
import { setResponse } from "../helpers/response-message";
import { setErrorMessages } from "../helpers/error-message";
import { CustomResponse } from "../types/custom-response-type";

export const printRequestResponseLog = async (req: Request, res: CustomResponse, next: NextFunction) => {
    const requestId =  req.get("X-REQUEST-ID") ?? ""
    // if (!requestId) {
    //     // const errorMessages = setErrorMessages()
    //     return response.status(400).json(setResponse(null, setErrorMessages("aaa")))
    //     // next()
    // }
    // const logRequestBody: Record<string, any> = request.body
    // delete logRequestBody["password"]
    const requestLog = {requestTime: new Date().toISOString(), app: "project-backend-product", method: req.method, requestId: requestId, host: req.hostname, urlPath: req.path, protocol: req.protocol, body: req.body, userAgent: req.headers["user-agent"], remoteAddr: req.ip, forwardedFor: req.header("x-forwarded-for")}
    // log := "responseTime": " + time.Now().String() + ", "app": "project-backend-user", "requestId": " + requestId + ", "response":  + responseBody 
    console.log(requestLog);
    if (!requestId) {
        const errorMessages = setErrorMessages("cannot find requestId")
        return res.status(400).json(setResponse(null, errorMessages))
    }

    next()

    const originalJson = res.json.bind(res)
    res.json = (body: any): Response => {
        res.body = body
        return originalJson(body)
    }
    const responseLog = {responseTime: new Date().toISOString(), app: "project-backend-product", requestId: requestId, response: res.body }
    // log := `{"responseTime": "` + time.Now().String() + `", "app": "project-backend-product", "requestId": "` + requestId + `", "response": ` + responseBody + `}`
    console.log(responseLog);
    
}