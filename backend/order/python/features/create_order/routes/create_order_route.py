from features.create_order.controllers.create_order_controller import CreateOrderController

def create_order_route(app):
    app.add_url_rule("/api/v1/order", view_func=CreateOrderController.create_order, methods=['POST'])