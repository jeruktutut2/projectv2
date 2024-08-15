from pydantic import BaseModel, ValidationError
from typing import Union, Type
from ..exception.response_exception import ResponseException
import json

def validate(schema: Type[BaseModel], data: dict) -> Union[BaseModel, ResponseException]:
    try:
        return schema.model_validate_json(json.dumps(data))
    except Exception as e:
        errorMessages = []
        for error in e.errors():
            field = str(error.get("loc"))
            field = field.replace("(", "").replace(")", "").replace("'", "").replace(",", ".").replace(" ", "")
            field = field[:-1] if field[-1] == "." else field
            errorMessage = {"field": field, "message": error.get("msg")}
            errorMessages.append(errorMessage)
        raise ResponseException(400, errorMessages, "validation error")
