from decimal import Decimal

class OrderItem:
    id: int
    order_id: int
    product_id: int
    price: Decimal
    quantity: int
    total: Decimal

    def __init__(self, id, order_id, product_id, price, quantity, total):
        self.id = id
        self.order_id = order_id
        self.product_id = product_id
        self.price = price
        self.quantity = quantity
        self.total = total