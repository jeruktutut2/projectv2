import { setErrorMessages } from "../exceptions/error-exception.js";
import responseMessageHelper from "../helpers/response-message-helper.js";
const printLogRequestResponse = (req, res, next) => {
    const requestId =  req.get("X-REQUEST-ID") ?? ""
    const requestLog = {requestTime: new Date().toISOString(), app: "project-backend-cart", method: req.method, requestId: requestId, host: req.hostname, urlPath: req.path, protocol: req.protocol, body: req.body, userAgent: req.headers["user-agent"], remoteAddr: req.ip, forwardedFor: req.header("x-forwarded-for")}
    console.log(requestLog);
    if (!requestId) {
        const errorMessages = setErrorMessages("cannot find requestId")
        return res.status(400).json(responseMessageHelper.setResponse(null, errorMessages))
    }

    const originalSend = res.send;
    const originalStatus = res.status
    let responseBody
    let responseStatus
    res.send = function (body) {
        responseBody = body
        return originalSend.call(this, body);
    };
    res.status = function(statusCode) {
        responseStatus = statusCode
        return originalStatus.call(this, statusCode)
    }

    next()

    res.on("finish", () => {
        const responseLog = {responseTime: new Date().toISOString(), app: "project-backend-cart", requestId: requestId, responseStatus: responseStatus, response: responseBody }
        console.log(responseLog);
    })
}

export default {
    printLogRequestResponse
}