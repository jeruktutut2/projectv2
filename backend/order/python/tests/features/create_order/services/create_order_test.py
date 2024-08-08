# import pytest
# from .....features.create_order.services.create_order_service import CreateOrderService
# from features.create_order.services.create_order_service import CreateOrderService

# requestId = "requestId"
# request = {"userId": 1, "orderItems": [{"product_id": 1, "quantity": 1}]}

# @pytest.fixture(scope='module', autouse=True)
# def setup_module():
#     # requestId = "requestId"
#     yield

# @pytest.fixture(scope='function', autouse=True)
# def setup_function():
#     yield

# def test_validation():
#     CreateOrderService.crate_order(requestId, request)