# from commons.setups.flask import app
from flask import request, jsonify
import datetime
import traceback

def set_log_middleware(app):

    @app.before_request
    def log_request():
        url_split = request.url.split(":")
        protocol = url_split[0]
        request_id = request.headers.get('X-REQUEST-ID', "")
        print({"requestTime": datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S.%f"), "app": "project-backend-order", "method": request.method, "requestId": request_id, "host": request.host, "urlPath": request.full_path, "protocol": protocol, "body": request.get_data(as_text=True), "userAgent": request.headers["user-agent"], "remoteAddr": request.remote_addr, "forwardedFor": request.headers.get("x-forwarded-for", "")})
        if request_id == "" or request_id == None:
            print({"logTime": datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S.%f"), "app": "project-backend-order", "requestId": request_id, "error": traceback.format_exc(), "message": "cannot find request_id"})
            return jsonify({"data": None, "errors": [{"field": "message", "message": "cannot find request_id"}]}), 400

    @app.after_request
    def log_response(response):
        request_id = request.headers.get('X-REQUEST-ID', "")
        print({"responseTime": datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S.%f"), "app": "project-backend-order", "requestId": request_id, "responseStatus": response.status_code, "response": response.get_data(as_text=True)})
        return response