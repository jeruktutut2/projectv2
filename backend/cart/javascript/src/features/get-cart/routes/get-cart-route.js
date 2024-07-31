import express from "express";
import getCartController from "../controllers/get-cart-controller.js";

const getCartRouter = express.Router()
getCartRouter.get("/api/v1/carts", getCartController.getCart)

export default {
    getCartRouter
}