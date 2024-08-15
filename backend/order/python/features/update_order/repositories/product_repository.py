from features.update_order.models.product import Product

class ProductRepository:

    @classmethod
    def find_in_id(self, cursor, ids) -> list[Product]:
        placeholders = ""
        params = []
        for id in ids:
            placeholders += "%s,"
            params.append(id)
        placeholders = placeholders[:-1]
        query = "SELECT id, name, description, price FROM products WHERE id IN ("+placeholders+")"
        cursor.execute(query, params)
        products = []
        rows = cursor.fetchall()
        if rows is not None:
            for row in rows:
                products.append(Product(row[0], 0, row[1], row[2], row[3], 0))
        return products
