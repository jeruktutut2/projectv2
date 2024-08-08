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
exports.UpdateProductByIdService = void 0;
const mysql_utils_1 = require("../../../utils/mysql-utils");
const error_exception_1 = require("../../../exceptions/error-exception");
const update_product_by_id_repository_1 = require("../repositories/update-product-by-id-repository");
const response_exception_1 = require("../../../exceptions/response-exception");
const elasticsearch_util_1 = require("../../../utils/elasticsearch-util");
const validation_1 = require("../../../validation/validation");
const update_product_by_id_validation_schema_1 = require("../validation-schema/update-product-by-id-validation-schema");
const error_message_1 = require("../../../helpers/error-message");
class UpdateProductByIdService {
    static updateProductById(requestId, updateProductByIdRequest) {
        return __awaiter(this, void 0, void 0, function* () {
            let poolConnection = null;
            try {
                updateProductByIdRequest = validation_1.Validation.validate(update_product_by_id_validation_schema_1.UpdateProductByIdValidationSchema.UPDATE, updateProductByIdRequest);
                poolConnection = yield mysql_utils_1.MysqlUtil.getPool().getConnection();
                yield poolConnection.beginTransaction();
                const [rows] = yield update_product_by_id_repository_1.UpdateProductByIdRepository.getById(poolConnection, updateProductByIdRequest.id);
                if (rows.length !== 1) {
                    const errorMessage = "cannot find product with id:" + updateProductByIdRequest.id.toString();
                    throw new response_exception_1.ResponseException(400, (0, error_message_1.setErrorMessages)(errorMessage), errorMessage);
                }
                const product = {
                    id: updateProductByIdRequest.id,
                    name: updateProductByIdRequest.name,
                    description: updateProductByIdRequest.description
                };
                const [resultSetHeader] = yield update_product_by_id_repository_1.UpdateProductByIdRepository.updateNameAndDescriptionById(poolConnection, product);
                if (resultSetHeader.affectedRows !== 1) {
                    const errorMessage = "number of affected rows when creating product is not one:" + resultSetHeader.affectedRows.toString();
                    throw new response_exception_1.ResponseException(500, (0, error_message_1.setErrorMessages)(errorMessage), errorMessage);
                }
                yield elasticsearch_util_1.ElasticsearchUtil.getClient().update({
                    index: "products",
                    id: updateProductByIdRequest.id.toString(),
                    doc: {
                        name: updateProductByIdRequest.name,
                        description: updateProductByIdRequest.description
                    }
                });
                yield poolConnection.commit();
                const response = {
                    name: updateProductByIdRequest.name,
                    description: updateProductByIdRequest.description
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
exports.UpdateProductByIdService = UpdateProductByIdService;
