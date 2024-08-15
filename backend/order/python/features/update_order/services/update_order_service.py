from commons.exception.error_exception import error_handler
from features.update_order.schema_validation.update_order_schema_validation import UpdateOrderValidation
from commons.validations.validation import validate
from features.update_order.repositories.order_repository import OrderRepository
from commons.utils.mysql_util import MysqlUtil
from commons.exception.response_exception import ResponseException
from commons.exception.error_exception import set_error_messages
from features.update_order.repositories.order_item_repository import OrderItemRepository
from features.update_order.repositories.product_repository import ProductRepository
from features.update_order.models.order_item import OrderItem
from decimal import Decimal
from features.update_order.models.update_order_response import UpdateOrderResponse, OrderItemsUpdateOrderResponse

class UpdateOrderService:

    @staticmethod
    def update_order(request_id, request_body) -> UpdateOrderResponse:
        connection = None
        cursor = None
        try:
            request = validate(UpdateOrderValidation, request_body)

            connection = MysqlUtil.get_connection()
            cursor = connection.cursor()
            connection.start_transaction()

            order = OrderRepository.find_by_id(cursor, request.id)
            if order == None:
                raise ResponseException(404, set_error_messages("cannot find order id: " + str(request_body.id), "cannot find order id: " + str(request_body.id)))
            
            product_ids = []
            for order_item in request.order_items:
                product_ids.append(order_item.product_id)
            
            products = ProductRepository.find_in_id(cursor, product_ids)
            if len(products) != len(product_ids):
                raise ResponseException(400, set_error_messages("number of products are not same "+str(len(products))+":"+str(len(product_ids))), "number of products are not same "+str(len(products))+":"+str(len(product_ids)))
            
            OrderItemRepository.delete_by_order_id(cursor, order.id)
            if cursor.rowcount != len(product_ids):
                raise ResponseException(400, set_error_messages("number of order items ar not same "+str(cursor.rowcount)+":"+str(len(product_ids))), "number of order items ar not same "+str(cursor.rowcount)+":"+str(len(product_ids)))
            
            inserted_order_items = []
            total_order = Decimal(0.0)
            for order_item in request.order_items:
                for product in products:
                    if product.id == order_item.product_id:
                        total_order_item = Decimal(order_item.quantity) * Decimal(product.price)
                        inserted_order_items.append(OrderItem(0, order.id, order_item.product_id, product.price, order_item.quantity, total_order_item))
                        total_order += Decimal(total_order_item)
            OrderItemRepository.create(cursor, inserted_order_items)
            if cursor.rowcount != len(inserted_order_items):
                raise Exception("number of created order items are not same " + str(cursor.rowcount) + ":" + str(len(inserted_order_items)))
            
            OrderRepository.update_total(cursor, total_order, order.id)
            if cursor.rowcount != 1:
                raise Exception("number of updated total order is not one: " + str(cursor.rowcount))
            
            order_items = OrderItemRepository.find_by_order_id(cursor, order.id)
            if len(order_items) != len(inserted_order_items):
                raise Exception("number of order items and inserted order items is not same " + str(len(order_items)) + ":" + str(len(inserted_order_items)))

            connection.commit()

            order_item_update_order_response = []
            for order_item in order_items:
                order_item_update_order_response.append(OrderItemsUpdateOrderResponse(id=order_item.id, productId=order_item.product_id, price=order_item.price, quantity=order_item.quantity, total=order_item.total))

            update_order_response = UpdateOrderResponse(id=order.id, total=total_order, paid=0, orderItems=order_item_update_order_response)

            return update_order_response
        except Exception as e:
            if connection:
                connection.rollback()
            
            error_handler(e, request_id)
        finally:
            if cursor:
                cursor.close()
            
            if connection:
                connection.close()