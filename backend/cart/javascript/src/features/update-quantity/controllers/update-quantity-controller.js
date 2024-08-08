import updateQuantityService from "../services/update-quantity-service.js";
import responseMessageHelper from "../../../commons/helpers/response-message-helper.js";
import { errorHandlerResponse } from "../../../commons/exceptions/error-exception.js";
const updateQuantity = async(req, res, next) => {
    const requestId = req.get("X-REQUEST-ID")
    try {
        const updateQuantityResponse = await updateQuantityService.updateQuantity(requestId, req.body)
        const response = responseMessageHelper.setResponse(updateQuantityResponse, null)
        return res.status(200).json(response)
    } catch(e) {
        return errorHandlerResponse(res, e)
    }
}

export default {
    updateQuantity
}