import supertest from "supertest";
import { web } from "../../../../src/commons/setups/express.js";
import mysqlUtil from "../../../../src/commons/utils/mysql-util.js";
import initializeCart from "../../../initialize/cart.js";
import initializeProduct from "../../../initialize/product.js";
import initializeUser from "../../../initialize/user.js";

describe("delete cart DELETE /api/v1/carts", () => {
    beforeAll(async () => {
        mysqlUtil.mysqlPool = await mysqlUtil.newConnection()
    })

    beforeEach(() => {

    })

    afterEach(() => {

    })

    afterAll(async () => {
        await mysqlUtil.closeConnection(mysqlUtil.mysqlPool)
    })

    it("should return validation error when request is empty", async() => {
        const response = await supertest(web)
        .delete("/api/v1/carts")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({})
        expect(response.status).toEqual(400)
        expect(response.body.data).toEqual(null)
        expect(response.body.errors).toEqual([
            { field: 'id', message: 'id is required' },
            { field: 'userId', message: 'userId is required' },
            { field: 'productId', message: 'productId is required' }
        ])
    })

    it("should return internal server error when there is no table exists", async() => {
        const connection = await mysqlUtil.getConnection(mysqlUtil.mysqlPool)
        await initializeCart.dropTableCart(connection)
        await initializeProduct.dropTableProduct(connection)
        await initializeUser.dropUserTable(connection)

        const response = await supertest(web)
        .delete("/api/v1/carts")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({
            id: 1,
            userId: 1,
            productId: 1
        })
        expect(response.status).toEqual(500)
        expect(response.body.data).toEqual(null)
        expect(response.body.errors).toEqual([ { field: 'message', message: 'internal server error' } ])
    })

    it("should return bad request cannot find product in cart", async() => {
        const connection = await mysqlUtil.getConnection(mysqlUtil.mysqlPool)
        await initializeCart.dropTableCart(connection)
        await initializeProduct.dropTableProduct(connection)
        await initializeUser.dropUserTable(connection)
        await initializeUser.createTableUser(connection)
        await initializeProduct.createTableProduct(connection)
        await initializeUser.createDataUser(connection)
        await initializeProduct.createDataProduct(connection)
        const response = await supertest(web)
        .delete("/api/v1/carts")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({
            id: 0,
            userId: 0,
            productId: 0
        })
        expect(response.status).toEqual(400)
        expect(response.body.data).toEqual(null)
        expect(response.body.errors).toEqual([
            { field: 'id', message: 'id must be greater than or equal to 1' },
            { field: 'id', message: 'id must be a positive number' },
            {
              field: 'userId',
              message: 'userId must be greater than or equal to 1'
            },
            { field: 'userId', message: 'userId must be a positive number' },
            {
              field: 'productId',
              message: 'productId must be greater than or equal to 1'
            },
            {
              field: 'productId',
              message: 'productId must be a positive number'
            }
        ])
    })

    it("should success", async() => {
        const connection = await mysqlUtil.getConnection(mysqlUtil.mysqlPool)
        await initializeCart.dropTableCart(connection)
        await initializeProduct.dropTableProduct(connection)
        await initializeUser.dropUserTable(connection)
        await initializeUser.createTableUser(connection)
        await initializeProduct.createTableProduct(connection)
        await initializeCart.createTableCart(connection)
        await initializeUser.createDataUser(connection)
        await initializeProduct.createDataProduct(connection)
        await initializeCart.createDataCart(connection)

        const [rows1] =  await initializeCart.getDataCart(connection, 1)
        expect(rows1.length).toEqual(1)

        const response = await supertest(web)
        .delete("/api/v1/carts")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({
            id: 1,
            userId: 1,
            productId: 1
        })

        const [rows2] = await initializeCart.getDataCart(connection, 1)
        expect(rows2.length).toEqual(0)

        expect(response.status).toEqual(200)
        expect(response.body.data).toEqual({ message: 'successfully delete cart' })
        expect(response.body.errors).toEqual(null)
    })
})