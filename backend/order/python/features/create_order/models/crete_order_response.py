from decimal import Decimal
from pydantic import BaseModel, Field, ConfigDict

class OrderItemsCreateOrderResponse(BaseModel):
    model_config = ConfigDict(
        extra="allow",
        title="Order Items Create Order Response",
        description="Represents order items create order response",
        arbitrary_types_allowed=True,
        populate_by_name=True
    )
    
    id: int = Field(..., alias="id")
    order_id: int = Field(..., alias="orderId")
    product_id: int = Field(..., alias="productId")
    price: Decimal = Field(..., alias="price")
    quantity: int = Field(..., alias="quantity")
    total: Decimal = Field(..., alias="total")

class CreateOrderResponse(BaseModel):
    model_config = ConfigDict(
        extra="allow",
        title="Create Order Response",
        description="Represents create order response",
        arbitrary_types_allowed=True,
        populate_by_name=True
    )
    id: int = Field(..., alias="id")
    total: Decimal = Field(..., alias="total")
    order_items: list[OrderItemsCreateOrderResponse] = Field(..., alias="orderItems")