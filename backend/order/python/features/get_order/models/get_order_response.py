from decimal import Decimal
from pydantic import BaseModel, Field, ConfigDict

class OrderItemResponse(BaseModel):
    model_config = ConfigDict(
        extra="allow",
        title="Order Items Response",
        description="Represents order items response",
        arbitrary_types_allowed=True,
        populate_by_name=True
    )
    id: int = Field(..., alias="id")
    order_id: int = Field(..., alias="orderId")
    product_id: int = Field(..., alias="productId")
    price: Decimal = Field(..., alias="price")
    quantity: int = Field(..., alias="quantity")
    total: Decimal = Field(..., alias="total")

class GetOrderResponse(BaseModel):
    model_config = ConfigDict(
        extra="allow",
        title="Get Order Items Response",
        description="Represents get order items response",
        arbitrary_types_allowed=True,
        populate_by_name=True
    )
    id: int = Field(..., alias="id")
    user_id: int = Field(..., alias="userId")
    total: Decimal = Field(..., alias="total")
    paid: int = Field(..., alias="paid")
    created_at: int = Field(..., alias="createdAt")
    order_items: list[OrderItemResponse] = Field(..., alias="orderItems")