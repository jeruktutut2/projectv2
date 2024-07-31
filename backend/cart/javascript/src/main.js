import mysqlUtil from "./commons/utils/mysql-util.js";
import { web } from "./commons/setups/express.js";

mysqlUtil.mysqlPool = await mysqlUtil.newConnection()
const app = web.listen(process.env.PROJECT_CART_APPLICATION_HOST, () => {
    console.log(new Date().toISOString(), "express: listening on " + process.env.PROJECT_CART_APPLICATION_HOST);
})

process.on("SIGINT", async () => {
    app.close(() => {
        console.log(new Date().toISOString(), "express: closed on " + process.env.PROJECT_CART_APPLICATION_HOST);
    })
    await mysqlUtil.closeConnection(mysqlUtil.mysqlPool)
    process.exit(0)
})
