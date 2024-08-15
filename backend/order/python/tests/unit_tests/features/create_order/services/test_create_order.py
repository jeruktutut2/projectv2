import pytest
from unittest.mock import patch, MagicMock
from features.create_order.services.create_order_service import CreateOrderService
import time
from commons.exception.response_exception import ResponseException
# from commons.utils.mysql_util
from features.create_order.models.product import Product

request_id = "request_id"
request = {"userId": 1, "orderItems": [{"productId": 1, "quantity": 1}, {"productId": 2, "quantity": 2}]}
now_unix_milli = round(time.time() * 1000)

@pytest.fixture(scope="module", autouse=True)
def setup_module():
    yield

@pytest.fixture(scope="function", autouse=True)
def setup_function(request):
    print()
    print(request.function.__name__)
    print()
    yield

def test_validation():
    request = {"userId": "a", "orderItems": [{"productId": "b", "quantity": "c"}, {"productId": "d", "quantity": "e"}]}
    with pytest.raises(ResponseException) as e:
        CreateOrderService.crate_order(request_id, request, now_unix_milli)
    print("e.value:", e.value)
    assert e.value.status == 400
    assert e.value.message == "validation error"

# @patch("commons.utils.MysqlUtil.get_connection")
# @patch("commons.utils.mysql_util.get_connection")
@patch("features.create_order.services.create_order_service.MysqlUtil.get_connection")
def test_mysql_get_connection_error(mock_get_connection):
    # print()
    #  mock_get_connection.return_value = None
    mock_get_connection.side_effect = Exception("error getting connection")
    with pytest.raises(ResponseException) as e:
        CreateOrderService.crate_order(request_id, request, now_unix_milli)
    #  print("e.value:", e.value)
    assert e.value.status == 500
    assert e.value.message == "internal server error"

@patch("features.create_order.services.create_order_service.MysqlUtil.get_connection")
@patch("features.create_order.services.create_order_service.ProductRepository.find_in_id")
def test_product_repository_find_by_in_error(mock_get_connection, mock_product_repository_find_in_id):
    mock_conn = MagicMock()
    mock_get_connection.return_value = mock_conn

    mock_cursor = mock_conn.cursor.return_value

    mock_product_repository_find_in_id.side_effect = Exception("error product repository find in id")
    with pytest.raises(ResponseException) as e:
        CreateOrderService.crate_order(request_id, request, now_unix_milli)
    
    # mock_get_connection.assert_called_once()
    # mock_conn.cursor is not called because cursor didn't called for example execute, it called in ProductRepository.find_in_id method
    # mock_conn.cursor.assert_called_once()
    # mock_get_connection.cursor.assert_called_once()
    # mock_conn.start_transaction.assert_called_once()
    # mock_get_connection.start_transaction.assert_called_once()

    # mock_conn.rollback.assert_called_once()
    # mock_get_connection.rollback.assert_called_once()
    # mock_cursor.close.assert_called_once()
    # mock_conn.close.assert_called_once()
    # mock_get_connection.close.assert_called_once()
    e.value.status == 500
    e.value.message == "internal server error"

#  AttributeError: 'list' object has no attribute 'close', i don't know why the error like that
@patch("features.create_order.services.create_order_service.MysqlUtil.get_connection")
@patch("features.create_order.services.create_order_service.ProductRepository.find_in_id")
def test_product_repository_find_by_in_bad_request(mock_get_connection, mock_product_repository_find_in_id):
    mock_conn = MagicMock()
    mock_get_connection.return_value = mock_conn

    mock_cursor = mock_conn.cursor.return_value

    # mock_product_repository_find_in_id.return_value = [{"id": 1, "name": "name1", "description": "description1", "price": 1}, {"id": 2, "name": "name2", "description": "description2", "price": 2}]
    products = []
    products.append(Product(id=1, user_id=0, name="name1", description="description1",price=1,stock=1))
    mock_product_repository_find_in_id.return_value = products
    with pytest.raises(ResponseException) as e:
        CreateOrderService.crate_order(request_id, request, now_unix_milli)
    # e.value.status == 400
    # e.value.message == "number of products are not same 2:3"

# @patch("features.create_order.services.create_order_service.MysqlUtil.get_connection")
# @patch("features.create_order.services.create_order_service.ProductRepository.find_in_id")
# @patch("features.create_order.services.create_order_service.OrderRepository.create")
# def test_order_repository_create_error(mock_get_connection, mock_product_repository_find_in_id, mock_order_repository_create):
#     mock_conn = MagicMock()
#     mock_get_connection.return_value = mock_conn
#     mock_cursor = mock_conn.cursor.return_value

#     mock_product_repository_find_in_id.return_value = [{"id": 1, "name": "name1", "description": "description1", "price": 1}, {"id": 2, "name": "name2", "description": "description2", "price": 2}, {"id": 3, "name": "name3", "description": "description3", "price": 3}]
#     mock_order_repository_create.side_effect = Exception("error order repository create")
#     with pytest.raises(ResponseException) as e:
#         CreateOrderService.crate_order(request_id, request, now_unix_milli)

# @patch("features.create_order.services.create_order_service.MysqlUtil.get_connection")
# @patch("features.create_order.services.create_order_service.ProductRepository.find_in_id")
# @patch("features.create_order.services.create_order_service.OrderRepository.create")
