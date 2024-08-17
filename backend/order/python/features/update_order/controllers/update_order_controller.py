from commons.exception.error_exception import error_handler_response
from features.update_order.services.update_order_service import UpdateOrderService
from flask import request, jsonify

class UpdateOrderController:

    @staticmethod
    def update_order():
        try:
            request_id = request.headers['X-REQUEST-ID']
            update_order_response = UpdateOrderService.update_order(request_id, request.json)
            return jsonify({"data": update_order_response.model_dump(by_alias=True), "errors": None}), 200
        except Exception as e:
            return error_handler_response(e)