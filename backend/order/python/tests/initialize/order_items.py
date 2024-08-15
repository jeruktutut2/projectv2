def create_table_order_items(cursor):
    try:
        query = """
        CREATE TABLE order_items (
  	        id int NOT NULL AUTO_INCREMENT,
            order_id INT NOT NULL,
            product_id INT NOT NULL,
            price DECIMAL(18,2) NOT NULL,
            quantity TINYINT NOT NULL,
            total DECIMAL(18,2) NOT NULL,
  	        PRIMARY KEY (id),
	        FOREIGN KEY fk_order_items_orders(order_id) REFERENCES orders(id),
            FOREIGN KEY fk_order_items_products(product_id) REFERENCES products(id)
        ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
        """
        cursor.execute(query)
        # print("create_table_order_items cursor.rowcount", cursor.rowcount)
        # if cursor.rowcount == 1:
        print("successfully create table order_items")
    except Exception as e:
        print("error when creating table order_items", e)

def create_data_order_items(cursor):
    try:
        query = "INSERT INTO order_items(id, order_id, product_id, price, quantity, total) VALUES (1, 1, 1, 1, 1, 1), (2, 1, 2, 2, 2, 2), (3, 1, 3, 3, 3, 3);"
        cursor.execute(query)
        # print("create_data_order_items cursor.rowcount:", cursor.rowcount)
        if cursor.rowcount == 3:
            print("successfully create data order_items")
    except Exception as e:
        print("error when creating data order_items", e)

def get_data_order_items(cursor):
    query = "SELECT id, order_id, product_id, price, quantity, total FROM order_items WHERE order_id = 1"
    cursor.execute(query)
    return cursor.fetchall()

def delete_table_order_items(cursor):
    try:
        query = "DROP TABLE IF EXISTS order_items;"
        cursor.execute(query)
        # print("delete_table_order_items cursor.rowcount:", cursor, cursor.rowcount)
        # if cursor.rowcount == -1:
        print("successfully delete table order_items")
    except Exception as e:
        print("error when deleting table order_items", e)