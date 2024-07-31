import express from "express";
import getCartController from "../controllers/get-cart-controller.js";

const getCartRouter = express.Router()
getCartRouter.post("", getCartController.getCart)

export default {
    getCartRouter
}