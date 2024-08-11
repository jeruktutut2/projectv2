from decimal import Decimal
class Product:
    id: int
    user_id: int
    name: str
    description: str
    price: Decimal
    stock: int
    def __init__(self, id, user_id, name, description, price, stock):
        self.id = id
        self.user_id = user_id
        self.name = name
        self.description = description
        self.price = price
        self.stock = stock
    
    def __str__(self):
        return f"id: {self.id}, user_id: {self.user_id}, name: {self.name}, description: {self.description}, price: {self.price}, stock: {self.stock} "
