import pytest
# from .....features.create_order.services.create_order_service import CreateOrderService
from features.create_order.services.create_order_service import CreateOrderService

requestId = "requestId"
request = {"userId": 1, "orderItems": [{"product_id": 1, "quantity": 1}]}

@pytest.fixture(scope='module', autouse=True)
def setup_module():
    # requestId = "requestId"
    # print("setup_module1")
    yield
    # print("setup_module2")

@pytest.fixture(scope='function', autouse=True)
def setup_function():
    # print("setup_function1")
    yield
    # print("setup_function2")

# def test_1():
#     print("test_1")

# @pytest.fixture(scope="module", autouse=True)
# def setup_module():
#     print("setup_module: Ini dijalankan sekali sebelum semua tes dalam modul")
#     yield
#     print("teardown_module: Ini dijalankan sekali setelah semua tes dalam modul")

# @pytest.fixture(scope="function", autouse=True)
# def setup_function():
#     print("setup_function: Ini dijalankan sebelum setiap tes fungsi")
#     yield
#     print("teardown_function: Ini dijalankan setelah setiap tes fungsi")

# def test_upper():
#     print("test_upper")
#     assert 'foo'.upper() == 'FOO'

def test_validation():
    print("mantap test")
    CreateOrderService.crate_order(requestId, request)

# def test_isupper():
#     print("test_isupper")
#     assert 'FOO'.isupper()
#     assert not 'Foo'.isupper()

# def test_split():
#     print("test_split")
#     s = 'hello world'
#     assert s.split() == ['hello', 'world']
#     with pytest.raises(TypeError):
#         s.split(2)