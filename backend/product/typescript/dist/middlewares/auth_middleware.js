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
exports.authenticate = void 0;
const redis_util_1 = require("../utils/redis-util");
const response_message_1 = require("../helpers/response-message");
const error_message_1 = require("../helpers/error-message");
const error_exception_1 = require("../exceptions/error-exception");
const authenticate = (req, res, next) => __awaiter(void 0, void 0, void 0, function* () {
    var _a, _b;
    const requestId = (_a = req.get("X-REQUEST-ID")) !== null && _a !== void 0 ? _a : "";
    try {
        const sessionIdUser = (_b = req.get("X-SESSION-USER-ID")) !== null && _b !== void 0 ? _b : "";
        if (!sessionIdUser) {
            const errorMessages = (0, error_message_1.setErrorMessages)("unauthorized");
            const response = (0, response_message_1.setResponse)(null, errorMessages);
            return res.status(401).json(response);
        }
        const user = yield redis_util_1.RedisUtil.getClient().get(sessionIdUser);
        if (!user) {
            const errorMessages = (0, error_message_1.setErrorMessages)("unauthorized");
            const response = (0, response_message_1.setResponse)(null, errorMessages);
            return res.status(401).json(response);
        }
        next();
    }
    catch (e) {
        return (0, error_exception_1.errorHandlerResponse)(res, e, requestId);
    }
});
exports.authenticate = authenticate;
