import express from "express";
import logRequestResponseMiddleware from "../middlewares/log-request-response-middleware.js";
import createCartRouter from "../../features/create-cart/routes/create-cart-route.js";
import deleteCartRouter from "../../features/delete-cart/routes/delete-cart-route.js";
import getCartRouter from "../../features/get-cart/routes/get-cart-route.js";
import updateQuantityRouter from "../../features/update-quantity/routes/update-quantity-route.js";

export const web = express()
web.use(express.json())
web.use(logRequestResponseMiddleware.printLogRequestResponse())
web.use(createCartRouter.createCartRouter)
web.use(deleteCartRouter.deleteCartRouter)
web.use(getCartRouter.getCartRouter)
web.use(updateQuantityRouter.updateQuantityRouter)
