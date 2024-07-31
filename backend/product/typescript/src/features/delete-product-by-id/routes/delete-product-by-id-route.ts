import express from "express";
import { DeleteProductByIdController } from "../controllers/delete-product-by-id-controller";
import { authenticate } from "../../../middlewares/auth_middleware";

export const deleteProductByIdRouter = express.Router()
deleteProductByIdRouter.post("/", DeleteProductByIdController.deleteProductById, authenticate)