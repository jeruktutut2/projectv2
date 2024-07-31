import express from "express";
import { CreateProductController } from "../controllers/create-product-controller";
import { authenticate } from "../../../middlewares/auth_middleware";

export const createProductRouter = express.Router()
createProductRouter.post("/", CreateProductController.createProduct, authenticate)