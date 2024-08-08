import supertest from "supertest"
import { ElasticsearchUtil } from "../../../../src/utils/elasticsearch-util"
import { MysqlUtil } from "../../../../src/utils/mysql-utils"
import { web } from "../../../../src/setups/express"
import { createDataProductsElasticsearch } from "../../../initialize/products"

describe("search product GET /api/v1/products/search", () => {
    beforeAll(async () => {
        await MysqlUtil.getInstance()
        await ElasticsearchUtil.getInstance()
    })

    beforeEach(() => {

    })

    afterEach(() => {

    })

    afterAll(async () => {
        await MysqlUtil.close()
        await ElasticsearchUtil.closeClient()
    })

    it("should return not found when data don't exists with certain keyword", async () => {
        const client = ElasticsearchUtil.getClient()
        await createDataProductsElasticsearch(client)

        const response = await supertest(web)
        .get("/api/v1/products/search")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({
            keyword: "nam"
        })
        expect(response.status).toEqual(404)
        expect(response.body.data).toEqual(null)
        expect(response.body.errors).toEqual([ { field: 'message', message: 'cannot find data with keyword: nam' } ])
    })

    it("should success", async() => {
        const client = ElasticsearchUtil.getClient()
        await createDataProductsElasticsearch(client)

        const response = await supertest(web)
        .get("/api/v1/products/search")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({
            keyword: "name1"
        })
        expect(response.status).toEqual(200)
        expect(response.body.data).toEqual([
            { id: '4', userId: '1', name: 'name1', description: 'description1' },
            {
              id: '1',
              userId: '1',
              name: 'name1',
              description: 'description1',
              stock: 1
            }
          ])
        expect(response.body.errors).toEqual(null)
    })
})