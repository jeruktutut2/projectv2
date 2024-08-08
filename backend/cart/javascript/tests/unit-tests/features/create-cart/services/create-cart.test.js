import createCartService from "../../../../../src/features/create-cart/services/create-cart-service.js";
import mysqlUtil from "../../../../../src/commons/utils/mysql-util.js";
import { ResponseException } from "../../../../../src/commons/exceptions/response-exception";
import { setInternalServerErrorMessage } from "../../../../../src/commons/exceptions/error-exception";
import createCartRepository from "../../../../../src/features/create-cart/repositories/create-cart-repository.js";
jest.mock("../../../../../src/commons/utils/mysql-util.js")
jest.mock("../../../../../src/features/create-cart/repositories/create-cart-repository.js")

describe("create cart", () => {
    const requestId = "requestId"
    let request = {
        userId: 1,
        productId: 1,
        quantity: 1
    }

    beforeAll(async () => {

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

    })

    it("should return bad request when request is empty", async() => {
        request = {
            userId: 0,
            productId: 0,
            quantity: 0
        }
        await expect(async () => await createCartService.createCart(requestId, request)).rejects.toThrow("validation error");
    })

    it("should return internal server error when cannot get connection from connection pooling", async() => {
        await mysqlUtil.getConnection.mockImplementation(() => {
            throw new ResponseException(500, setInternalServerErrorMessage(), "internal server error")
        })
        await expect(async () => await createCartService.createCart(requestId, request)).rejects.toThrow("internal server error");
    })

    it("should return internal server when creating cart", async() => {
        createCartRepository.createCart.mockImplementation((connection, cart) => {
            throw new Error("internal server error")
        })
        await expect(async () => await createCartService.createCart(requestId, request)).rejects.toThrow("internal server error");
    })

    it("should return internal server error because rows affected not one", async() => {
        createCartRepository.createCart.mockImplementation((connection, cart) => {
            return [[], []]
        })
        await expect(async () => await createCartService.createCart(requestId, request)).rejects.toThrow("internal server error");
    })

    it("should success", async() => {
        await createCartRepository.createCart.mockImplementation((connection, cart) => {
            return [
                {
                  fieldCount: 0,
                  affectedRows: 1,
                  insertId: 1,
                  info: '',
                  serverStatus: 2,
                  warningStatus: 0,
                  changedRows: 0
                },
                undefined
              ]
        })
        await expect(createCartService.createCart(requestId, request)).resolves.toEqual({"userId": 1, "productId": 1, "quantity": 1});
    })
})