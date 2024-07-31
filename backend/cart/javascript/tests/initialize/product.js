const createTableProduct = async(connection) => {
    try {
        const query = `CREATE TABLE products (
  	        id int NOT NULL AUTO_INCREMENT,
	        user_id int NOT NULL,
	        name varchar(255) NOT NULL,
	        description TEXT NOT NULL,
	        stock MEDIUMINT NOT NULL,
  	        PRIMARY KEY (id),
  	        -- KEY fk_user_permission_user (user_id),
  	        -- CONSTRAINT user_permission_ibfk_1 FOREIGN KEY (user_id) REFERENCES user (id)
            FOREIGN KEY fk_products_users(user_id) REFERENCES users(id)
        ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`
        await connection.execute(query)
    } catch(e) {
        console.log("error when creating product table", e);
    }
}

const createDataProduct = async(connection) => {
    try {
        const query = `INSERT INTO products (id, user_id, name, description, stock) VALUES (1, 1, "name", "description", 1), (2, 2, "name2", "description2", 2), (3, 3, "name3", "description3", 3);`
        await connection.execute(query)
    } catch(e) {
        console.log("error when creating product data");
    }
}

const dropTableProduct = async(connection) => {
    try {
        const query = `DROP TABLE IF EXISTS products;`
        await connection.execute(query)
    } catch(e) {
        console.log("error when dropping product table", e);
    }
}

export default {
    createTableProduct,
    createDataProduct,
    dropTableProduct
}