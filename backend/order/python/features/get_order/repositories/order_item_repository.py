from features.get_order.models.order_item import OrderItem

class OrderItemsRepository:

    @classmethod
    def find_by_order_id(self, cursor, order_id) -> list[OrderItem]:
        query = "SELECT id, order_id, product_id, price, quantity, total FROM order_items WHERE order_id = %s;"
        params = (order_id,)
        cursor.execute(query, params)
        order_items = []
        rows = cursor.fetchall()
        if cursor.fetchall() is not None:
            for row in rows:
                order_items.append(OrderItem(row[0], row[1], row[2], row[3], row[4], row[5]))
        
        return order_items