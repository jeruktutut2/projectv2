from commons.exception.error_exception import error_handler_response
from features.delete_order.services.delete_order_service import DeleteOrderService
from flask import request, jsonify

class DeleteOrderController:

    @staticmethod
    def delete_order():
        try:
            request_id = request.headers['X-REQUEST-ID']
            orderId = request.json.get("id")
            delete_order_response = DeleteOrderService.delete_order(request_id, orderId)
            return jsonify({"data": delete_order_response.model_dump(by_alias=True), "errors": None}), 200
        except Exception as e:
            return error_handler_response(e)