const getByUserIdAndProductId = async(connection, userId, productId) => {
    const query = `SELECT id, user_id, product_id, quantity FROM carts WHERE user_id = ? AND product_id = ?;`
    const params = [userId, productId]
    return await connection.execute(query, params)
}

const updateQuantityById = async(connection, quantity, id) => {
    const query = `UPDATE carts SET quantity = ? WHERE id = ?;`
    const params = [quantity, id]
    return await connection.execute(query, params)
}

export default {
    getByUserIdAndProductId,
    updateQuantityById
}