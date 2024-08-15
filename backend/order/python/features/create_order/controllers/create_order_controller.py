from fastapi import APIRouter, Request
import json

router = APIRouter()

class CreateOrderController:

    @staticmethod
    @router.post("/api/v1/create-order")
    # async def create_order(request: Request):
    def create_order(request: Request):
        # await request.json()
        request_body = request.body()
        json_body = json.loads(request_body)
        print("json_body:", json_body)
