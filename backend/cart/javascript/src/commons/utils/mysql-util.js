import mysql from "mysql2/promise";

// why use this, because i don't know why this variable cannot be accesses in this file, it should be pass in parameter
let mysqlPool
const newConnection = async () => {
    try {
        console.log(new Date().toISOString(), "mysql: connecting to " + process.env.PROJECT_CART_MYSQL_HOST);
        const pool = mysql.createPool("mysql://"+process.env.PROJECT_CART_MYSQL_USERNAME+":"+process.env.PROJECT_CART_MYSQL_PASSWORD+"@"+process.env.PROJECT_CART_MYSQL_HOST+"/"+process.env.PROJECT_CART_MYSQL_DATABASE);
        console.log(new Date().toISOString(), "mysql: connected to " + process.env.PROJECT_CART_MYSQL_HOST);

        console.log(new Date().toISOString(), "mysql: pinging to " + process.env.PROJECT_CART_MYSQL_HOST);
        const connection = await pool.getConnection()
        await connection.ping()
        console.log(new Date().toISOString(), "mysql: pinged to " + process.env.PROJECT_CART_MYSQL_HOST);
        return pool
    } catch(e) {
        console.log(new Date().toISOString(), "mysql: error when connecting to " + process.env.PROJECT_CART_MYSQL_HOST, e);
    }
}

const getConnection = async (pool) => {
    try {
        return pool.getConnection()
    } catch(e) {
        console.log(new Date().toISOString(), "mysql: error when get connection to " + process.env.PROJECT_CART_MYSQL_HOST, e);
    }
}

// why do this, because unit test has to mock this
// but it doesn't work, i don't know, if i move begin to this function, it doesn't return the begin connection
const beginTransaction = async(connection) => {
    return await connection.beginTransaction()
}

const commit = async(connection) => {
    return await connection.commit()
}

const rollback = async(connection) => {
    return await connection.rollback()
}

const releaseConnection = async(pool, connection) => {
    try {
        return await pool.releaseConnection(connection)    
    } catch(e) {
        console.log(new Date().toISOString(), "mysql: error when release connection to " + process.env.PROJECT_CART_MYSQL_HOST, e);
    }
}

const closeConnection = async (pool) => {
    try {
        console.log(new Date().toISOString(), "mysql: closing to " + process.env.PROJECT_CART_MYSQL_HOST);
        await pool.end()
        console.log(new Date().toISOString(), "mysql: closed to " + process.env.PROJECT_CART_MYSQL_HOST);
    } catch(e) {
        console.log(new Date().toISOString(), "mysql: error when end connection to " + process.env.PROJECT_CART_MYSQL_HOST, e);
    }
}

export default {
    mysqlPool,
    newConnection,
    getConnection,
    beginTransaction,
    commit,
    rollback,
    releaseConnection,
    closeConnection
}