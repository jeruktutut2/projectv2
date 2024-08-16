import pytest
from commons.utils.mysql_util import MysqlUtil
from commons.setups.flask import app
from tests.initialize.orders import create_table_orders, create_data_orders, get_data_orders, delete_table_orders
from tests.initialize.order_items import create_table_order_items, create_data_order_items, get_data_order_items, delete_table_order_items

connection = None
cursor = None
# client = None
request = {"userId": 1, "orderItems": [{"productId": 1, "quantity": 1}, {"productId": 2, "quantity": 2}]}
headers = {"X-REQUEST-ID": "requestId"}

@pytest.fixture(scope='module', autouse=True)
def setup_module():
    MysqlUtil.get_instance()
    global connection
    connection = MysqlUtil.get_connection()
    global cursor
    cursor = connection.cursor()
    # global client
    # app.run(debug=True)
    # client = app.test_client()
    yield
    cursor.close()
    connection.close()

@pytest.fixture(scope='function', autouse=True)
def setup_function(request):
    print()
    print(request.function.__name__)
    print()
    yield

@pytest.fixture
def client():
    # app.config["TESTING"] = True
    with app.test_client() as client:
        yield client

def test_request_id_doesnt_exists(client):
    response = client.post("/api/v1/order", json=request)
    # print("response:", response, response.status, response.status_code, response.json, response.json.get("data"))
    assert response.status_code == 400
    assert response.json.get("data") == None
    assert response.json.get("errors") == [{'field': 'message', 'message': 'cannot find request_id'}]

def test_validation(client):
    # headers = {"X-REQUEST-ID": "requestId"}
    request = {"userId": "a", "orderItems": [{"productId": "b", "quantity": "c"}, {"productId": "d", "quantity": "e"}]}
    # request = {"name": "Alice", "age": 30}
    # print("client1:", client)
    response = client.post("/api/v1/order", json=request, headers=headers)
    # print("client2:", client)
    # print("response:", response, response.status_code, response.json)
    assert response.status_code == 400
    assert response.json.get("data") == None
    assert response.json.get("errors") == [{'field': 'userId', 'message': 'Input should be a valid integer, unable to parse string as an integer'}, {'field': 'orderItems.0.productId', 'message': 'Input should be a valid integer, unable to parse string as an integer'}, {'field': 'orderItems.0.quantity', 'message': 'Input should be a valid integer, unable to parse string as an integer'}, {'field': 'orderItems.1.productId', 'message': 'Input should be a valid integer, unable to parse string as an integer'}, {'field': 'orderItems.1.quantity', 'message': 'Input should be a valid integer, unable to parse string as an integer'}]

def test_internal_server_error_no_table(client):
    connection.start_transaction()
    delete_table_order_items(cursor)
    delete_table_orders(cursor)
    connection.commit()
    response = client.post("/api/v1/order", json=request, headers=headers)
    # print("response:", response, response.json)
    assert response.status_code == 500
    assert response.json.get("data") == None
    assert response.json.get("errors") == [{'field': 'message', 'message': 'internal server error'}]

def test_success(client):
    connection.start_transaction()
    delete_table_order_items(cursor)
    delete_table_orders(cursor)
    create_table_orders(cursor)
    create_table_order_items(cursor)
    connection.commit()
    response = client.post("/api/v1/order", json=request, headers=headers)
    # print("response:", response, response.json)
    assert response.status_code == 201
    assert response.json.get("data") == '{"id":1,"total":"3.00","orderItems":[{"id":0,"orderId":0,"productId":1,"price":"1.00","quantity":1,"total":"1.00"},{"id":0,"orderId":0,"productId":2,"price":"1.00","quantity":2,"total":"2.00"}]}'
    assert response.json.get("errors") == None
