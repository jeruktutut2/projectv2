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
exports.GetProductByIdService = void 0;
const error_exception_1 = require("../../../exceptions/error-exception");
const mysql_utils_1 = require("../../../utils/mysql-utils");
const get_product_by_id_repository_1 = require("../repositories/get-product-by-id-repository");
const response_exception_1 = require("../../../exceptions/response-exception");
const error_message_1 = require("../../../helpers/error-message");
class GetProductByIdService {
    static getProductById(requestId, id) {
        return __awaiter(this, void 0, void 0, function* () {
            let poolConnection = null;
            try {
                poolConnection = yield mysql_utils_1.MysqlUtil.getPool().getConnection();
                yield poolConnection.beginTransaction();
                const [rows] = yield get_product_by_id_repository_1.GetProductByIdRepository.getProductById(poolConnection, id);
                if (rows.length !== 1) {
                    const errorMessage = "cannot find product with id:" + id.toString();
                    throw new response_exception_1.ResponseException(400, (0, error_message_1.setErrorMessages)(errorMessage), errorMessage);
                }
                yield poolConnection.commit();
                const response = {
                    id: rows[0].id,
                    name: rows[0].name,
                    description: rows[0].description,
                    stoct: rows[0].stock
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
exports.GetProductByIdService = GetProductByIdService;
