import pytest
from commons.utils.mysql_util import MysqlUtil
from commons.setups.flask import app
from tests.initialize.orders import create_table_orders, create_data_orders, get_data_orders, delete_table_orders
from tests.initialize.order_items import create_table_order_items, create_data_order_items, get_data_order_items, delete_table_order_items

connection = None
cursor = None
headers = {"X-REQUEST-ID": "requestId"}

@pytest.fixture(scope='module', autouse=True)
def setup_module():
    MysqlUtil.get_instance()
    global connection
    connection = MysqlUtil.get_connection()
    global cursor
    cursor = connection.cursor()
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
    app.config["TESTING"] = True
    with app.test_client() as client:
        yield client

def test_request_id_doesnt_exists(client):
    response = client.get("/api/v1/order?orderId=1")
    assert response.status_code == 400
    assert response.json.get("data") == None
    assert response.json.get("errors") == [{'field': 'message', 'message': 'cannot find request_id'}]

def test_internal_server_error_no_table(client):
    connection.start_transaction()
    delete_table_order_items(cursor)
    delete_table_orders(cursor)
    connection.commit()
    response = client.get("/api/v1/order?orderId=1", headers=headers)
    assert response.status_code == 500
    assert response.json.get("data") == None
    assert response.json.get("errors") == [{'field': 'message', 'message': 'internal server error'}]

def test_cannot_found_data_order(client):
    connection.start_transaction()
    delete_table_order_items(cursor)
    delete_table_orders(cursor)
    create_table_orders(cursor)
    create_table_order_items(cursor)
    connection.commit()
    response = client.get("/api/v1/order?orderId=1", headers=headers)
    assert response.status_code == 404
    assert response.json.get("data") == None
    assert response.json.get("errors") == [{'field': 'message', 'message': 'cannot find order by id: 1'}]

def test_success(client):
    connection.start_transaction()
    delete_table_order_items(cursor)
    delete_table_orders(cursor)
    create_table_orders(cursor)
    create_table_order_items(cursor)
    create_data_orders(cursor)
    create_data_order_items(cursor)
    connection.commit()
    response = client.get("/api/v1/order?orderId=1", headers=headers)
    assert response.status_code == 200
    assert response.json.get("data") == '{"id":1,"userId":1,"total":"10.00","paid":0,"createdAt":1722390867657,"orderItems":[{"id":1,"orderId":1,"productId":1,"price":"1.00","quantity":1,"total":"1.00"},{"id":2,"orderId":1,"productId":2,"price":"2.00","quantity":2,"total":"2.00"},{"id":3,"orderId":1,"productId":3,"price":"3.00","quantity":3,"total":"3.00"}]}'
    assert response.json.get("errors") == None