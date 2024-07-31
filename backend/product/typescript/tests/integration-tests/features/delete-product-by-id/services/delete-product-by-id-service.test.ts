import { PoolConnection } from "mysql2/promise"
import { ElasticsearchUtil } from "../../../../../src/utils/elasticsearch-util"
import { MysqlUtil } from "../../../../../src/utils/mysql-utils"
import { createDataProducts, createDataProductsElasticsearch, createTableProducts, deleteTableProducts, getDataProduct, getDataProductsElasticsearch } from "../../../../initialize/products"
import { DeleteProductByIdService } from "../../../../../src/features/delete-product-by-id/services/delete-product-by-id-service";
import { Client } from "@elastic/elasticsearch";
import { setDataMessage } from "../../../../../src/helpers/data-message";

describe("delete product by id", () => {

    const requestId = "requestId"
    const id = 1

    beforeAll( async () => {
        await MysqlUtil.getInstance()
        await ElasticsearchUtil.getInstance()
    })

    beforeEach(() => {

    })

    afterEach(() => {
        jest.resetAllMocks();
    })

    afterAll( async () => {
        await MysqlUtil.close()
        await ElasticsearchUtil.closeClient()
    })

    it("should return internal server error", async () => {
        const poolConnection: PoolConnection = await MysqlUtil.getPool().getConnection()
        await deleteTableProducts(poolConnection)
        await expect(async () => await DeleteProductByIdService.deleteProductById(requestId, id)).rejects.toThrow("internal server error")
    })

    it("should return number affected rows is not one", async () => {
        const poolConnection: PoolConnection = await MysqlUtil.getPool().getConnection()
        await deleteTableProducts(poolConnection)
        await createTableProducts(poolConnection)
        await expect(async () => await DeleteProductByIdService.deleteProductById(requestId, id)).rejects.toThrow("number of affected rows when deleting product is not one:0")
    })

    it("should success", async () => {
        const poolConnection: PoolConnection = await MysqlUtil.getPool().getConnection()
        await deleteTableProducts(poolConnection)
        await createTableProducts(poolConnection)
        await createDataProducts(poolConnection)
        const client: Client = ElasticsearchUtil.getClient()
        await createDataProductsElasticsearch(client)

        const result = await DeleteProductByIdService.deleteProductById(requestId, id)

        const [rows] = await getDataProduct(poolConnection, 1)
        expect(rows.length).toEqual(0)

        // const dataProductElasticsearch  = await getDataProductsElasticsearch(client)
        // console.log("dataProductElasticsearch:", dataProductElasticsearch);
        // await expect(async () => await getDataProductsElasticsearch(client)).rejects.toThrow(`{"_index":"products_v1","_type":"_doc","_id":"1","found":false}`)
        await expect(getDataProductsElasticsearch(client)).resolves.toEqual(undefined)
        
        // expect(dataProductElasticsearch.found).toEqual(0)

        const dataMessage = setDataMessage("successfully delete product")
        expect(result).toEqual(dataMessage)

        // await expect(DeleteProductByIdService.DeleteProductById(requestId, id)).resolves.toEqual(true)
    })
})