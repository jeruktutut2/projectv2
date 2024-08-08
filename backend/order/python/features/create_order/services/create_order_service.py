# from ....commons.validations.validation import validate
# from ..schema_validation.create_order_schema_validation import CreateOrderValidation

from commons.validations.validation import validate
from features.create_order.schema_validation.create_order_schema_validation import CreateOrderValidation

class CreateOrderService:

    @staticmethod
    def crate_order(requestId, request):
        try:
            print("mantap")
            validate(CreateOrderValidation, request)
        except Exception as e:
            print(e)
        finally:
            print()
