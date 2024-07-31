import { errorHandler } from "../../../commons/exceptions/error-exception";
import mysqlUtil from "../../../commons/utils/mysql-util.js";
import getCartRepository from "../repositories/get-cart-repository.js";

const getCart = async(requestId, userId) => {
    let connection
    try {
        connection = await mysqlUtil.getConnection(mysqlUtil.mysqlPool)
        const [rows] = await getCartRepository.getCartByUserId(connection, userId)
        // console.log("rows:", rows);
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