class OrderItemRepository:

    @classmethod
    def delete_by_order_id(self, cursor, order_id):
        query = "DELETE FROM order_items WHERE order_id = %s;"
        params = (order_id,)
        return cursor.execute(query, params)