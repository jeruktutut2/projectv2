import { web } from "./setups/express";
import { MysqlUtil } from "./utils/mysql-utils";
import { ElasticsearchUtil } from "./utils/elasticsearch-util";

// const mysqlUtil = await MysqlUtil.getInstance()

// let mysqlUtil: MysqlUtil
const app = web.listen(process.env.PROJECT_PRODUCT_APPLICATION_HOST, async() => {
    try {
        // mysqlUtil = await MysqlUtil.getInstance()
        await MysqlUtil.getInstance()
        await ElasticsearchUtil.getInstance()
        console.log(new Date().toISOString(), "express: listening on " + process.env.PROJECT_PRODUCT_APPLICATION_HOST);
    } catch(e) {
        console.log("error:", e);
        process.exit(1);
    } finally {
        // if (mysqlUtil) {
        //     mysqlUtil
        // }
        await MysqlUtil.close()
    }
})

process.on("SIGINT", async () => {
    app.close(() => {
        console.log(new Date().toISOString(), "express: closed on " + process.env.PROJECT_PRODUCT_APPLICATION_HOST);
    })
    await MysqlUtil.close()
    process.exit(0)
})