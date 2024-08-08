import { MysqlUtil } from "../../../../../src/utils/mysql-utils"
import { ElasticsearchUtil } from "../../../../../src/utils/elasticsearch-util";
import { UpdateProductByIdRequest } from "../../../../../src/features/update-product-by-id/models/update-product-by-id-request";
import { UpdateProductByIdService } from "../../../../../src/features/update-product-by-id/services/update-product-by-id-service";
import { PoolConnection } from "mysql2/promise";
import { createDataProducts, createDataProductsElasticsearch, createTableProducts, deleteTableProducts, getDataProduct } from "../../../../initialize/products";

describe("update product by id", () => {

    const requestId = "requestId"
    let updateProductByIdRequest: UpdateProductByIdRequest = {
        id: 1,
        name: "name edit",
        description: "description edit"
    }

    beforeAll( async () => {
        await MysqlUtil.getInstance()
        await ElasticsearchUtil.getInstance()
    })

    beforeEach(() => {
        updateProductByIdRequest = {
            id: 1,
            name: "name edit",
            description: "description edit"
        }
    })

    afterEach(() => {
        jest.resetAllMocks();
    })

    afterAll( async () => {
        await MysqlUtil.close()
        await ElasticsearchUtil.closeClient()
    })

    it("should return error when update product by id request is empty", async () => {
        updateProductByIdRequest = {
            id: 0,
            name: "",
            description: "'"
        }
        await expect(async () => await UpdateProductByIdService.updateProductById(requestId, updateProductByIdRequest)).rejects.toThrow("validation error")
    })

    it("should return internal server error", async () => {
        const poolConnection: PoolConnection = await MysqlUtil.getPool().getConnection()
        await deleteTableProducts(poolConnection)
        await expect(async () => await UpdateProductByIdService.updateProductById(requestId, updateProductByIdRequest)).rejects.toThrow("internal server error")
    })

    it("should return response exception when there is no data product", async () => {
        const poolConnection: PoolConnection = await MysqlUtil.getPool().getConnection()
        await deleteTableProducts(poolConnection)
        await createTableProducts(poolConnection)
        await expect(async () => await UpdateProductByIdService.updateProductById(requestId, updateProductByIdRequest)).rejects.toThrow("cannot find product with id:1")
    })

    it("should success", async () => {
        const poolConnection: PoolConnection = await MysqlUtil.getPool().getConnection()
        await deleteTableProducts(poolConnection)
        await createTableProducts(poolConnection)
        await createDataProducts(poolConnection)

        const client = ElasticsearchUtil.getClient()
        await createDataProductsElasticsearch(client)

        const result = await UpdateProductByIdService.updateProductById(requestId, updateProductByIdRequest)
        const [rows] = await getDataProduct(poolConnection, 1)

        expect(rows.length).toEqual(1)
        expect(rows[0].name).toEqual("name edit")
        expect(rows[0].description).toEqual("description edit")

        const resultElasticsearch = await ElasticsearchUtil.getClient().get({
            index: "products",
            id: "1"
        })
        expect(resultElasticsearch.found).toEqual(true)

        expect(result.name).toEqual("name edit")
        expect(result.description).toEqual("description edit")
    })
})