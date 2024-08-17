def create_table_orders(cursor):
    try:
        query = """
            CREATE TABLE orders (
  	            id int NOT NULL AUTO_INCREMENT,
	            user_id int NOT NULL,
                total DECIMAL(18,2) NOT NULL,
	            paid TINYINT NOT NULL,
                created_at BIGINT NOT NULL,
  	            PRIMARY KEY (id),
	            FOREIGN KEY fk_orders_users(user_id) REFERENCES users(id)
            ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
        """
        cursor.execute(query)
        print("successfully create table orders")
    except Exception as e:
        print("error when creating table orders", e)

def create_data_orders(cursor):
    try:
        query = "INSERT INTO orders(id, user_id, total, paid, created_at) VALUES (1, 1, 10, 0, 1722390867657), (2, 2, 20, 0, 1722390867657), (3, 3, 30, 0, 1722390867657);"
        cursor.execute(query)
        if cursor.rowcount == 3:
            print("successfully create data orders")
    except Exception as e:
        print("error when creating data orders", e)

def get_data_orders(cursor):
    query = "SELECT id, user_id, total, paid, created_at FROM orders WHERE id = 1"
    cursor.execute(query)
    return cursor.fetchall()

def delete_table_orders(cursor):
    try:
        query = "DROP TABLE IF EXISTS orders;"
        cursor.execute(query)
        print("successfully delete table orders")
    except Exception as e:
        print("error when deleting table orders", e)