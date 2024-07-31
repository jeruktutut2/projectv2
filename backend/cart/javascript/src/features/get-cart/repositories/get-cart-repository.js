const getCartByUserId = async(connection, userId) => {
    const query = `SELECT id, user_id, product_id, quantity FROM carts WHERE user_id = ?;`
    const params = [userId]
    return await connection.execute(query, params)
}

export default {
    getCartByUserId
}