import { setInternalServerErrorMessage } from "../../../../../src/commons/exceptions/error-exception.js";
import { ResponseException } from "../../../../../src/commons/exceptions/response-exception.js";
import mysqlUtil from "../../../../../src/commons/utils/mysql-util.js";
import getCartService from "../../../../../src/features/get-cart/services/get-cart-service.js";
import getCartRepository from "../../../../../src/features/get-cart/repositories/get-cart-repository.js";
jest.mock("../../../../../src/commons/utils/mysql-util.js")
jest.mock("../../../../../src/features/get-cart/repositories/get-cart-repository.js")

describe("get cart", () => {

    const requestId = "requestId"
    let request = {
        userId: 1
    }

    beforeAll(async () => {

    })

    beforeEach(() => {
        request = {
            userId: 1
        }
    })

    afterEach(() => {
        jest.resetAllMocks();
    })

    afterAll(async () => {

    })

    it("should return validation error when request body is empty", async() => {
        request = {
            userId: 0
        }
        await expect(async () => await getCartService.getCart(requestId, request)).rejects.toThrow("validation error");
    })

    it("should return internal server error when cannot get connection from connection pooling", async() => {
        await mysqlUtil.getConnection.mockImplementation(() => {
            throw new ResponseException(500, setInternalServerErrorMessage(), "internal server error")
        })
        await expect(async () => await getCartService.getCart(requestId, request)).rejects.toThrow("internal server error");
    })

    it("should return internal server error when getting cart by user id", async() => {
        await getCartRepository.getCartByUserId.mockImplementation((connection, userid) => {
            throw new Error("internal server error")
        })
        await expect(async () => await getCartService.getCart(requestId, request)).rejects.toThrow("internal server error");
    })

    it("should return success", async() => {
        getCartRepository.getCartByUserId.mockImplementation(() => {
            return [[{id: 1, user_id: 1, product_id:1, quantity: 1}], []]
        })
        await expect(getCartService.getCart(requestId, request)).resolves.toEqual([{"productId": 1, "quantity": 1, "userId": 1}]);
    })
})