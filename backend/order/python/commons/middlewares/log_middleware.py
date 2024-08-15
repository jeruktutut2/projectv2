from commons.setups.flask import app
from flask import request, jsonify
import datetime
import traceback

@app.before_request
def log_request():
    url_split = request.url.split(":")
    protocol = url_split[0]
    print({"requestTime": datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S.%f"), "app": "project-backend-order", "method": request.method, "requestId": request.headers['X-REQUEST-ID'], "host": request.host, "urlPath": request.full_path, "protocol": protocol, "body": request.get_data(as_text=True), "userAgent": request.headers["user-agent"], "remoteAddr": request.remote_addr, "forwardedFor": request.headers["x-forwarded-for"]})

    if 'X-REQUEST-ID' not in request.headers:
        print({"logTime": datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S.%f"), "app": "project-backend-order", "requestId": request.headers['X-REQUEST-ID'], "error": traceback.format_exc(), "message": "cannot find request_id"})
        return jsonify({"data": None, "error": [{"field": "message", "message": "cannot find request_id"}]}), 400

@app.after_request
def log_response(response):
    print({"responseTime": datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S.%f"), "app": "project-backend-order", "requestId": request.headers["X-REQUEST-ID"], "responseStatus": response.status_code, "response": response.get_data(as_text=True)})