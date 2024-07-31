import deleteCartService from "../services/delete-cart-service.js";
import responseMessageHelper from "../../../commons/helpers/response-message-helper.js";
import { errorHandlerResponse } from "../../../commons/exceptions/error-exception";
const deleteCart = async(req, res, next) => {
    const requestId = req.get("X-REQUEST-ID")
    try {
        // req.get("X-REQUEST-ID")
        const deleteCartResponse = await deleteCartService.deleteCart(requestId, req.body)
        const response = responseMessageHelper.setResponse(deleteCartResponse, null)
        return res.status(200).json(response)
    } catch(e) {
        return errorHandlerResponse(res, e)
    }
}

export default {
    deleteCart
}