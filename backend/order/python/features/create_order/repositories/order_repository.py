class OrderRepository:
    # @staticmethod
    @classmethod
    def create(self, cursor, order):
        query = "INSERT INTO orders(user_id, total, paid, created_at) VALUES (%s, %s, %s, %s);"
        params = (order.user_id, order.total, order.paid, order.created_at)
        return cursor.execute(query, params)
    
    # def create_order_items(self, cursor, order_items):
    #     query = "INSERT INTO order_items(order_id, product_id, price, quantity, total) VALUES (%s, %s, %s, %s, %s);"
    #     params = []
    #     for order_item in order_items:
    #         param = (order_item.order_id, order_item.product_id, order_item.price, order_item.quantity, order_item.total)
    #         params.append(param)
    #     # params = []
    #     return cursor.executemany(query, params)