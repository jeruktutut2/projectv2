import { Product } from "../../../../../src/features/create-product/models/create-product";
import { CreateProductRepository } from "../../../../../src/features/create-product/repositories/create-product-repository";
import { MysqlUtil } from "../../../../../src/utils/mysql-utils";
import { PoolConnection } from 'mysql2/promise';
import { CreateProductService } from "../../../../../src/features/create-product/services/create-product-service";
import { CreateProductRequest } from "../../../../../src/features/create-product/models/create-product-request";
import { createTableProducts, deleteTableProducts, getDataProduct } from "../../../../initialize/products";
import { ElasticsearchUtil } from "../../../../../src/utils/elasticsearch-util";

describe("create product", () => {

    const requestId = "requestId"
    let createProductRequest: CreateProductRequest = {
        userId: 0,
        name: "",
        description: "",
        stock: 0
    }

    beforeAll(async () => {
        await MysqlUtil.getInstance()
        await ElasticsearchUtil.getInstance()
    })

    beforeEach(() => {
        createProductRequest = {
            userId: 1,
            name: "name",
            description: "description",
            stock: 1
        }
    })

    afterEach(() => {
        jest.resetAllMocks();
    })

    afterAll(async () => {
        await MysqlUtil.close()
        await ElasticsearchUtil.closeClient()
    })

    it("should return error validation when giving an empty create product request", async () => {
        createProductRequest = {
            userId: 0,
            name: "",
            description: "",
            stock: 0
        }
        await expect(async () => await CreateProductService.create(requestId, createProductRequest)).rejects.toThrow("validation error");
    })

    it("should return internal server error when there is no table products", async () => {
        const poolConnection: PoolConnection = await MysqlUtil.getPool().getConnection()
        await deleteTableProducts(poolConnection)
        await expect(async () => await CreateProductService.create(requestId, createProductRequest)).rejects.toThrow("internal server error");
    })

    it("should success", async () => {
        const poolConnection: PoolConnection = await MysqlUtil.getPool().getConnection()
        await deleteTableProducts(poolConnection)
        await createTableProducts(poolConnection)
        const result = await CreateProductService.create(requestId, createProductRequest)

        const [rows] = await getDataProduct(poolConnection, 1)

        const resultElasticsearch = await ElasticsearchUtil.getClient().get({
            index: "products",
            id: "1"
        })
        expect(resultElasticsearch.found).toEqual(true)

        expect(rows.length).toEqual(1)
        expect(rows[0].name).toEqual("name1")
        expect(rows[0].description).toEqual("description1")
        expect(rows[0].stock).toEqual(1)

        expect(result.name).toEqual("name1")
        expect(result.description).toEqual("description1")
        expect(result.stock).toEqual(1)
    })
})