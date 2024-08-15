import pytest
from commons.utils.mysql_util import MysqlUtil
from tests.initialize.orders import create_table_orders, create_data_orders, get_data_orders, delete_table_orders
from tests.initialize.order_items import create_table_order_items, create_data_order_items, get_data_order_items, delete_table_order_items
from commons.exception.response_exception import ResponseException
from features.get_order.services.get_order_service import GetOrderService
from decimal import Decimal

connection = None
cursor = None

requestId = "requestId"
id = 1

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

def test_internal_server_error_no_table():
    connection.start_transaction()
    delete_table_order_items(cursor)
    delete_table_orders(cursor)
    connection.commit()

    with pytest.raises(ResponseException) as e:
        GetOrderService.get_order(requestId, id)
    assert e.value.status == 500
    assert e.value.message == "internal server error"

def test_cannot_found_data_order():
    connection.start_transaction()
    delete_table_order_items(cursor)
    delete_table_orders(cursor)
    create_table_orders(cursor)
    create_table_order_items(cursor)
    connection.commit()

    with pytest.raises(ResponseException) as e:
        GetOrderService.get_order(requestId, id)
        
    assert e.value.status == 404
    assert e.value.message == "cannot find order by id: " + str(id)

def test_success():
    connection.start_transaction()
    delete_table_order_items(cursor)
    delete_table_orders(cursor)
    create_table_orders(cursor)
    create_table_order_items(cursor)
    create_data_orders(cursor)
    create_data_order_items(cursor)
    connection.commit()
    result = GetOrderService.get_order(requestId, id)
    assert result.id == 1
    assert result.user_id == 1
    assert result.total == Decimal(10.00)
    assert result.paid == 0