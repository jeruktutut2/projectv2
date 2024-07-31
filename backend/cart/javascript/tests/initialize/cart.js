const createTableCart = async (connection) => {
    try {
        const query = `CREATE TABLE carts (
            id INT NOT NULL AUTO_INCREMENT,
            user_id INT NOT NULL,
            product_id INT NOT NULL,
	        quantity INT NOT NULL,
            PRIMARY KEY (id),
            FOREIGN KEY fk_carts_user(user_id) REFERENCES users(id),
            FOREIGN KEY fk_carts_product(product_id) REFERENCES products(id)
        ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`
        await connection.execute(query)
    } catch(e) {
        console.log("error when creating cart table", e);
    }
}

const createDataCart = async(connection) => {
    try {
        const query = `INSERT INTO carts (user_id, product_id, quantity) VALUES (1, 1, 1), (2, 2, 2), (3, 3, 3), (1, 2, 2);`
        await connection.execute(query)
    } catch(e) {
        console.log("error when creating cart data", e);
    }
}

const dropTableCart = async(connection) => {
    try {
        const query = `DROP TABLE IF EXISTS carts;`
        await connection.execute(query)
    } catch(e) {
        console.log("error when dropping cart table", e);
    }
}

const getDataCart = async(connection, id) => {
    try {
        const query = `SELECT id, user_id, product_id, quantity FROM carts WHERE id = ?;`
        const params = [id]
        return await connection.execute(query, params)
    } catch(e) {
        console.log("error when getting cart data", e);
    }
}

export default {
    createTableCart,
    createDataCart,
    dropTableCart,
    getDataCart
}