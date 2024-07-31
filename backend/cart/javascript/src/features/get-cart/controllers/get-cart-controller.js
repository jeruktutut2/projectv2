import getCartService from "../services/get-cart-service.js";
import responseMessageHelper from "../../../commons/helpers/response-message-helper.js";
import { errorHandlerResponse } from "../../../commons/exceptions/error-exception";
const getCart = async(req, res, next) => {
    const requestId = req.get("X-REQUEST-ID")
    try {
        // req.get("X-REQUEST-ID")
        const getCartResponse = await getCartService.getCart(requestId, req.body)
        const response = responseMessageHelper.setResponse(getCartResponse, null)
        return res.status(200).json(response)
    } catch(e) {
        return errorHandlerResponse(res, e)
    }
}

export default {
    getCart
}