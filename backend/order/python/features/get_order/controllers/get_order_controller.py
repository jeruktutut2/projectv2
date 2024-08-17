from commons.exception.error_exception import error_handler_response
from features.get_order.services.get_order_service import GetOrderService
from flask import request, jsonify

class GetOrderController:

    @staticmethod
    def get_order():
        try:
            request_id = request.headers['X-REQUEST-ID']
            orderId = request.args.get("orderId")
            get_order_response = GetOrderService.get_order(request_id, orderId)
            return jsonify({"data": get_order_response.model_dump(by_alias=True), "errors": None}), 200
        except Exception as e:
            return error_handler_response(e)