class OrderRepository:

    @classmethod
    def delete_by_id(self, cursor, id):
        query = "DELETE FROM orders WHERE id = %s;"
        params = (id,)
        return cursor.execute(query, params)