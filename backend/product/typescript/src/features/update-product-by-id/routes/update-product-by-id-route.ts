import express from "express";
import { UpdateProductByIdController } from "../controllers/update-product-by-id-controller";
import { authenticate } from "../../../middlewares/auth_middleware";

export const updateProductByIdRoute = express.Router()
updateProductByIdRoute.patch("/api/v1/products", UpdateProductByIdController.updateProductById, authenticate)