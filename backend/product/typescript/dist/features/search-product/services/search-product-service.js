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
exports.SearchProductService = void 0;
const elasticsearch_util_1 = require("../../../utils/elasticsearch-util");
const response_exception_1 = require("../../../exceptions/response-exception");
const error_exception_1 = require("../../../exceptions/error-exception");
const error_message_1 = require("../../../helpers/error-message");
class SearchProductService {
    static searchProduct(requestId, keyword) {
        return __awaiter(this, void 0, void 0, function* () {
            let client = null;
            try {
                client = elasticsearch_util_1.ElasticsearchUtil.getClient();
                const result = yield client.search({
                    index: "products",
                    query: {
                        bool: {
                            should: [
                                { match: { name: keyword } },
                                { match: { description: keyword } }
                            ]
                        }
                    }
                });
                let searchResponses = [];
                const hits = result.hits.hits;
                if (hits.length < 1) {
                    const errorMessage = "cannot find data with keyword: " + keyword;
                    throw new response_exception_1.ResponseException(404, (0, error_message_1.setErrorMessages)(errorMessage), errorMessage);
                }
                hits.forEach((hit) => {
                    const searchResponse = {
                        id: hit._source.id,
                        userId: hit._source.userId,
                        name: hit._source.name,
                        description: hit._source.description,
                        stock: hit._source.stock
                    };
                    searchResponses.push(searchResponse);
                });
                return Promise.resolve(searchResponses);
            }
            catch (e) {
                (0, error_exception_1.errorHandler)(e, requestId);
                return Promise.reject(e);
            }
        });
    }
}
exports.SearchProductService = SearchProductService;
