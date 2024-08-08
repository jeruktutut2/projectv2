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
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.MysqlUtil = void 0;
const promise_1 = __importDefault(require("mysql2/promise"));
class MysqlUtil {
    constructor() {
        var _a, _b, _c;
        console.log(new Date().toISOString(), "mysql: connecting to", process.env.PROJECT_PRODUCT_MYSQL_HOST);
        const connectionString = `mysql://` + process.env.PROJECT_PRODUCT_MYSQL_USERNAME + `:` + process.env.PROJECT_PRODUCT_MYSQL_PASSWORD + `@` + process.env.PROJECT_PRODUCT_MYSQL_HOST + `/` + process.env.PROJECT_PRODUCT_MYSQL_DATABASE + ``;
        const access = {
            uri: connectionString,
            connectionLimit: parseInt((_a = process.env.PROJECT_PRODUCT_MYSQL_MAX_OPEN_CONNECTION) !== null && _a !== void 0 ? _a : "0"),
            maxIdle: parseInt((_b = process.env.PROJECT_PRODUCT_MYSQL_MAX_IDLE_CONNECTION) !== null && _b !== void 0 ? _b : "0"),
            idleTimeout: parseInt((_c = process.env.PROJECT_PRODUCT_MYSQL_CONNECTION_MAX_IDLETIME) !== null && _c !== void 0 ? _c : "0")
        };
        MysqlUtil.pool = promise_1.default.createPool(access);
        console.log(new Date().toISOString(), "mysql: connected to", process.env.PROJECT_PRODUCT_MYSQL_HOST);
    }
    static getInstance() {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                if (!MysqlUtil.instance) {
                    MysqlUtil.instance = new MysqlUtil();
                }
                else {
                    MysqlUtil.instance;
                }
                console.log(new Date().toISOString(), "mysql: pinging to", process.env.PROJECT_PRODUCT_MYSQL_HOST);
                const pool = MysqlUtil.getPool();
                const connection = yield pool.getConnection();
                yield connection.ping();
                console.log(new Date().toISOString(), "mysql: pingged to", process.env.PROJECT_PRODUCT_MYSQL_HOST);
                return Promise.resolve(MysqlUtil.instance);
            }
            catch (e) {
                console.log("error when creating connection mysql to " + process.env.PROJECT_PRODUCT_MYSQL_HOST + " : " + e);
                return Promise.reject(e);
            }
        });
    }
    static getPool() {
        return this.pool;
    }
    static close() {
        return __awaiter(this, void 0, void 0, function* () {
            if (this.pool) {
                console.log(new Date().toISOString(), "mysql: closing to", process.env.PROJECT_PRODUCT_MYSQL_HOST);
                yield this.pool.end();
                console.log(new Date().toISOString(), "mysql: closed to", process.env.PROJECT_PRODUCT_MYSQL_HOST);
            }
        });
    }
}
exports.MysqlUtil = MysqlUtil;
