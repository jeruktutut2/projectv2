from decimal import Decimal
from pydantic import BaseModel, Field, ConfigDict

class OrderItemsUpdateOrderResponse(BaseModel):
    model_config = ConfigDict(
        extra="allow",
        title="Order Items Update Order Response",
        description="Represents order items update order response",
        arbitrary_types_allowed=True,
        populate_by_name=True
    )
    id: int = Field(..., alias="id")
    product_id: int = Field(..., alias="productId")
    price: Decimal = Field(..., alias="price")
    quantity: int = Field(..., alias="quantity")
    total: Decimal = Field(..., alias="total")
    
class UpdateOrderResponse(BaseModel):
    model_config = ConfigDict(
        extra="allow",
        title="Update Order Response",
        description="Represents update order response",
        arbitrary_types_allowed=True,
        populate_by_name=True
    )
    id: int = Field(..., alias="id")
    total: Decimal = Field(..., alias="total")
    paid: int = Field(..., alias="paid")
    order_items: list[OrderItemsUpdateOrderResponse] = Field(..., alias="orderItems")