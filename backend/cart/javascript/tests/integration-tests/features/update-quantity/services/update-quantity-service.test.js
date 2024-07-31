import updateQuantityService from "../../../../../src/features/update-quantity/services/update-quantity-service.js";
import mysqlUtil from "../../../../../src/commons/utils/mysql-util.js";
import initializeCart from "../../../../initialize/cart.js";
import initializeProduct from "../../../../initialize/product.js";
import initializeUser from "../../../../initialize/user.js";

describe("", () => {

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
            quantity: 2
        }
    })

    afterEach(() => {
        jest.resetAllMocks();
    })

    afterAll(async () => {
        await mysqlUtil.closeConnection(mysqlUtil.mysqlPool)
    })

    it("should return validation error when request is empty", async() => {
        request = {
            userId: 0,
            productId: 0,
            quantity: 0
        }
        await expect(async () => await updateQuantityService.updateQuantity(requestId, request)).rejects.toThrow("validation error");
    })

    it("should return internal server error when there is no table exists", async () => {
        const connection = await mysqlUtil.getConnection(mysqlUtil.mysqlPool)
        await initializeCart.dropTableCart(connection)
        await initializeProduct.dropTableProduct(connection)
        await initializeUser.dropUserTable(connection)
        await expect(async () => await updateQuantityService.updateQuantity(requestId, request)).rejects.toThrow("internal server error");
    })

    it("should return bad request when cart data doesn't exists", async() => {
        const connection = await mysqlUtil.getConnection(mysqlUtil.mysqlPool)
        await initializeCart.dropTableCart(connection)
        await initializeProduct.dropTableProduct(connection)
        await initializeUser.dropUserTable(connection)
        await initializeUser.createTableUser(connection)
        await initializeProduct.createTableProduct(connection)
        await initializeCart.createTableCart(connection)
        await expect(async () => await updateQuantityService.updateQuantity(requestId, request)).rejects.toThrow("cannot find product in cart");
    })

    it("shoult return success", async() => {
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

        const  [rows1] = await initializeCart.getDataCart(connection, 1)
        // console.log("rows1:", rows1);
        expect(rows1[0].quantity).toEqual(1)

        const result = await updateQuantityService.updateQuantity(requestId, request)

        const [rows2] = await initializeCart.getDataCart(connection, 1)
        expect(rows2[0].quantity).toEqual(request.quantity)

        expect(result.userId).toEqual(request.userId)
        expect(result.productId).toEqual(request.productId)
        expect(result.quantity).toEqual(request.quantity)
    })
})