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
exports.GetProductByIdController = void 0;
const get_product_by_id_service_1 = require("../services/get-product-by-id-service");
const response_message_1 = require("../../../helpers/response-message");
const error_exception_1 = require("../../../exceptions/error-exception");
class GetProductByIdController {
    static getProductById(req, res, next) {
        return __awaiter(this, void 0, void 0, function* () {
            var _a;
            const requestId = (_a = req.get("X-REQUEST-ID")) !== null && _a !== void 0 ? _a : "";
            try {
                const getProductByIdRequest = req.body;
                const getProductByIdResponse = yield get_product_by_id_service_1.GetProductByIdService.getProductById(requestId, Number(getProductByIdRequest.id));
                const response = (0, response_message_1.setResponse)(getProductByIdResponse, null);
                return res.status(200).json(response);
            }
            catch (e) {
                return (0, error_exception_1.errorHandlerResponse)(res, e, requestId);
            }
        });
    }
}
exports.GetProductByIdController = GetProductByIdController;
