import express from "express";
import { CreateProductController } from "../controllers/create-product-controller";
import { authenticate } from "../../../commons/middlewares/auth_middleware";

export const createProductRouter = express.Router()
createProductRouter.post("/api/v1/products", CreateProductController.createProduct, authenticate)