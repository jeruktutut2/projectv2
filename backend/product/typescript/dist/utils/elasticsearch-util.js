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
exports.ElasticsearchUtil = void 0;
const elasticsearch_1 = require("@elastic/elasticsearch");
class ElasticsearchUtil {
    constructor() {
        console.log(new Date().toISOString(), "elasticsearch: connecting to", process.env.PROJECT_PRODUCT_ELASTICSEARCH_NODE);
        ElasticsearchUtil.client = new elasticsearch_1.Client({
            node: process.env.PROJECT_PRODUCT_ELASTICSEARCH_NODE
        });
        console.log(new Date().toISOString(), "elasticsearch: connected to", process.env.PROJECT_PRODUCT_ELASTICSEARCH_NODE);
    }
    static getInstance() {
        return __awaiter(this, void 0, void 0, function* () {
            try {
                if (!ElasticsearchUtil.instance) {
                    ElasticsearchUtil.instance = new ElasticsearchUtil();
                }
                else {
                    ElasticsearchUtil.instance;
                }
                console.log(new Date().toISOString(), "elasticsearch: pinging to", process.env.PROJECT_PRODUCT_ELASTICSEARCH_NODE);
                yield ElasticsearchUtil.client.ping();
                console.log(new Date().toISOString(), "elasticsearch: pinged to", process.env.PROJECT_PRODUCT_ELASTICSEARCH_NODE);
                return Promise.resolve(ElasticsearchUtil.instance);
            }
            catch (e) {
                console.log("error when creating connection elasticsearch to " + process.env.PROJECT_PRODUCT_ELASTICSEARCH_NODE + " : " + e);
                return Promise.reject(e);
            }
        });
    }
    static getClient() {
        return this.client;
    }
    static closeClient() {
        return __awaiter(this, void 0, void 0, function* () {
            if (ElasticsearchUtil.client) {
                console.log(new Date().toISOString(), "elasticsearch: closing to", process.env.PROJECT_PRODUCT_ELASTICSEARCH_NODE);
                yield ElasticsearchUtil.client.close();
                console.log(new Date().toISOString(), "elasticsearch: closed to", process.env.PROJECT_PRODUCT_ELASTICSEARCH_NODE);
            }
        });
    }
}
exports.ElasticsearchUtil = ElasticsearchUtil;
