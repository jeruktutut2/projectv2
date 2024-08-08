import updateQuantityService from "../../../../../src/features/update-quantity/services/update-quantity-service.js";
import mysqlUtil from "../../../../../src/commons/utils/mysql-util.js";
import { ResponseException } from "../../../../../src/commons/exceptions/response-exception";
import { setInternalServerErrorMessage } from "../../../../../src/commons/exceptions/error-exception";
import updateQuantityRepository from "../../../../../src/features/update-quantity/repositories/update-quantity-repository.js";
jest.mock("../../../../../src/commons/utils/mysql-util.js")
jest.mock("../../../../../src/features/update-quantity/repositories/update-quantity-repository.js")

describe("update quantity", () => {
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

    it("should return validation error when request body is empty", async() => {
        request = {
            userId: 0,
            productId: 0,
            quantity: 0
        }
        await expect(async () => await updateQuantityService.updateQuantity(requestId, request)).rejects.toThrow("validation error");
    })

    it("should return internal server error when cannot get connection from connection pooling", async() => {
        mysqlUtil.getConnection.mockImplementation(() => {
            throw new ResponseException(500, setInternalServerErrorMessage(), "internal server error")
        })
        await expect(async () => await updateQuantityService.updateQuantity(requestId, request)).rejects.toThrow("internal server error");
    })

    it("should return internal server error when getting by user id and product id", async() => {
        updateQuantityRepository.getByUserIdAndProductId.mockImplementation((connection, userId, productId) => {
            throw new Error("internal server error")
        })
        await expect(async () => await updateQuantityService.updateQuantity(requestId, request)).rejects.toThrow("internal server error");
    })
})