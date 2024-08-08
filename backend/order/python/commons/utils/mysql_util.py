# import mysql.connector
import mysql
import os
import datetime

import mysql.connector

class MysqlUtil:
    _instance = None
    pool = None

    def __new__(cls, *args, **kwargs):
        if not cls._instance:
            cls._instance = super().__new__(cls, *args, **kwargs)
            # print(os.environ.get['PROJECT_ORDER_MYSQL_HOST'])
            # print(os.environ.get('PROJECT_ORDER_MYSQL_HOST'))
            # print("pool:", os.environ.get("PROJECT_ORDER_MYSQL_MAX_OPEN_CONNECTION", ""))
            # print("host:", os.environ.get("PROJECT_ORDER_MYSQL_HOST"))
            # print("database:", os.environ.get("PROJECT_ORDER_MYSQL_DATABASE"))
            # print("username:", os.environ.get("PROJECT_ORDER_MYSQL_USERNAME"))
            # print("password:", os.environ.get("PROJECT_ORDER_MYSQL_PASSWORD"))
            try:
                print(datetime.datetime.now(), "mysql: new connecting to", os.environ.get("PROJECT_ORDER_MYSQL_HOST"))
                hostport = os.environ.get("PROJECT_ORDER_MYSQL_HOST", ":").split(":")
                # print(hostport[0], hostport[1])
                cls.pool = mysql.connector.pooling.MySQLConnectionPool(
                    pool_name="ordermysqlpool",
                    # pool_size=10, #os.environ.get("PROJECT_ORDER_MYSQL_MAX_OPEN_CONNECTION", 0),
                    pool_size=int(os.environ.get("PROJECT_ORDER_MYSQL_MAX_OPEN_CONNECTION", 0)),
                    pool_reset_session=True,
                    # host="localhost", #os.environ.get("PROJECT_ORDER_MYSQL_HOST"),
                    host=hostport[0],
                    # port=3309,
                    port=int(hostport[1]),
                    database=os.environ.get("PROJECT_ORDER_MYSQL_DATABASE", ""),
                    user=os.environ.get("PROJECT_ORDER_MYSQL_USERNAME", ""),
                    password=os.environ.get("PROJECT_ORDER_MYSQL_PASSWORD", "")
                    # option_files=f"mysql+mysqlconnector://{os.environ.get("PROJECT_ORDER_MYSQL_USERNAME")}:{os.environ.get("PROJECT_ORDER_MYSQL_PASSWORD")}@{os.environ.get("PROJECT_ORDER_MYSQL_HOST")}/{os.environ.get("PROJECT_ORDER_MYSQL_DATABASE")}"
                )
                print(datetime.datetime.now(), "mysql: new connected to", os.environ.get("PROJECT_ORDER_MYSQL_HOST"))
                print(datetime.datetime.now(), "mysql: getting connection to", os.environ.get("PROJECT_ORDER_MYSQL_HOST"))
                connection = cls.pool.get_connection()
                print(datetime.datetime.now(), "mysql: got connection to", os.environ.get("PROJECT_ORDER_MYSQL_HOST"))
                print(datetime.datetime.now(), "mysql: pinging connection to", os.environ.get("PROJECT_ORDER_MYSQL_HOST"))
                connection.ping(reconnect=True, attempts=3, delay=2)
                print(datetime.datetime.now(), "mysql: pinged connection to", os.environ.get("PROJECT_ORDER_MYSQL_HOST"))

            # except mysql.connector.Error as error:
            except Exception as e:
                # print("error when new connection:", error)
                print("error when new connection:", e)
            # print(datetime.datetime.now(), "mysql: pooling connection to", os.environ.get("PROJECT_ORDER_MYSQL_HOST"))
            # cls.pool = mysql.connector.pooling.MySQLConnectionPool(
            #     pool_name="ordermysqlpool",
            #     pool_size=os.environ.get("PROJECT_ORDER_MYSQL_MAX_OPEN_CONNECTION"),
            #     pool_reset_session=True,
            #     host=os.environ.get("PROJECT_ORDER_MYSQL_HOST"),
            #     database=os.environ.get("PROJECT_ORDER_MYSQL_DATABASE"),
            #     user=os.environ.get("PROJECT_ORDER_MYSQL_USERNAME"),
            #     password=os.environ.get("PROJECT_ORDER_MYSQL_PASSWORD")
            # )
            # print(datetime.datetime.now(), "mysql: pooled connection to", os.environ.get("PROJECT_ORDER_MYSQL_HOST"))
            # return cls._instance
        return cls._instance
    
    def get_connection(self):
        # try:
            # connection = self.pool.get_connection()
            # return connection
            return self.pool.get_connection()
            
            # connection.ping(reconnect=True, attempts=3, delay=2)
        # except mysql.connector.Error as error:
        #     print("error when getting and pinging connection", error)
        # finally:
        #     connection.close()
        

mysqlUtil = MysqlUtil()
# mysqlUtil.get_connection()