const createTableUser = async (connection) => {
    try {
        const query = `CREATE TABLE users (
  	        id int NOT NULL AUTO_INCREMENT,
  	        username varchar(50) NOT NULL,
  	        email varchar(100) NOT NULL,
  	        PRIMARY KEY (id),
  	        UNIQUE KEY username_unique_index (username),
  	        UNIQUE KEY email_unique_index (email)
        ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`
        await connection.execute(query)
    } catch(e) {
        console.log("error when creating user table", e);
    }
}

const createDataUser = async(connection) => {
    try {
        const query = `INSERT INTO users (id,username,email) VALUES (1,"username","email@email.com"), (2,"username2","email2@email.com"), (3,"username3","email3@email.com");`
        await connection.execute(query)
    } catch(e) {
        console.log("error when creating user data", e);
    }
    // const query = `INSERT INTO user (id,username,email) VALUES (1,"username","email@email.com");`
}

const dropUserTable = async(connection) => {
    try {
        const query = `DROP TABLE IF EXISTS users;`
        await connection.execute(query)
    } catch(e) {
        console.log("error when dropping user table", e);
    }
}

export default {
    createTableUser,
    createDataUser,
    dropUserTable
}