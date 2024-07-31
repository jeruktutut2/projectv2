import express from "express";
import createCartControlle from "../controllers/create-cart-controller.js";

const createCartRouter = express.Router()
createCartRouter.post("/api/v1/carts", createCartControlle.createCart)

export default {
    createCartRouter
}