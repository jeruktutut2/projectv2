from decimal import Decimal

class Order:
    id: int
    user_id: int
    total: Decimal
    paid: int
    created_at: int

    def __init__(self, id, user_id, total, paid, created_at):
        self.id = id
        self.user_id = user_id
        self.total = total
        self.paid = paid
        self.created_at = created_at