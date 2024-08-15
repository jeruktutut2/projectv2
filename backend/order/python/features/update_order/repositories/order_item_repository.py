from features.update_order.models.order_item import OrderItem

class OrderItemRepository:
    
    @classmethod
    def create(self, cursor, order_items):
        query = "INSERT INTO order_items(order_id, product_id, price, quantity, total) VALUES (%s, %s, %s, %s, %s);"
        params = []
        for order_item in order_items:
            param = (order_item.order_id, order_item.product_id, order_item.price, order_item.quantity, order_item.total)
            params.append(param)
        return cursor.executemany(query, params)
    
    @classmethod
    def delete_by_order_id(self, cursor, order_id):
        query = "DELETE FROM order_items WHERE order_id = %s;"
        params = (order_id,)
        return cursor.execute(query, params)
    
    @classmethod
    def find_by_order_id(self, cursor, order_id) -> list[OrderItem]:
        query = "SELECT id, order_id, product_id, price, quantity, total FROM order_items WHERE order_id = %s;"
        params = (order_id,)
        cursor.execute(query, params)
        order_items = []
        rows = cursor.fetchall()
        if rows is not None:
            for row in rows:
                order_items.append(OrderItem(row[0], row[1], row[2], row[3], row[4], row[5]))
        return order_items