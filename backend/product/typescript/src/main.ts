import { web } from "./commons/setups/express";
import { MysqlUtil } from "./commons/utils/mysql-utils";
import { ElasticsearchUtil } from "./commons/utils/elasticsearch-util";

const app = web.listen(process.env.PROJECT_PRODUCT_APPLICATION_HOST, async() => {
    try {
        await MysqlUtil.getInstance()
        await ElasticsearchUtil.getInstance()
        console.log(new Date().toISOString(), "express: listening on " + process.env.PROJECT_PRODUCT_APPLICATION_HOST);
    } catch(e) {
        await MysqlUtil.close()
        process.exit(1);
    } finally {

    }
})

process.on("SIGINT", async () => {
    await MysqlUtil.close()
    app.close(() => {
        console.log(new Date().toISOString(), "express: closed on " + process.env.PROJECT_PRODUCT_APPLICATION_HOST);
    })
    process.exit(0)
})