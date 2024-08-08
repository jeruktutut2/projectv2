import createCartController from "../services/create-cart-service.js";
import responseMessageHelper from "../../../commons/helpers/response-message-helper.js";
import { errorHandlerResponse } from "../../../commons/exceptions/error-exception.js";
const createCart = async(req, res, next) => {
    const requestId = req.get("X-REQUEST-ID")
    try {
        const createCartResponse = await createCartController.createCart(requestId, req.body)
        const response = responseMessageHelper.setResponse(createCartResponse, null)
        return res.status(201).json(response)
    } catch(e) {
        return errorHandlerResponse(res, e)
    }
}

export default {
    createCart
}