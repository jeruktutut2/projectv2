from features.create_order.models.product import Product
class ProductRepository:

    # @staticmethod
    @classmethod
    def find_in_id(self, cursor, ids):
        # print(cursor, ids)
        placeholders = ""
        params = []
        for id in ids:
            placeholders += "%s,"
            params.append(id)

        placeholders = placeholders[:-1]
        query = "SELECT id, name, description, price FROM products WHERE id IN ("+placeholders+")"
        cursor.execute(query, params)
        # print("cursor.fetchall():", cursor.fetchall())
        # return cursor.execute(query, params)
        products = []
        for product in cursor.fetchall():
            products.append(Product(product[0], 0, product[1], product[2], product[3], 0))
        # return cursor.fetchall()
        return products