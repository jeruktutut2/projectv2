import deleteCartService from "../../../../../src/features/delete-cart/services/delete-cart-service.js";
import mysqlUtil from "../../../../../src/commons/utils/mysql-util.js";
import initializeCart from "../../../../initialize/cart.js";
import initializeProduct from "../../../../initialize/product.js";
import initializeUser from "../../../../initialize/user.js";

describe("delete cart", () => {

    const requestId = "requestId"
    let request = {
        id: 0,
        userId: 0,
        productId: 0
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
            id: 1,
            userId: 1,
            productId: 1
        }
    })

    afterEach(() => {
        jest.resetAllMocks();
    })

    afterAll(async () => {
        await mysqlUtil.closeConnection(mysqlUtil.mysqlPool)
    })

    it("should return validate error when request is empty", async() => {
        request = {
            id: 0,
            userId: 0,
            productId: 0
        }
        await expect(async () => await deleteCartService.deleteCart(requestId, request)).rejects.toThrow("validation error");
    })

    it("should return internal server error because there is no table exists", async () => {
        const connection = await mysqlUtil.getConnection(mysqlUtil.mysqlPool)
        await initializeCart.dropTableCart(connection)
        await initializeProduct.dropTableProduct(connection)
        await initializeUser.dropUserTable(connection)
        await expect(async () => await deleteCartService.deleteCart(requestId, request)).rejects.toThrow("internal server error");
    })

    it("should return bad request cannot find product in cart", async() => {
        const connection = await mysqlUtil.getConnection(mysqlUtil.mysqlPool)
        await initializeCart.dropTableCart(connection)
        await initializeProduct.dropTableProduct(connection)
        await initializeUser.dropUserTable(connection)
        await initializeUser.createTableUser(connection)
        await initializeProduct.createTableProduct(connection)
        await initializeCart.createTableCart(connection)
        await initializeUser.createDataUser(connection)
        await initializeProduct.createDataProduct(connection)
        await expect(async () => await deleteCartService.deleteCart(requestId, request)).rejects.toThrow("cannot find product in cart");
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

        const [rows1] = await initializeCart.getDataCart(connection, request.id)
        expect(rows1.length).toEqual(1)

        const result = await deleteCartService.deleteCart(requestId, request)

        const [rows2] = await initializeCart.getDataCart(connection, request.id)
        expect(rows2.length).toEqual(0)

        expect(result.message).toEqual("successfully delete cart")
    })
})