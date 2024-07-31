import express from "express";
import { SearchProductController } from "../controllers/service-product-controller";

export const searchProductRoute = express.Router()
searchProductRoute.post("/", SearchProductController.searchProduct)