import express from "express";
import updateQuantityController from "../controllers/update-quantity-controller.js";

const updateQuantityRouter = express.Router()
updateQuantityRouter.post("", updateQuantityController.updateQuantity)

export default {
    updateQuantityRouter
}