"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.printRequestResponseLog = void 0;
const response_message_1 = require("../helpers/response-message");
const error_message_1 = require("../helpers/error-message");
const printRequestResponseLog = (req, res, next) => __awaiter(void 0, void 0, void 0, function* () {
    var _a;
    const requestId = (_a = req.get("X-REQUEST-ID")) !== null && _a !== void 0 ? _a : "";
    const requestLog = { requestTime: new Date().toISOString(), app: "project-backend-product", method: req.method, requestId: requestId, host: req.hostname, urlPath: req.path, protocol: req.protocol, body: req.body, userAgent: req.headers["user-agent"], remoteAddr: req.ip, forwardedFor: req.header("x-forwarded-for") };
    console.log(requestLog);
    if (!requestId) {
        const errorMessages = (0, error_message_1.setErrorMessages)("cannot find requestId");
        return res.status(400).json((0, response_message_1.setResponse)(null, errorMessages));
    }
    const originalSend = res.send;
    const originalStatus = res.status;
    let responseBody;
    let responseStatus;
    res.send = function (body) {
        responseBody = body;
        return originalSend.call(this, body);
    };
    res.status = function (statusCode) {
        responseStatus = statusCode;
        return originalStatus.call(this, statusCode);
    };
    next();
    res.on("finish", () => {
        const responseLog = { responseTime: new Date().toISOString(), app: "project-backend-product", requestId: requestId, responseStatus: responseStatus, response: responseBody };
        console.log(responseLog);
    });
});
exports.printRequestResponseLog = printRequestResponseLog;
