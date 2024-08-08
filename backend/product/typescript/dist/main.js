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
const express_1 = require("./setups/express");
const mysql_utils_1 = require("./utils/mysql-utils");
const elasticsearch_util_1 = require("./utils/elasticsearch-util");
const app = express_1.web.listen(process.env.PROJECT_PRODUCT_APPLICATION_HOST, () => __awaiter(void 0, void 0, void 0, function* () {
    try {
        yield mysql_utils_1.MysqlUtil.getInstance();
        yield elasticsearch_util_1.ElasticsearchUtil.getInstance();
        console.log(new Date().toISOString(), "express: listening on " + process.env.PROJECT_PRODUCT_APPLICATION_HOST);
    }
    catch (e) {
        yield mysql_utils_1.MysqlUtil.close();
        process.exit(1);
    }
    finally {
    }
}));
process.on("SIGINT", () => __awaiter(void 0, void 0, void 0, function* () {
    yield mysql_utils_1.MysqlUtil.close();
    app.close(() => {
        console.log(new Date().toISOString(), "express: closed on " + process.env.PROJECT_PRODUCT_APPLICATION_HOST);
    });
    process.exit(0);
}));
