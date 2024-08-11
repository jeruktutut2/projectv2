import mysql
import os
import datetime

import mysql.connector

class MysqlUtil:
    _instance = None
    pool = None

    def __new__(cls, *args, **kwargs):
        cls._instance = super().__new__(cls, *args, **kwargs)
        try:
            print(datetime.datetime.now(), "mysql: new connecting to", os.environ.get("PROJECT_ORDER_MYSQL_HOST"))
            hostport = os.environ.get("PROJECT_ORDER_MYSQL_HOST", ":").split(":")
            cls.pool = mysql.connector.pooling.MySQLConnectionPool(
                pool_name="ordermysqlpool",
                pool_size=int(os.environ.get("PROJECT_ORDER_MYSQL_MAX_OPEN_CONNECTION", 0)),
                pool_reset_session=True,
                host=hostport[0],
                port=int(hostport[1]),
                database=os.environ.get("PROJECT_ORDER_MYSQL_DATABASE", ""),
                user=os.environ.get("PROJECT_ORDER_MYSQL_USERNAME", ""),
                password=os.environ.get("PROJECT_ORDER_MYSQL_PASSWORD", "")
            )
            print(datetime.datetime.now(), "mysql: new connected to", os.environ.get("PROJECT_ORDER_MYSQL_HOST"))
        except Exception as e:
            print("error when new connection:", e)
        return cls._instance
    
    @classmethod
    def get_instance(self):
        if not self._instance:
            self._instance = MysqlUtil()
        
        print(datetime.datetime.now(), "mysql: getting connection to", os.environ.get("PROJECT_ORDER_MYSQL_HOST"))
        connection = self.pool.get_connection()
        print(datetime.datetime.now(), "mysql: got connection to", os.environ.get("PROJECT_ORDER_MYSQL_HOST"))

        print(datetime.datetime.now(), "mysql: pinging connection to", os.environ.get("PROJECT_ORDER_MYSQL_HOST"))
        connection.ping(reconnect=True, attempts=3, delay=2)
        print(datetime.datetime.now(), "mysql: pinged connection to", os.environ.get("PROJECT_ORDER_MYSQL_HOST"))
        return self._instance
    
    @classmethod
    def get_connection(self):
        return self.pool.get_connection()