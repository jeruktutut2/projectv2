import { errorHandler } from "../../../commons/exceptions/error-exception.js";
import mysqlUtil from "../../../commons/utils/mysql-util.js";
import { validate } from "../../../commons/validations/validation.js";
import getCartRepository from "../repositories/get-cart-repository.js";
import { getCartValidation } from "../schema-validation/get-cart-schema-validation.js";

const getCart = async(requestId, request) => {
    let connection
    try {
        validate(getCartValidation, request)
        connection = await mysqlUtil.getConnection(mysqlUtil.mysqlPool)
        const [rows] = await getCartRepository.getCartByUserId(connection, request.userId)
        const cart = []
        rows.forEach(element => {
            cart.push({userId: element.user_id, productId: element.product_id, quantity: element.quantity})
        });
        return cart
    } catch(e) {
        errorHandler(e, requestId)
    } finally {
        if (connection) {
            mysqlUtil.releaseConnection(mysqlUtil.mysqlPool, connection)
        }
    }
}

export default {
    getCart
}