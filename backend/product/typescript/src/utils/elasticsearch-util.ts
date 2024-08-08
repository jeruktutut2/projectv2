import { Client } from "@elastic/elasticsearch";

export class ElasticsearchUtil {

    private static instance: ElasticsearchUtil
    private static client: Client

    private constructor() {
        console.log(new Date().toISOString(), "elasticsearch: connecting to", process.env.PROJECT_PRODUCT_ELASTICSEARCH_NODE);
        ElasticsearchUtil.client = new Client({
            node: process.env.PROJECT_PRODUCT_ELASTICSEARCH_NODE
        })
        console.log(new Date().toISOString(), "elasticsearch: connected to", process.env.PROJECT_PRODUCT_ELASTICSEARCH_NODE);
    }

    public static async getInstance(): Promise<ElasticsearchUtil> {
        try {
            if (!ElasticsearchUtil.instance) {
                ElasticsearchUtil.instance = new ElasticsearchUtil()
            } else {
                ElasticsearchUtil.instance
            }
            console.log(new Date().toISOString(), "elasticsearch: pinging to", process.env.PROJECT_PRODUCT_ELASTICSEARCH_NODE);
            await ElasticsearchUtil.client.ping()
            console.log(new Date().toISOString(), "elasticsearch: pinged to", process.env.PROJECT_PRODUCT_ELASTICSEARCH_NODE);
            return Promise.resolve(ElasticsearchUtil.instance)
        } catch(e: unknown) {
            console.log("error when creating connection elasticsearch to " + process.env.PROJECT_PRODUCT_ELASTICSEARCH_NODE + " : " + e);
            return Promise.reject(e)
        }
    }

    public static getClient(): Client {
        return this.client
    }

    public static async closeClient() {
        if (ElasticsearchUtil.client) {
            console.log(new Date().toISOString(), "elasticsearch: closing to", process.env.PROJECT_PRODUCT_ELASTICSEARCH_NODE);
            await ElasticsearchUtil.client.close()
            console.log(new Date().toISOString(), "elasticsearch: closed to", process.env.PROJECT_PRODUCT_ELASTICSEARCH_NODE);
        }
    }
}