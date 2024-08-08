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
exports.RedisUtil = void 0;
const ioredis_1 = require("ioredis");
class RedisUtil {
    constructor() {
        var _a;
        console.log(new Date().toISOString(), " redis: connecting to redis: " + process.env.PROJECT_USER_REDIS_HOST + ":" + process.env.PROJECT_USER_REDIS_PORT);
        RedisUtil.client = new ioredis_1.Redis({
            host: (_a = process.env.PROJECT_USER_REDIS_HOST) !== null && _a !== void 0 ? _a : "",
            port: Number(process.env.PROJECT_USER_REDIS_PORT),
            db: Number(process.env.PROJECT_USER_REDIS_DATABASE)
        });
        console.log(new Date().toISOString(), " redis: connecteed to redis: " + process.env.PROJECT_USER_REDIS_HOST + ":" + process.env.PROJECT_USER_REDIS_PORT);
    }
    static getInstance() {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                if (!RedisUtil) {
                    RedisUtil.instance = new RedisUtil();
                }
                else {
                    RedisUtil.instance;
                }
                console.log(new Date().toISOString(), " redis: pinging to redis: " + process.env.PROJECT_USER_REDIS_HOST + ":" + process.env.PROJECT_USER_REDIS_PORT);
                yield RedisUtil.client.ping();
                console.log(new Date().toISOString(), " redis: pinged to redis: " + process.env.PROJECT_USER_REDIS_HOST + ":" + process.env.PROJECT_USER_REDIS_PORT);
                return Promise.resolve(RedisUtil.instance);
            }
            catch (e) {
                console.log("error when creating connection to redis: " + process.env.PROJECT_USER_REDIS_HOST + ":" + process.env.PROJECT_USER_REDIS_PORT);
                return Promise.reject(e);
            }
        });
    }
    static getClient() {
        return this.client;
    }
    static disconnect() {
        return __awaiter(this, void 0, void 0, function* () {
            console.log(new Date().toISOString(), " redis: disconnecting to redis:" + process.env.PROJECT_USER_REDIS_HOST + ":" + process.env.PROJECT_USER_REDIS_PORT);
            this.client.disconnect();
            console.log(new Date().toISOString(), " redis: disconnected to redis:" + process.env.PROJECT_USER_REDIS_HOST + ":" + process.env.PROJECT_USER_REDIS_PORT);
        });
    }
}
exports.RedisUtil = RedisUtil;
