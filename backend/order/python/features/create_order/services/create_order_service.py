from commons.validations.validation import validate
from features.create_order.schema_validation.create_order_schema_validation import CreateOrderValidation
from features.create_order.repositories.order_repository import OrderRepository
from commons.utils.mysql_util import MysqlUtil
from features.create_order.models.order import Order
from features.create_order.repositories.product_repository import ProductRepository
from commons.exception.response_exception import ResponseException
from commons.exception.error_exception import set_error_messages, error_handler, set_internal_server_error
from decimal import Decimal
import datetime
from features.create_order.repositories.order_item_repository import OrderItemRepository
from features.create_order.models.order_item import OrderItem

class CreateOrderService:

    @staticmethod
    def crate_order(requestId, request_body, now_unix_milli):
        connection = None
        cursor = None
        try:
            request = validate(CreateOrderValidation, request_body)
            
            connection = MysqlUtil.get_connection()
            connection.start_transaction()
            cursor = connection.cursor()

            product_ids = []
            for order_item in request.order_items:
                product_ids.append(order_item.product_id)

            products = ProductRepository.find_in_id(cursor, product_ids)
            if len(products) != len(product_ids):
                raise ResponseException(400, set_error_messages("number orf products are not same " + len(products) +":"+ len(product_ids)), "number orf products are not same "+len(products)+":"+len(product_ids))
            
            total_order: Decimal = Decimal(0.0)
            order_items = []
            response_order_items = []
            for request_order_item in request.order_items:
                for product in products:
                    if request_order_item.product_id == product.id:
                        total_order += Decimal(request_order_item.quantity) * Decimal(product.price)

                        total_order_item: Decimal = 0.0
                        total_order_item = Decimal(request_order_item.quantity) * Decimal(product.price)
                        order_items.append(OrderItem(0, 0, request_order_item.product_id, product.price, request_order_item.quantity, total_order_item))

                        response_order_items.append({"productId": request_order_item.product_id, "price": product.price, "quantity": request_order_item.quantity, "total": total_order_item})

            order = Order(0, request.user_id, total_order, 0, now_unix_milli)
            OrderRepository.create(cursor, order)
            if cursor.rowcount != 1:
                print({"logTime": datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S.%f"), "app": "backend-order-python", "requestId": requestId, "error": "row count created order is not one: " + result.rowcount})
                raise ResponseException(500, set_internal_server_error(), "internal server error")
            order.id = cursor.lastrowid

            for order_item in order_items:
                order_item.order_id = order.id
            OrderItemRepository.create_many(cursor, order_items)

            connection.commit()

            return {"id": order.id, "total": order.total, "orderItems": response_order_items}
        except Exception as e:
            if connection:
                connection.rollback()
            
            error_handler(e, requestId)
        finally:
            if cursor:
                cursor.close()
            
            if connection:
                connection.close()
