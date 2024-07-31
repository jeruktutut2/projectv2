import express from "express";
import { SearchProductController } from "../controllers/service-product-controller";

export const searchProductRoute = express.Router()
searchProductRoute.get("/api/v1/products/search", SearchProductController.searchProduct)