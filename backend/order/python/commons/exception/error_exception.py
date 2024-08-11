from commons.exception.response_exception import ResponseException
import datetime
import traceback

def error_handler(error, requestId):
    print({"logTime": datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S.%f"), "app": "backend-order-python", "requestId": requestId, "error": traceback.format_exc()})
    if isinstance(error, ResponseException):
        raise ResponseException(error.status, error.error_messages, error.message)
    else:
        raise ResponseException(500, set_internal_server_error(), "internal server error")

def set_error_messages(message):
    return [{"field": "message", "message": message}]

def set_internal_server_error():
    return [{"field": "message", "message": "internal server error"}]