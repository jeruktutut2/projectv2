import express from "express";
import deleteCartController from "../controllers/delete-cart-controller.js";

const deleteCartRouter = express.Router()
deleteCartRouter.delete("/api/v1/carts", deleteCartController.deleteCart)

export default {
    deleteCartRouter
}