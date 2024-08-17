from flask import request, jsonify
from features.create_order.services.create_order_service import CreateOrderService
from commons.exception.error_exception import error_handler_response
import time

class CreateOrderController:

    @staticmethod
    def create_order():
        try:
            request_id = request.headers['X-REQUEST-ID']
            request_body = request.json
            now_unix_milli = round(time.time() * 1000)
            createOrder = CreateOrderService.crate_order(request_id, request_body, now_unix_milli)
            return jsonify({"data": createOrder.model_dump(by_alias=True), "errors": None}), 201
        except Exception as e:
            return error_handler_response(e)