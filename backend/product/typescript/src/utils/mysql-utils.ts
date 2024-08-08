import mysql, { Pool, PoolOptions } from "mysql2/promise";

export class MysqlUtil {

    private static instance: MysqlUtil
    private static pool: Pool

    private constructor() {
        console.log(new Date().toISOString(), "mysql: connecting to", process.env.PROJECT_PRODUCT_MYSQL_HOST);
        const connectionString = `mysql://`+process.env.PROJECT_PRODUCT_MYSQL_USERNAME+`:`+process.env.PROJECT_PRODUCT_MYSQL_PASSWORD+`@`+process.env.PROJECT_PRODUCT_MYSQL_HOST+`/`+process.env.PROJECT_PRODUCT_MYSQL_DATABASE+``
        const access: PoolOptions = {
            uri: connectionString,
            connectionLimit: parseInt(process.env.PROJECT_PRODUCT_MYSQL_MAX_OPEN_CONNECTION ?? "0"),
            maxIdle: parseInt(process.env.PROJECT_PRODUCT_MYSQL_MAX_IDLE_CONNECTION ?? "0"),
            idleTimeout: parseInt(process.env.PROJECT_PRODUCT_MYSQL_CONNECTION_MAX_IDLETIME ?? "0")
        };
        MysqlUtil.pool = mysql.createPool(access);
        console.log(new Date().toISOString(), "mysql: connected to", process.env.PROJECT_PRODUCT_MYSQL_HOST);
    }

    public static async getInstance(): Promise<MysqlUtil> {
        try {        
            if (!MysqlUtil.instance) {
                MysqlUtil.instance = new MysqlUtil()
            } else {
                MysqlUtil.instance
            }
            console.log(new Date().toISOString(), "mysql: pinging to", process.env.PROJECT_PRODUCT_MYSQL_HOST);
            const pool = MysqlUtil.getPool()
            const connection = await pool.getConnection()
            await connection.ping()
            console.log(new Date().toISOString(), "mysql: pingged to", process.env.PROJECT_PRODUCT_MYSQL_HOST);
            return Promise.resolve(MysqlUtil.instance)
        } catch(e: unknown) {
            console.log("error when creating connection mysql to " + process.env.PROJECT_PRODUCT_MYSQL_HOST + " : " + e);
            return Promise.reject(e)
        }
    }

    public static getPool(): Pool {
        return this.pool
    }

    public static async close() {
        if (this.pool) {
            console.log(new Date().toISOString(), "mysql: closing to", process.env.PROJECT_PRODUCT_MYSQL_HOST);
            await this.pool.end()
            console.log(new Date().toISOString(), "mysql: closed to", process.env.PROJECT_PRODUCT_MYSQL_HOST);
        }
    }
}