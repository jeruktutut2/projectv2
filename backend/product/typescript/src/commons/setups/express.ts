import express from "express";
import { printRequestResponseLog } from "../middlewares/log-request-response-middleware";
import { createProductRouter } from "../../features/create-product/routes/create-product-route";
import { deleteProductByIdRouter } from "../../features/delete-product-by-id/routes/delete-product-by-id-route";
import { getProductByIdRouter } from "../../features/get-product-by-id/routes/get-product-by-id-route";
import { searchProductRoute } from "../../features/search-product/routes/search-product-route";
import { updateProductByIdRoute } from "../../features/update-product-by-id/routes/update-product-by-id-route";

export const web = express()
web.use(express.json())
web.use(printRequestResponseLog)
web.use(createProductRouter)
web.use(deleteProductByIdRouter)
web.use(getProductByIdRouter)
web.use(searchProductRoute)
web.use(updateProductByIdRoute)


