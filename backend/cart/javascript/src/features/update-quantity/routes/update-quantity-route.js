import express from "express";
import updateQuantityController from "../controllers/update-quantity-controller.js";

const updateQuantityRouter = express.Router()
updateQuantityRouter.patch("/api/v1/carts/update-quantity", updateQuantityController.updateQuantity)

export default {
    updateQuantityRouter
}