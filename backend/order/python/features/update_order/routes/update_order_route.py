from features.update_order.controllers.update_order_controller import UpdateOrderController

def update_order_route(app):
    app.add_url_rule("/api/v1/order", view_func=UpdateOrderController.update_order, methods=['PATCH'])