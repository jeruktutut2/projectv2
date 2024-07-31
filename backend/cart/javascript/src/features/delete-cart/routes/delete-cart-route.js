import express from "express";
import deleteCartController from "../controllers/delete-cart-controller.js";

const deleteCartRouter = express.Router()
deleteCartRouter.post("", deleteCartController.deleteCart)

export default {
    deleteCartRouter
}