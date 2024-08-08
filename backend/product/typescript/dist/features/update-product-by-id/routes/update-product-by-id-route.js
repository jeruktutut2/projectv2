"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.updateProductByIdRoute = void 0;
const express_1 = __importDefault(require("express"));
const update_product_by_id_controller_1 = require("../controllers/update-product-by-id-controller");
const auth_middleware_1 = require("../../../middlewares/auth_middleware");
exports.updateProductByIdRoute = express_1.default.Router();
exports.updateProductByIdRoute.patch("/api/v1/products", update_product_by_id_controller_1.UpdateProductByIdController.updateProductById, auth_middleware_1.authenticate);
