import mysqlUtil from "../../../../../src/commons/utils/mysql-util.js";
import createCartService from "../../../../../src/features/create-cart/services/create-cart-service.js";
import initializeUser from "../../../../initialize/user.js";
import initializeProduct from "../../../../initialize/product.js";
import initializeCart from "../../../../initialize/cart.js";

describe("create cart service", () => {

    const requestId = "requestId"
    let request = {
        userId: 0,
        productId: 0,
        quantity: 0
    }

    beforeAll(async () => {
        process.env.PROJECT_CART_MYSQL_USERNAME = "root"
        process.env.PROJECT_CART_MYSQL_PASSWORD = "12345"
        process.env.PROJECT_CART_MYSQL_HOST = "localhost:3309"
        process.env.PROJECT_CART_MYSQL_DATABASE = "project_carts"
        mysqlUtil.mysqlPool = await mysqlUtil.newConnection()
    })

    beforeEach(() => {
        request = {
            userId: 1,
            productId: 1,
            quantity: 1
        }
    })

    afterEach(() => {
        jest.resetAllMocks();
    })

    afterAll(async () => {
        await mysqlUtil.closeConnection(mysqlUtil.mysqlPool)
    })

    it("should return validation error when request body is empty", async () => {
        request = {
            userId: 0,
            productId: 0,
            quantity: 0
        }
        await expect(async () => await createCartService.createCart(requestId, request)).rejects.toThrow("validation error");
    })

    it("should return internal server error because there is no table exists", async() => {
        const connection = await mysqlUtil.getConnection(mysqlUtil.mysqlPool)
        await initializeCart.dropTableCart(connection)
        await initializeProduct.dropTableProduct(connection)
        await initializeUser.dropUserTable(connection)
        await expect(async () => await createCartService.createCart(requestId, request)).rejects.toThrow("internal server error");
    })

    it("should return success", async () => {
        const connection = await mysqlUtil.getConnection(mysqlUtil.mysqlPool)
        await initializeCart.dropTableCart(connection)
        await initializeProduct.dropTableProduct(connection)
        await initializeUser.dropUserTable(connection)
        await initializeUser.createTableUser(connection)
        await initializeUser.createDataUser(connection)
        await initializeProduct.createTableProduct(connection)
        await initializeProduct.createDataProduct(connection)
        await initializeCart.createTableCart(connection)

        const result = await createCartService.createCart(requestId, request)

        const [rows] =  await initializeCart.getDataCart(connection, 1)
        expect(rows.length).toEqual(1)
        expect(rows[0].id).toEqual(1)
        expect(rows[0].user_id).toEqual(1)
        expect(rows[0].product_id).toEqual(1)
        expect(rows[0].quantity).toEqual(1)

        expect(result.userId).toEqual(1)
        expect(result.productId).toEqual(1)
        expect(result.quantity).toEqual(1)
    })
})