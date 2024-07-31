const getCartByIdAndUserIdAndProductId = async (connection, id, userId, productId) => {
    const query = `SELECT id, user_id, product_id, quantity FROM carts WHERE id = ? AND user_id = ? AND product_id = ?;`
    const params = [id, userId, productId]
    return await connection.execute(query, params)
}

const deleteCartById = async(connection, id) => {
    const query = `DELETE FROM carts WHERE id = ?;`
    const params = [id]
    return await connection.execute(query, params)
}

export default {
    getCartByIdAndUserIdAndProductId,
    deleteCartById
}