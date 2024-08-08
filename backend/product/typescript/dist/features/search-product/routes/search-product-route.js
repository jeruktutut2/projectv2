"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.searchProductRoute = void 0;
const express_1 = __importDefault(require("express"));
const service_product_controller_1 = require("../controllers/service-product-controller");
exports.searchProductRoute = express_1.default.Router();
exports.searchProductRoute.get("/api/v1/products/search", service_product_controller_1.SearchProductController.searchProduct);
