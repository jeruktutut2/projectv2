class OrderRepository:
    
    @classmethod
    def create(self, cursor, order):
        query = "INSERT INTO orders(user_id, total, paid, created_at) VALUES (%s, %s, %s, %s);"
        params = (order.user_id, order.total, order.paid, order.created_at)
        return cursor.execute(query, params)