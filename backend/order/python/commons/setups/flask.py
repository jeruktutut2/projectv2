from flask import Flask, request
import datetime
from features.create_order.routes.create_order_route import create_order_route
from commons.middlewares.log_middleware import set_log_middleware

app = Flask(__name__)
create_order_route(app)
set_log_middleware(app)