import supertest from "supertest"
import { ElasticsearchUtil } from "../../../../src/commons/utils/elasticsearch-util"
import { MysqlUtil } from "../../../../src/commons/utils/mysql-utils"
import { web } from "../../../../src/commons/setups/express"
import { createDataProducts, createDataProductsElasticsearch, createTableProducts, deleteTableProducts, getDataProduct, getDataProductsElasticsearch } from "../../../initialize/products"

describe("delete product DELETE /api/v1//products", () => {
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

    it("should return internal server error when no table exists", async () => {
        const poolConnection = await MysqlUtil.getPool().getConnection()
        await deleteTableProducts(poolConnection)
        
        const response = await supertest(web)
        .delete("/api/v1/products")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({
            id: 1
        })
        expect(response.status).toEqual(500)
        expect(response.body.data).toEqual(null)
        expect(response.body.errors).toEqual([ { field: 'message', message: 'internal server error' } ])
    })

    it("should return number affected rows is not one", async() => {
        const poolConnection = await MysqlUtil.getPool().getConnection()
        await deleteTableProducts(poolConnection)
        await createTableProducts(poolConnection)

        const response = await supertest(web)
        .delete("/api/v1/products")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({
            id: 1
        })
        expect(response.status).toEqual(500)
        expect(response.body.data).toEqual(null)
        expect(response.body.errors).toEqual([
            {
              field: 'message',
              message: 'number of affected rows when deleting product is not one:0'
            }
          ])
    })

    it("should success", async() => {
        const poolConnection = await MysqlUtil.getPool().getConnection()
        await deleteTableProducts(poolConnection)
        await createTableProducts(poolConnection)
        await createDataProducts(poolConnection)
        const client = ElasticsearchUtil.getClient()
        await createDataProductsElasticsearch(client)

        const response = await supertest(web)
        .delete("/api/v1/products")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({
            id: 1
        })

        const [rows] = await getDataProduct(poolConnection, 1)
        expect(rows.length).toEqual(0)

        await expect(getDataProductsElasticsearch(client)).resolves.toEqual(undefined)

        expect(response.status).toEqual(200)
        expect(response.body.data).toEqual({ message: 'successfully delete product' })
        expect(response.body.errors).toEqual(null)
    })
})