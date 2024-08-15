from commons.utils.mysql_util import MysqlUtil
import pytest
from tests.initialize.orders import create_table_orders, create_data_orders, delete_table_orders, get_data_orders
from tests.initialize.order_items import create_table_order_items, create_data_order_items, delete_table_order_items, get_data_order_items
from commons.exception.response_exception import ResponseException
from features.delete_order.services.delete_order_service import DeleteOrderService

request_id = "request_id"
id = 1

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

def test_internal_server_error_no_table():
    connection.start_transaction()
    delete_table_order_items(cursor)
    delete_table_orders(cursor)
    connection.commit()
    with pytest.raises(ResponseException) as e:
        DeleteOrderService.delete_order(request_id, id)
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
    result = DeleteOrderService.delete_order(request_id, id)
    # print("result:", result)
    assert result.message == "succesfully deleted order"