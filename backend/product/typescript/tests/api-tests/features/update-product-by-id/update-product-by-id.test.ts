import supertest from "supertest"
import { ElasticsearchUtil } from "../../../../src/commons/utils/elasticsearch-util"
import { MysqlUtil } from "../../../../src/commons/utils/mysql-utils"
import { web } from "../../../../src/commons/setups/express";
import { createDataProducts, createDataProductsElasticsearch, createTableProducts, deleteTableProducts, getDataProduct } from "../../../initialize/products";

describe("update product by id PATCH /api/v1/products", () => {
    beforeAll( async () => {
        await MysqlUtil.getInstance()
        await ElasticsearchUtil.getInstance()
    })

    beforeEach(() => {

    })

    afterEach(() => {

    })

    afterAll( async () => {
        await MysqlUtil.close()
        await ElasticsearchUtil.closeClient()
    })

    it("should return error when update product by id request is empty", async() => {
        const response = await supertest(web)
        .patch("/api/v1/products")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({})
        expect(response.status).toEqual(400)
        expect(response.body.data).toEqual(null)
        expect(response.body.errors).toEqual([
            { field: 'id', message: 'id is required' },
            { field: 'name', message: 'name is required' },
            { field: 'description', message: 'description is required' }
        ])
    })

    it("should return internal server error", async() => {
        const poolConnection = await MysqlUtil.getPool().getConnection()
        await deleteTableProducts(poolConnection)

        const response = await supertest(web)
        .patch("/api/v1/products")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({
            id: 1,
            name: "name edit",
            description: "description edit"
        })
        expect(response.status).toEqual(500)
        expect(response.body.data).toEqual(null)
        expect(response.body.errors).toEqual([ { field: 'message', message: 'internal server error' } ])
    })

    it("should return response exception when there is no data product", async() => {
        const poolConnection = await MysqlUtil.getPool().getConnection()
        await deleteTableProducts(poolConnection)
        await createTableProducts(poolConnection)
        const response = await supertest(web)
        .patch("/api/v1/products")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({
            id: 1,
            name: "name edit",
            description: "description edit"
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
        const client = ElasticsearchUtil.getClient()
        await createDataProductsElasticsearch(client)

        const response = await supertest(web)
        .patch("/api/v1/products")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({
            id: 1,
            name: "name edit",
            description: "description edit"
        })

        const [rows] = await getDataProduct(poolConnection, 1)
        expect(rows.length).toEqual(1)
        expect(rows[0].name).toEqual("name edit")
        expect(rows[0].description).toEqual("description edit")

        const resultElasticsearch = await ElasticsearchUtil.getClient().get({
            index: "products",
            id: "1"
        })
        expect(resultElasticsearch.found).toEqual(true)

        expect(response.status).toEqual(200)
        expect(response.body.data).toEqual({ name: 'name edit', description: 'description edit' })
        expect(response.body.errors).toEqual(null)
    })
})