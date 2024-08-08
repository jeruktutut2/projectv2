import supertest from "supertest"
import { web } from "../../../../src/setups/express"
import { MysqlUtil } from "../../../../src/utils/mysql-utils"
import { createDataProducts, createTableProducts, deleteTableProducts } from "../../../initialize/products"
import { ElasticsearchUtil } from "../../../../src/utils/elasticsearch-util"

describe("get product by id GET /api/v1/products", () => {
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

    it("should return internal server error when there is no table exists", async() => {
        const poolConnection = await MysqlUtil.getPool().getConnection()
        await deleteTableProducts(poolConnection)

        const response = await supertest(web)
        .get("/api/v1/products")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({
            id: 1
        })
        expect(response.status).toEqual(500)
        expect(response.body.data).toEqual(null)
        expect(response.body.errors).toEqual([ { field: 'message', message: 'internal server error' } ])
    })

    it("should return error when there is no product", async() => {
        const poolConnection = await MysqlUtil.getPool().getConnection()
        await deleteTableProducts(poolConnection)
        await createTableProducts(poolConnection)

        const response = await supertest(web)
        .get("/api/v1/products")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({
            id: 1
        })
        expect(response.status).toEqual(400)
        expect(response.body.data).toEqual(null)
        expect(response.body.errors).toEqual([ { field: 'message', message: 'cannot find product with id:1' } ])
    })

    it("should success", async() => {
        const poolConnection = await MysqlUtil.getPool().getConnection()
        await deleteTableProducts(poolConnection)
        await createTableProducts(poolConnection)
        await createDataProducts(poolConnection)

        const response = await supertest(web)
        .get("/api/v1/products")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({
            id: 1
        })
        expect(response.status).toEqual(200)
        expect(response.body.data).toEqual({ id: 1, name: 'name1', description: 'description1', stoct: 1 })
        expect(response.body.errors).toEqual(null)
    })
})