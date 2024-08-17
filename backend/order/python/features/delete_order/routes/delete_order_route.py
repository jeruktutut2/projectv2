from features.delete_order.controllers.delete_order_controller import DeleteOrderController
def delete_order_route(app):
    app.add_url_rule("/api/v1/order", view_func=DeleteOrderController.delete_order, methods=['DELETE'])