"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.deleteProductByIdRouter = void 0;
const express_1 = __importDefault(require("express"));
const delete_product_by_id_controller_1 = require("../controllers/delete-product-by-id-controller");
const auth_middleware_1 = require("../../../middlewares/auth_middleware");
exports.deleteProductByIdRouter = express_1.default.Router();
exports.deleteProductByIdRouter.delete("/api/v1/products", delete_product_by_id_controller_1.DeleteProductByIdController.deleteProductById, auth_middleware_1.authenticate);
