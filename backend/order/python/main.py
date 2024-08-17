from commons.setups.flask import app
from commons.utils.mysql_util import MysqlUtil
import os

MysqlUtil.get_instance()

if __name__ == "__main__":
    hostport = os.environ.get("PROJECT_ORDER_APPLICATION_HOST").split(":")
    port = int(hostport[1])
    app.run(debug=True, port=port)
