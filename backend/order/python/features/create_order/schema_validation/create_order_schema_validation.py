from pydantic import BaseModel, Field, ConfigDict
from typing import List

class OrderItemsCreateOrderValidation(BaseModel):
    model_config = ConfigDict(
        extra="allow",
        title="Order Items Create Order Validation",
        description="Represents order items create order validation",
        arbitrary_types_allowed=True
    )
    product_id: int = Field(..., gt=0, alias="productId")
    quantity: int = Field(..., gt=0, alias="quantity")

    # class Config:
    #     # allow_population_by_field_name = True
    #     populate_by_name = True

class CreateOrderValidation(BaseModel):
    model_config = ConfigDict(
        extra="allow",
        title="Create Order Validation",
        description="Represents create order validation",
        arbitrary_types_allowed=True
    )

    user_id: int = Field(..., alias="userId", gt=0)
    order_items: List[OrderItemsCreateOrderValidation] = Field(..., alias="orderItems")

    # class Config:
    #     # allow_population_by_field_name = True
    #     populate_by_name = True