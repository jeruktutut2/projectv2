import supertest from "supertest";
import { web } from "../../../../src/commons/setups/express.js";
import mysqlUtil from "../../../../src/commons/utils/mysql-util.js";
import initializeCart from "../../../initialize/cart.js";
import initializeProduct from "../../../initialize/product.js";
import initializeUser from "../../../initialize/user.js";

describe("get cart GET /api/v1/carts", () => {
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

    it("should return internal server error when there is no table exists", async() => {
        const connection = await mysqlUtil.getConnection(mysqlUtil.mysqlPool)
        await initializeCart.dropTableCart(connection)
        await initializeProduct.dropTableProduct(connection)
        await initializeUser.dropUserTable(connection)
        const response = await supertest(web)
        .get("/api/v1/carts")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({
            userId: 1
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
        await initializeProduct.createTableProduct(connection)
        await initializeCart.createTableCart(connection)
        await initializeUser.createDataUser(connection)
        await initializeProduct.createDataProduct(connection)
        await initializeCart.createDataCart(connection)
        const response = await supertest(web)
        .get("/api/v1/carts")
        .set("Content-Type", "application/json")
        .set("X-REQUEST-ID", "requestId")
        .send({
            userId: 1
        })
        expect(response.status).toEqual(200)
        expect(response.body.data).toEqual([
            { userId: 1, productId: 1, quantity: 1 },
            { userId: 1, productId: 2, quantity: 2 }
          ])
        expect(response.body.errors).toEqual(null)
    })
})