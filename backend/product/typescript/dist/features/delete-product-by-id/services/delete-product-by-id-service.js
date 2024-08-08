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
exports.DeleteProductByIdService = void 0;
const mysql_utils_1 = require("../../../utils/mysql-utils");
const error_exception_1 = require("../../../exceptions/error-exception");
const delete_product_by_id_repository_1 = require("../repositories/delete-product-by-id-repository");
const response_exception_1 = require("../../../exceptions/response-exception");
const elasticsearch_util_1 = require("../../../utils/elasticsearch-util");
const error_message_1 = require("../../../helpers/error-message");
const data_message_1 = require("../../../helpers/data-message");
class DeleteProductByIdService {
    static deleteProductById(requestId, id) {
        return __awaiter(this, void 0, void 0, function* () {
            let poolConnection = null;
            try {
                poolConnection = yield mysql_utils_1.MysqlUtil.getPool().getConnection();
                yield poolConnection.beginTransaction();
                const [resultSetHeader] = yield delete_product_by_id_repository_1.DeleteProductByIdRepository.DeleteProductById(poolConnection, id);
                if (resultSetHeader.affectedRows !== 1) {
                    const errorMessage = "number of affected rows when deleting product is not one:" + resultSetHeader.affectedRows.toString();
                    throw new response_exception_1.ResponseException(500, (0, error_message_1.setErrorMessages)(errorMessage), errorMessage);
                }
                yield elasticsearch_util_1.ElasticsearchUtil.getClient().delete({
                    index: "products",
                    id: id.toString()
                });
                yield poolConnection.commit();
                const dataMessage = (0, data_message_1.setDataMessage)("successfully delete product");
                return Promise.resolve(dataMessage);
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
exports.DeleteProductByIdService = DeleteProductByIdService;
