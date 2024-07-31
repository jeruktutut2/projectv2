const createCart = async (connection, cart) => {
    const query = `INSERT INTO carts (user_id, product_id, quantity) VALUES (?, ?, ?);`
    const params = [cart.userId, cart.productId, cart.quantity]
    return await connection.execute(query, params)
}

export default {
    createCart
}