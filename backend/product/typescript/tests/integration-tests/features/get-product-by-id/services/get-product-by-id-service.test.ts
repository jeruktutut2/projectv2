import { MysqlUtil } from "../../../../../src/commons/utils/mysql-utils"
import { GetProductByIdService } from "../../../../../src/features/get-product-by-id/services/get-product-by-id-service";
import { createDataProducts, createTableProducts, deleteTableProducts } from "../../../../initialize/products";
import { PoolConnection } from "mysql2/promise";

describe("get product by id", () => {

    const requestId = "requestId"
    const id = 1

    beforeAll( async () => {
        await MysqlUtil.getInstance()
    })

    beforeEach(() => {

    })

    afterEach(() => {
        jest.resetAllMocks();
    })

    afterAll( async () => {
        await MysqlUtil.close()
    })

    it("should return internal server error", async () => {
        const poolConnection: PoolConnection = await MysqlUtil.getPool().getConnection()
        await deleteTableProducts(poolConnection)
        await expect(async () => await GetProductByIdService.getProductById(requestId, id)).rejects.toThrow("internal server error")
    })

    it("should return error when no product", async () => {
        const poolConnection: PoolConnection = await MysqlUtil.getPool().getConnection()
        await deleteTableProducts(poolConnection)
        await createTableProducts(poolConnection)
        await expect(async () => await GetProductByIdService.getProductById(requestId, id)).rejects.toThrow("cannot find product with id:1")
    })

    it("should success", async () => {
        const poolConnection: PoolConnection = await MysqlUtil.getPool().getConnection()
        await deleteTableProducts(poolConnection)
        await createTableProducts(poolConnection)
        await createDataProducts(poolConnection)
        await expect(GetProductByIdService.getProductById(requestId, id)).resolves.toEqual({"description": "description1", "id": 1, "name": "name1", "stoct": 1})
    })
})