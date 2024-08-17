from features.get_order.controllers.get_order_controller import GetOrderController

def get_order_route(app):
    app.add_url_rule("/api/v1/order", view_func=GetOrderController.get_order, methods=['GET'])