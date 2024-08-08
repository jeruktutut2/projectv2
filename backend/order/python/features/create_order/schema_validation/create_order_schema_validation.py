from pydantic import BaseModel, Field
from typing import List

class OrderItemsCreateOrderValidation(BaseModel):
    product_id: int = Field(..., gt=0, alias="productId")
    quantity: int = Field(..., gt=0, alias="quantity")

    class Config:
        # allow_population_by_field_name = True
        populate_by_name = True

class CreateOrderValidation(BaseModel):
    user_id: int = Field(..., gt=0, alias="userId")
    order_items: List[OrderItemsCreateOrderValidation] = Field(..., alias="orderItems")

    class Config:
        # allow_population_by_field_name = True
        populate_by_name = True