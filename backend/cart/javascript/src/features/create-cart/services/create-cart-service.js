import mysqlUtil from "../../../commons/utils/mysql-util.js";
import { validate } from "../../../commons/validations/validation.js";
import { createCartValidation } from "../schema-validation/create-cart-schema-validation.js";
import createCartRepository from "../repositories/create-cart-repository.js";
import { errorHandler, setInternalServerErrorMessage } from "../../../commons/exceptions/error-exception.js";
import { ResponseException } from "../../../commons/exceptions/response-exception.js";

const createCart = async (requestId, request) => {
    let connection
    try {
        validate(createCartValidation, request)

        connection = await mysqlUtil.getConnection(mysqlUtil.mysqlPool)

        const cart = {
            userId: request.userId,
            productId: request.productId,
            quantity: request.quantity
        }
        const result = await createCartRepository.createCart(connection, cart)
        if (result[0].affectedRows !== 1) {
            console.log(JSON.stringify({logTime: (new Date()).toISOString(), app: "backend-cart-nodejs", requestId: requestId, error: "affected rows create cart not one: " + result[0].affectedRows}));
            throw new ResponseException(500, setInternalServerErrorMessage(), "internal server error")
        }
        return cart
    } catch(e) {
        errorHandler(e, requestId)
    } finally {
        if (connection) {
            await mysqlUtil.releaseConnection(mysqlUtil.mysqlPool, connection)
        }
    }
}

export default {
    createCart
}