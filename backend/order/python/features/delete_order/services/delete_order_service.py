from commons.utils.mysql_util import MysqlUtil
from features.delete_order.repositories.order_repository import OrderRepository
from features.delete_order.repositories.order_item_repository import OrderItemRepository
from commons.exception.error_exception import error_handler
from features.delete_order.models.delete_order_response import DeleteOrderResponse

class DeleteOrderService:

    @staticmethod
    def delete_order(request_id, id) -> DeleteOrderResponse:
        connection = None
        cursor = None
        try:
            connection = MysqlUtil.get_connection()
            cursor = connection.cursor()
            connection.start_transaction()

            OrderItemRepository.delete_by_order_id(cursor, id)
            if cursor.rowcount < 1:
                raise Exception("number of deleted are less than one: " + str(cursor.rowcount))
            
            OrderRepository.delete_by_id(cursor, id)
            if cursor.rowcount != 1:
                raise Exception("number of deleted are not one:" + str(cursor.rowcount))
            
            connection.commit()

            delete_order_response = DeleteOrderResponse(message="succesfully deleted order")
            return delete_order_response
        except Exception as e:
            if connection:
                connection.rollback()
            
            error_handler(e, request_id)
        finally:
            if cursor:
                cursor.close()
            
            if connection:
                connection.close()