import express from "express";
import createCartControlle from "../controllers/create-cart-controller.js";

const createCartRouter = express.Router()
createCartRouter.post("", createCartControlle.createCart)

export default {
    createCartRouter
}