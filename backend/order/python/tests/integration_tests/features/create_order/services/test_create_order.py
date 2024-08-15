import pytest
from features.create_order.services.create_order_service import CreateOrderService
from commons.exception.response_exception import ResponseException
from tests.initialize.orders import create_table_orders, create_data_orders, delete_table_orders, get_data_orders
from commons.utils.mysql_util import MysqlUtil
from tests.initialize.order_items import create_table_order_items, create_data_order_items, delete_table_order_items, get_data_order_items
import time
from decimal import Decimal

requestId = "requestId"
request = {"userId": 1, "orderItems": [{"productId": 1, "quantity": 1}, {"productId": 2, "quantity": 2}]}
now_unix_milli = round(time.time() * 1000)

connection = None
cursor = None

@pytest.fixture(scope='module', autouse=True)
def setup_module():
    MysqlUtil.get_instance()
    global connection
    connection = MysqlUtil.get_connection()
    global cursor
    cursor = connection.cursor()
    # cursor = MysqlUtil.get_cursor(connection)
    yield
    cursor.close()
    # MysqlUtil.close_cursor(cursor)
    connection.close()
    # MysqlUtil.close_connection(connection)


@pytest.fixture(scope='function', autouse=True)
def setup_function(request):
    print()
    print(request.function.__name__)
    print()
    yield

def test_validation():
    request = {"userId": "a", "orderItems": [{"productId": "b", "quantity": "c"}, {"productId": "d", "quantity": "e"}]}
    with pytest.raises(ResponseException) as e:
        CreateOrderService.crate_order(requestId, request, now_unix_milli)
    assert e.value.status == 400
    assert e.value.message == "validation error"

def test_internal_server_error_no_table():
    # MysqlUtil.start_transaction(connection)
    connection.start_transaction()
    delete_table_order_items(cursor)
    delete_table_orders(cursor)
    # MysqlUtil.commit(connection)
    connection.commit()
    
    with pytest.raises(ResponseException) as e:
        CreateOrderService.crate_order(requestId, request, now_unix_milli)
    assert e.value.status == 500
    assert e.value.message == "internal server error"

def test_success():
    # MysqlUtil.start_transaction(connection)
    connection.start_transaction()
    delete_table_order_items(cursor)
    delete_table_orders(cursor)
    create_table_orders(cursor)
    create_table_order_items(cursor)
    # MysqlUtil.commit(connection)
    connection.commit()
    # got this error when trying to wrap mysql connection into method, ReferenceError: weakly-referenced object no longer exists
    result = CreateOrderService.crate_order(requestId, request, now_unix_milli)
    assert result.id == 1
    assert result.total == Decimal('3.00')
    assert len(result.order_items) == 2
    order_items = result.order_items
    # print("order_items:", order_items)
    # print("order_items[0].id:",order_items[0].id)
    assert order_items[0].product_id == 1
    assert order_items[0].price == Decimal('1.00')
    assert order_items[0].quantity == 1
    assert order_items[0].total == Decimal('1.00')
    assert order_items[1].product_id == 2
    assert order_items[1].price == Decimal('1.00')
    assert order_items[1].quantity == 2
    assert order_items[1].total == Decimal('2.00')

    orders = get_data_orders(cursor)
    assert len(orders) == 1
    assert orders[0][0] == 1
    assert orders[0][1] == 1
    assert orders[0][2] == Decimal('3.00')
    assert orders[0][3] == 0
    order_items = get_data_order_items(cursor)
    assert len(order_items) == 2
    assert order_items[0][0] == 1
    assert order_items[0][1] == 1
    assert order_items[0][2] == 1
    assert order_items[0][3] == Decimal('1.00')
    assert order_items[0][4] == 1
    assert order_items[0][5] == Decimal('1.00')
    assert order_items[1][0] == 2
    assert order_items[1][1] == 1
    assert order_items[1][2] == 2
    assert order_items[1][3] == Decimal('1.00')
    assert order_items[1][4] == 2
    assert order_items[1][5] == Decimal('2.00')