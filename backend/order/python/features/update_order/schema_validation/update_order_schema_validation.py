from pydantic import BaseModel, Field, ConfigDict
from typing import List

class OrderItemsUpdateOrderValidaion(BaseModel):
    model_config = ConfigDict(
        extra="allow",
        title="Order Items Update Order Validation",
        description="Represents order items update order validation",
        arbitrary_types_allowed=True
    )
    product_id: int = Field(..., gt=0, alias="productId")
    quantity: int = Field(..., gt=0, alias="quantity")


class UpdateOrderValidation(BaseModel):
    model_config = ConfigDict(
        extra="allow",
        title="Update Order Validation",
        description="Represents update order validation",
        arbitrary_types_allowed=True
    )
    id: int = Field(..., gt=0, alias="id")
    order_items: List[OrderItemsUpdateOrderValidaion] = Field(..., alias="orderItems")
