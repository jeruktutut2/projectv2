from commons.utils.mysql_util import MysqlUtil
import pytest
from commons.exception.response_exception import ResponseException
from features.update_order.services.update_order_service import UpdateOrderService
from tests.initialize.orders import create_table_orders, create_data_orders, delete_table_orders, get_data_orders
from tests.initialize.order_items import create_table_order_items, create_data_order_items, delete_table_order_items, get_data_order_items
from decimal import Decimal

request_id = "requestId"
request = {"id": 1, "orderItems": [{"productId": 1, "quantity": 2}, {"productId": 2, "quantity": 3}, {"productId": 3, "quantity": 1}]}

connection = None
cursor = None

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

def test_validation():
    request = {"id": 0, "orderItems": [{"productId": 0, "quantity": 0}]}
    with pytest.raises(ResponseException) as e:
        UpdateOrderService.update_order(request_id, request)
    e.value.status == 400
    e.value.message == "validation status"


def test_internal_server_error_no_table():
    connection.start_transaction()
    delete_table_order_items(cursor)
    delete_table_orders(cursor)
    # delete_table_order_items(cursor)
    connection.commit()
    with pytest.raises(ResponseException) as e:
        UpdateOrderService.update_order(request_id, request)
    # print("e:", e)
    e.value.status == 500
    e.value.message == "internal server error"

def test_success():
    connection.start_transaction()
    delete_table_order_items(cursor)
    delete_table_orders(cursor)
    create_table_orders(cursor)
    create_table_order_items(cursor)
    create_data_orders(cursor)
    create_data_order_items(cursor)
    connection.commit()
    result = UpdateOrderService.update_order(request_id, request)
    # print("result:", result)
    assert result.id == 1
    assert result.total == Decimal(6.0)
    assert len(result.order_items) == 3
    assert result.order_items[0].product_id == 1
    assert result.order_items[0].price == Decimal(1.0)
    assert result.order_items[0].quantity == 2
    assert result.order_items[0].total == 2
    assert result.order_items[1].product_id == 2
    assert result.order_items[1].price == 1
    assert result.order_items[1].quantity == 3
    assert result.order_items[1].total == 3
    assert result.order_items[2].product_id == 3
    assert result.order_items[2].price == 1
    assert result.order_items[2].quantity == 1
    assert result.order_items[2].total == 1
