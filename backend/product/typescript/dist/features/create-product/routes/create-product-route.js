"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.createProductRouter = void 0;
const express_1 = __importDefault(require("express"));
const create_product_controller_1 = require("../controllers/create-product-controller");
const auth_middleware_1 = require("../../../middlewares/auth_middleware");
exports.createProductRouter = express_1.default.Router();
exports.createProductRouter.post("/api/v1/products", create_product_controller_1.CreateProductController.createProduct, auth_middleware_1.authenticate);
