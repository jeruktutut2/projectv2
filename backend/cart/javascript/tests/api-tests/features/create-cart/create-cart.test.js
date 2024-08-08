import supertest from "supertest";
import { web } from "../../../../src/commons/setups/express.js";
import mysqlUtil from "../../../../src/commons/utils/mysql-util.js";
import initializeCart from "../../../initialize/cart.js";
import initializeProduct from "../../../initialize/product.js";
import initializeUser from "../../../initialize/user.js";

describe("create cart POST /api/v1/carts", () => {
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

    it("should return when request body is empty", async() => {
        const response = await supertest(web)
        .post("/api/v1/carts")
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
        .post("/api/v1/carts")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({
            userId: 1,
            productId: 1,
            quantity: 1
        })
        expect(response.status).toEqual(500)
        expect(response.body.data).toEqual(null)
        expect(response.body.errors).toEqual([ { field: 'message', message: 'internal server error' } ])
    })

    it("should success", async() => {
        const connection = await mysqlUtil.getConnection(mysqlUtil.mysqlPool)
        await initializeCart.dropTableCart(connection)
        await initializeProduct.dropTableProduct(connection)
        await initializeUser.dropUserTable(connection)
        await initializeUser.createTableUser(connection)
        await initializeUser.createDataUser(connection)
        await initializeProduct.createTableProduct(connection)
        await initializeProduct.createDataProduct(connection)
        await initializeCart.createTableCart(connection)

        const response = await supertest(web)
        .post("/api/v1/carts")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({
            userId: 1,
            productId: 1,
            quantity: 1
        })

        const [rows] = await initializeCart.getDataCart(connection, 1)
        expect(rows.length).toEqual(1)
        expect(rows[0].id).toEqual(1)
        expect(rows[0].user_id).toEqual(1)
        expect(rows[0].product_id).toEqual(1)
        expect(rows[0].quantity).toEqual(1)

        expect(response.status).toEqual(201)
        expect(response.body.data).toEqual({ userId: 1, productId: 1, quantity: 1 })
        expect(response.body.errors).toEqual(null)
    })
})