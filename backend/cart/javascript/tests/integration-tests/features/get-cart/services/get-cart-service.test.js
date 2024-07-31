import mysqlUtil from "../../../../../src/commons/utils/mysql-util.js";
import initializeCart from "../../../../initialize/cart.js";
import initializeProduct from "../../../../initialize/product.js";
import initializeUser from "../../../../initialize/user.js";
import getCartService from "../../../../../src/features/get-cart/services/get-cart-service.js";

describe("", () => {

    const requestId = "requestId"
    const userId = 1

    beforeAll(async () => {
        process.env.PROJECT_CART_MYSQL_USERNAME = "root"
        process.env.PROJECT_CART_MYSQL_PASSWORD = "12345"
        process.env.PROJECT_CART_MYSQL_HOST = "localhost:3309"
        process.env.PROJECT_CART_MYSQL_DATABASE = "project_carts"
        mysqlUtil.mysqlPool = await mysqlUtil.newConnection()
    })

    beforeEach(() => {

    })

    afterEach(() => {
        jest.resetAllMocks();
    })

    afterAll(async () => {
        await mysqlUtil.closeConnection(mysqlUtil.mysqlPool)
    })

    it("should return internal server error when there is no table exists", async() => {
        const connection = await mysqlUtil.getConnection(mysqlUtil.mysqlPool)
        await initializeCart.dropTableCart(connection)
        await initializeProduct.dropTableProduct(connection)
        await initializeUser.dropUserTable(connection)
        await expect(async () => await getCartService.getCart(requestId, userId)).rejects.toThrow("internal server error");
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
        const result = await getCartService.getCart(requestId, userId)
        // console.log("result:", result);
        expect(result.length).toEqual(2)
    })
})