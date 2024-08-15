class OrderItemRepository:
    
    @classmethod
    def create_many(self, cursor, order_items):
        query = "INSERT INTO order_items(order_id, product_id, price, quantity, total) VALUES (%s, %s, %s, %s, %s);"
        params = []
        for order_item in order_items:
            param = (order_item.order_id, order_item.product_id, order_item.price, order_item.quantity, order_item.total)
            params.append(param)
        return cursor.executemany(query, params)
