import express from "express";
import { GetProductByIdController } from "../controllers/get-product-by-id-controller";

export const getProductByIdRouter = express.Router()
getProductByIdRouter.post("/", GetProductByIdController.getProductById)