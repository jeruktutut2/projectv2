from pydantic import BaseModel, Field, ConfigDict

class DeleteOrderResponse(BaseModel):
    model_config = ConfigDict(
        extra="allow",
        title="Delete Order Response",
        description="Represents delete order response",
        arbitrary_types_allowed=True,
        populate_by_name=True
    )
    message: str = Field(..., alias="message")