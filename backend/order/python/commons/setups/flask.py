from flask import Flask, request
import datetime

app = Flask(__name__)

# @app.before_request
# def log_request():
#     url_split = request.url.split(":")
#     protocol = url_split[0]
#     print({"requestTime": datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S.%f"), "app": "project-backend-order", "method": request.method, "requestId": request.headers['X-REQUEST-ID'], "host": request.host, "urlPath": request.full_path, "protocol": protocol, "body": request.get_data(as_text=True), "userAgent": request.headers["user-agent"], "remoteAddr": request.remote_addr, "forwardedFor": request.headers["x-forwarded-for"]})

# @app.after_request
# def log_response(response):
#     print({"responseTime": datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S.%f"), "app": "project-backend-order", "requestId": request.headers["X-REQUEST-ID"], "responseStatus": response.status_code, "response": response.get_data(as_text=True)})