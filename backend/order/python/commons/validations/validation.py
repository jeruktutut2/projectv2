from pydantic import BaseModel, ValidationError
from typing import Union, Type
from ..exception.response_exception import ResponseException

def validate(schema: Type[BaseModel], data: dict) -> Union[BaseModel, ResponseException]:
    try:
        schema.model_validate(data)
    except ValidationError as e:
        print("validation e:", e)