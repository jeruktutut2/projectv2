from features.update_order.models.order import Order
class OrderRepository:

    @classmethod
    def find_by_id(self, cursor, id) -> Order:
        query = "SELECT id, user_id, total, paid, created_at FROM orders WHERE id = %s FOR UPDATE;"
        params = (id,)
        cursor.execute(query, params)

        order = None
        row = cursor.fetchone()
        if row is not None:
            order = Order(row[0], row[1], row[2], row[3], row[4])
        return order
    
    @classmethod
    def update_total(self, cursor, total, id):
        query = "UPDATE orders SET total = %s WHERE id = %s;"
        params = (total, id)
        return cursor.execute(query, params)