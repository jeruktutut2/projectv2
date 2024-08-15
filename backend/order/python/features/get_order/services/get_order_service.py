from commons.exception.error_exception import error_handler
from commons.utils.mysql_util import MysqlUtil
from features.get_order.repositories.order_repository import OrderRepository
from features.get_order.repositories.order_item_repository import OrderItemsRepository
from commons.exception.response_exception import ResponseException
from commons.exception.error_exception import set_error_messages
from features.get_order.models.get_order_response import GetOrderResponse, OrderItemResponse

class GetOrderService:

    @staticmethod
    def get_order(requestId, id) -> GetOrderResponse:
        connection = None
        cursor = None
        try:
            connection = MysqlUtil.get_connection()
            cursor = connection.cursor()
            # why always use transaction, because by default connection autocommit false
            connection.start_transaction()
            
            order = OrderRepository.find_by_id(cursor, id)
            if order == None :
                raise ResponseException(404, set_error_messages("cannot find order by id: " + str(id)), "cannot find order by id: " + str(id))
            
            order_items = OrderItemsRepository.find_by_order_id(cursor, id)
            order_itmes_response = []
            for order_item in order_items:
                order_itmes_response.append(OrderItemResponse(id=order_item.id, orderId=order_item.order_id, productId=order_item.product_id, price=order_item.price, quantity=order_item.quantity, total=order_item.total))
            get_order_response = GetOrderResponse(id=order.id, userId=order.user_id, total=order.total, paid=order.paid, createdAt=order.created_at, orderItems=order_itmes_response)

            connection.commit()
            return get_order_response
        except Exception as e:
            if connection:
                connection.rollback()
            
            error_handler(e, requestId)
        finally:
            if cursor:
                cursor.close()

            if connection:
                connection.close()