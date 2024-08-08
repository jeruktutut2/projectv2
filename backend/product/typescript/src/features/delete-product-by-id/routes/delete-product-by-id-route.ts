import express from "express";
import { DeleteProductByIdController } from "../controllers/delete-product-by-id-controller";
import { authenticate } from "../../../commons/middlewares/auth_middleware";

export const deleteProductByIdRouter = express.Router()
deleteProductByIdRouter.delete("/api/v1/products", DeleteProductByIdController.deleteProductById, authenticate)