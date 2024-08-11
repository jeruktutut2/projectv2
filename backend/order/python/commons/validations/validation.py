from pydantic import BaseModel, ValidationError
from typing import Union, Type
from ..exception.response_exception import ResponseException
import json

def validate(schema: Type[BaseModel], data: dict) -> Union[BaseModel, ResponseException]:
    try:
        # schema.model_validate_json(**data)
        # result = schema.model_validate_json(json.dumps(data))
        return schema.model_validate_json(json.dumps(data))
        # print("result:", result)
    # except ValidationError as e:
    #     print("ValidationError e:", e)
        # return result
    except Exception as e:
        # print(1)
        # print("2 Exception e:", e)
        # print(3)
        # print("4 e.error():", e.errors())
        # print(5)
        # print(type(e.errors()))
        errorMessages = []
        for error in e.errors():
            # print(type(error))
            # print("error:", error.get("loc"), error.get("msg"))
            # field = string.replace(error.get("loc"), "")
            # field = error.get("loc").replace(",", ".")
            field = str(error.get("loc"))
            field = field.replace("(", "").replace(")", "").replace("'", "").replace(",", ".").replace(" ", "")
            field = field[:-1] if field[-1] == "." else field
            errorMessage = {"field": field, "message": error.get("msg")}
            errorMessages.append(errorMessage)

        # print("errorMessages:", errorMessages)
        raise ResponseException(400, errorMessages, "validation error")
