import { Request, Response, NextFunction } from "express";
import { setResponse } from "../helpers/response-message";
import { setErrorMessages } from "../helpers/error-message";
import { CustomResponse } from "../types/custom-response-type";

export const printRequestResponseLog = async (req: Request, res: CustomResponse, next: NextFunction) => {
    const requestId =  req.get("X-REQUEST-ID") ?? ""
    const requestLog = {requestTime: new Date().toISOString(), app: "project-backend-product", method: req.method, requestId: requestId, host: req.hostname, urlPath: req.path, protocol: req.protocol, body: req.body, userAgent: req.headers["user-agent"], remoteAddr: req.ip, forwardedFor: req.header("x-forwarded-for")}
    console.log(requestLog);
    if (!requestId) {
        const errorMessages = setErrorMessages("cannot find requestId")
        return res.status(400).json(setResponse(null, errorMessages))
    }

    const originalSend = res.send;
    const originalStatus = res.status
    let responseBody: any
    let responseStatus: any
    res.send = function (body?: any): Response {
        responseBody = body
        return originalSend.call(this, body);
    };
    res.status = function(statusCode: number): Response {
        responseStatus = statusCode
        return originalStatus.call(this, statusCode)
    }

    next()

    res.on("finish", () => {
        const responseLog = {responseTime: new Date().toISOString(), app: "project-backend-product", requestId: requestId, responseStatus: responseStatus, response: responseBody }
        console.log(responseLog);
    })
}