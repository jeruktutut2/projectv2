const printLogRequestResponse = (req, res, next) => {
    const requestId =  req.get("X-REQUEST-ID") ?? ""
    const requestLog = {requestTime: new Date().toISOString(), app: "project-backend-cart", method: req.method, requestId: requestId, host: req.hostname, urlPath: req.path, protocol: req.protocol, body: req.body, userAgent: req.headers["user-agent"], remoteAddr: req.ip, forwardedFor: req.header("x-forwarded-for")}
    console.log(requestLog);
    if (!requestId) {
        const errorMessages = setErrorMessages("cannot find requestId")
        return res.status(400).json(setResponse(null, errorMessages))
    }
    next()
    const originalJson = res.json.bind(res)
    res.json = (body) => {
        res.body = body
        return originalJson(body)
    }
    const responseLog = {responseTime: new Date().toISOString(), app: "project-backend-cart", requestId: requestId, response: res.body }
}

export default {
    printLogRequestResponse
}