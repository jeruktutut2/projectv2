import { updateQuantityValidation } from "../schema-validation/update-quantity-schema-valiation.js";
import { validate } from "../../../commons/validations/validation.js";
import mysqlUtil from "../../../commons/utils/mysql-util.js";
import { errorHandler, setErrorMessages, setInternalServerErrorMessage } from "../../../commons/exceptions/error-exception.js";
import updateQuantityRepository from "../repositories/update-quantity-repository.js";
import { ResponseException } from "../../../commons/exceptions/response-exception.js";

const updateQuantity = async(requestId, request) => {
    let connection
    try {
        validate(updateQuantityValidation, request)
        connection = await mysqlUtil.getConnection(mysqlUtil.mysqlPool)
        await connection.beginTransaction()
        // this should be by idCart, userId and productId
        const [rows] = await updateQuantityRepository.getByUserIdAndProductId(connection, request.userId, request.productId)
        console.log("rows:", rows);
        if (rows.length !== 1) {
            throw new ResponseException(400, setErrorMessages("cannot find product in cart or many same product in cart"), "cannot find product in cart")
        }
        const result = await updateQuantityRepository.updateQuantityById(connection, request.quantity, rows[0].id)
        if (result[0].affectedRows !== 1) {
            console.log(JSON.stringify({logTime: (new Date()).toISOString(), app: "backend-cart-nodejs", requestId: requestId, error: "affected rows update quantity not one: " + result[0].affectedRows}));
            throw new ResponseException(500, setInternalServerErrorMessage(), "internal server error")
        }
        await connection.commit()
        return {userId: request.userId, productId: request.productId, quantity: request.quantity}
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
    updateQuantity
}