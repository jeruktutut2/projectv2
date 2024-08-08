"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.getProductByIdRouter = void 0;
const express_1 = __importDefault(require("express"));
const get_product_by_id_controller_1 = require("../controllers/get-product-by-id-controller");
exports.getProductByIdRouter = express_1.default.Router();
exports.getProductByIdRouter.get("/api/v1/products", get_product_by_id_controller_1.GetProductByIdController.getProductById);
