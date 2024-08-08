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
exports.CreateProductService = void 0;
const error_exception_1 = require("../../../exceptions/error-exception");
const validation_1 = require("../../../validation/validation");
const create_product_validation_schema_1 = require("../validation-schema/create-product-validation-schema");
const mysql_utils_1 = require("../../../utils/mysql-utils");
const create_product_repository_1 = require("../repositories/create-product-repository");
const response_exception_1 = require("../../../exceptions/response-exception");
const elasticsearch_util_1 = require("../../../utils/elasticsearch-util");
const error_message_1 = require("../../../helpers/error-message");
class CreateProductService {
    static create(requestId, createProductRequest) {
        return __awaiter(this, void 0, void 0, function* () {
            let poolConnection = null;
            try {
                createProductRequest = validation_1.Validation.validate(create_product_validation_schema_1.CreateProductValidationSchema.CREATE, createProductRequest);
                poolConnection = yield mysql_utils_1.MysqlUtil.getPool().getConnection();
                yield poolConnection.beginTransaction();
                const product = {
                    userId: createProductRequest.userId,
                    name: createProductRequest.name,
                    description: createProductRequest.description,
                    stock: createProductRequest.stock
                };
                const [resultSetHeader] = yield create_product_repository_1.CreateProductRepository.create(poolConnection, product);
                if (resultSetHeader.affectedRows !== 1) {
                    const errorMessage = "number of affected rows when creating product is not one:" + resultSetHeader.affectedRows.toString();
                    throw new response_exception_1.ResponseException(500, (0, error_message_1.setErrorMessages)(errorMessage), errorMessage);
                }
                yield elasticsearch_util_1.ElasticsearchUtil.getClient().index({
                    index: "products",
                    id: resultSetHeader.insertId.toString(),
                    document: {
                        id: resultSetHeader.insertId.toString(),
                        userId: product.userId.toString(),
                        name: product.name,
                        description: product.description
                    }
                });
                yield poolConnection.commit();
                const response = {
                    name: product.name,
                    description: product.description,
                    stock: product.stock
                };
                return Promise.resolve(response);
            }
            catch (e) {
                if (poolConnection) {
                    yield poolConnection.rollback();
                }
                (0, error_exception_1.errorHandler)(e, requestId);
                return Promise.reject(e);
            }
            finally {
                if (poolConnection) {
                    poolConnection.release();
                }
            }
        });
    }
}
exports.CreateProductService = CreateProductService;
