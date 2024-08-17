from flask import Flask, request
import datetime
from features.create_order.routes.create_order_route import create_order_route
from commons.middlewares.log_middleware import set_log_middleware
from features.get_order.routes.get_order_route import get_order_route
from features.update_order.routes.update_order_route import update_order_route
from features.delete_order.routes.delete_order_route import delete_order_route

app = Flask(__name__)
create_order_route(app)
get_order_route(app)
update_order_route(app)
delete_order_route(app)
set_log_middleware(app)