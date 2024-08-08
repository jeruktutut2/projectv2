import supertest from "supertest"
import { web} from "../../../../src/setups/express";
import { MysqlUtil } from "../../../../src/utils/mysql-utils";
import { ElasticsearchUtil } from "../../../../src/utils/elasticsearch-util";
import { createDataProducts, createTableProducts, deleteTableProducts, getDataProduct } from "../../../initialize/products";

describe("create product POST /api/v1/products", () => {

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

    it("should return error bad request when request body is empty", async() => {
        const response = await supertest(web)
        .post("/api/v1/products")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({})
        expect(response.status).toEqual(400)
        expect(response.body.data).toEqual(null)
        expect(response.body.errors).toEqual([
            { field: 'userId', message: 'userId is required' },
            { field: 'name', message: 'name is required' },
            { field: 'description', message: 'description is required' },
            { field: 'stock', message: 'stock is required' }
          ])
    }, 60000)

    it("should return internal server error when there is no table exists", async() => {
        const poolConnection = await MysqlUtil.getPool().getConnection()
        await deleteTableProducts(poolConnection)
        const response = await supertest(web)
        .post("/api/v1/products")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({
            userId:1,
            name: "name1",
            description: "description1",
            stock: 1
        })
        expect(response.status).toEqual(500)
        expect(response.body.data).toEqual(null)
        expect(response.body.errors).toEqual([ { field: 'message', message: 'internal server error' } ])
    }, 60000)
    
    it("should success", async() => {
        const poolConnection = await MysqlUtil.getPool().getConnection()
        await deleteTableProducts(poolConnection)
        await createTableProducts(poolConnection)
        await createDataProducts(poolConnection)
        const response = await supertest(web)
        .post("/api/v1/products")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({
            userId:1,
            name: "name1",
            description: "description1",
            stock: 1
        })
        expect(response.status).toEqual(201)
        expect(response.body.data).toEqual({ name: 'name1', description: 'description1', stock: 1 })
        expect(response.body.errors).toEqual(null)
        
        const [rows] = await  getDataProduct(poolConnection, 1)

        const resultElasticsearch = await ElasticsearchUtil.getClient().get({
            index: "products",
            id: "1"
        })
        expect(resultElasticsearch.found).toEqual(true)

        expect(rows.length).toEqual(1)
        expect(rows[0].name).toEqual("name1")
        expect(rows[0].description).toEqual("description1")
        expect(rows[0].stock).toEqual(1)
        
    })
})