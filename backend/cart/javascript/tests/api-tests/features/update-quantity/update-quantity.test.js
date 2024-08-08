import supertest from "supertest";
import { web } from "../../../../src/commons/setups/express.js";
import mysqlUtil from "../../../../src/commons/utils/mysql-util.js";
import initializeCart from "../../../initialize/cart.js";
import initializeProduct from "../../../initialize/product.js";
import initializeUser from "../../../initialize/user.js";

describe("update quantity PATCH /api/v1/carts/update-quantity", () => {
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
        .patch("/api/v1/carts/update-quantity")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({})
        expect(response.status).toEqual(400)
        expect(response.body.data).toEqual(null)
        expect(response.body.errors).toEqual([
            { field: 'userId', message: 'userId is required' },
            { field: 'productId', message: 'productId is required' },
            { field: 'quantity', message: 'quantity is required' }
        ])
    })

    it("should return internal server error when there is no table exists", async() => {
        const connection = await mysqlUtil.getConnection(mysqlUtil.mysqlPool)
        await initializeCart.dropTableCart(connection)
        await initializeProduct.dropTableProduct(connection)
        await initializeUser.dropUserTable(connection)
        const response = await supertest(web)
        .patch("/api/v1/carts/update-quantity")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({
            userId: 1,
            productId: 1,
            quantity: 2
        })
        expect(response.status).toEqual(500)
        expect(response.body.data).toEqual(null)
        expect(response.body.errors).toEqual([ { field: 'message', message: 'internal server error' } ])
    })

    it("should return bad requst when cart data doesn't exists", async() => {
        const connection = await mysqlUtil.getConnection(mysqlUtil.mysqlPool)
        await initializeCart.dropTableCart(connection)
        await initializeProduct.dropTableProduct(connection)
        await initializeUser.dropUserTable(connection)
        await initializeUser.createTableUser(connection)
        await initializeProduct.createTableProduct(connection)
        await initializeCart.createTableCart(connection)
        const response = await supertest(web)
        .patch("/api/v1/carts/update-quantity")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({
            userId: 1,
            productId: 1,
            quantity: 2
        })
        expect(response.status).toEqual(400)
        expect(response.body.data).toEqual(null)
        expect(response.body.errors).toEqual([ { field: 'message', message: 'cannot find product in cart' } ])
    })

    it("should return success", async() => {
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

        const [rows1] = await initializeCart.getDataCart(connection, 1)
        expect(rows1[0].quantity).toEqual(1)
        const response = await supertest(web)
        .patch("/api/v1/carts/update-quantity")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({
            userId: 1,
            productId: 1,
            quantity: 2
        })
        const [rows2] = await initializeCart.getDataCart(connection, 1)
        expect(rows2[0].quantity).toEqual(2)

        expect(response.status).toEqual(200)
        expect(response.body.data).toEqual({ userId: 1, productId: 1, quantity: 2 })
        expect(response.body.errors).toEqual(null)
    })
})