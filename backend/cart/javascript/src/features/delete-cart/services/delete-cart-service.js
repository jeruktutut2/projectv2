import { errorHandler, setErrorMessages, setInternalServerErrorMessage } from "../../../commons/exceptions/error-exception.js";
import { ResponseException } from "../../../commons/exceptions/response-exception.js";
import mysqlUtil from "../../../commons/utils/mysql-util.js";
import deleteCartRepository from "../repositories/delete-cart-repository.js";
import { validate } from "../../../commons/validations/validation.js";
import { deleteCartValidation } from "../schema-validation/delete-cart-schema-validation.js";
import { setDataMessage } from "../../../commons/helpers/data-message-helper.js";

const deleteCart = async(requestId, request) => {
    let connection
    try {
        validate(deleteCartValidation, request)
        connection = await mysqlUtil.getConnection(mysqlUtil.mysqlPool)
        await connection.beginTransaction()
        const [rows] = await deleteCartRepository.getCartByIdAndUserIdAndProductId(connection, request.id, request.userId, request.productId)
        if (rows.length !== 1) {
            throw new ResponseException(400, setErrorMessages("cannot find product in cart"), "cannot find product in cart")
        }
        const result = await deleteCartRepository.deleteCartById(connection, request.id)
        if (result[0].affectedRows !== 1) {
            console.log(JSON.stringify({logTime: (new Date()).toISOString(), app: "backend-cart-nodejs", requestId: requestId, error: "affected rows delete cart not one: " + result[0].affectedRows}));
            throw new ResponseException(500, setInternalServerErrorMessage(), "internal server error")
        }
        await connection.commit()
        const dataMessage = setDataMessage("successfully delete cart")
        return dataMessage
    } catch(e) {
        if (connection) {
            await connection.rollback()
        }
        errorHandler(e, requestId)
    } finally {
        if (connection) {
            mysqlUtil.releaseConnection(mysqlUtil.mysqlPool, connection)
        }
    }
}

export default {
    deleteCart
}