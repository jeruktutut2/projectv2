import { Product } from "../../../../../src/features/create-product/models/create-product";
import { CreateProductRepository } from "../../../../../src/features/create-product/repositories/create-product-repository";
import { MysqlUtil } from "../../../../../src/utils/mysql-utils";
import { PoolConnection } from 'mysql2/promise';
import { CreateProductService } from "../../../../../src/features/create-product/services/create-product-service";
import { CreateProductRequest } from "../../../../../src/features/create-product/models/create-product-request";
// import {  } from "module";
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
        // const requestId = "requestId"
        await MysqlUtil.getInstance()
        await ElasticsearchUtil.getInstance()
        // createProductRequest = {
        //     userId: 1,
        //     name: "name",
        //     description: "description",
        //     stock: 1
        // }
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

        // console.log("MysqlUtil.getPool():", await MysqlUtil.getPool().getConnection());
        
        // let poolConnection: PoolConnection | null = null

        // try {
        //     poolConnection = await MysqlUtil.getPool().getConnection()
        //     // console.log("poolConnection1:", poolConnection);
            
        //     await poolConnection.beginTransaction()
        //     // console.log("poolConnection2:", poolConnection);

        //     // console.log("poolConnection:", poolConnection);
        

        //     const product: Product = {
        //         userId: 1,
        //         name: "name",
        //         description: "description",
        //         stock: 1
        //     }
        //     CreateProductRepository.create(poolConnection, product)
        //     await poolConnection.commit()
        // } catch(e) {
        //     console.log("e:", e);
        //     if (poolConnection) {
        //         await poolConnection.rollback()
        //     }
        //     // await poolConnection.rollback()
        // } finally {
        //     if (poolConnection) {
        //         poolConnection.release()
        //     }
        //     // poolConnection.release()
        // }
        
        // poolConnection.release()

        createProductRequest = {
            userId: 0,
            name: "",
            description: "",
            stock: 0
        }
        // const await CreateProductService.create(requestId, createProductRequest)
        // await expect(async () => await CreateProductService.create(requestId, createProductRequest)).rejects.toThrow("[{\"field\":\"userId\",\"message\":\"Number must be greater than 0\"},{\"field\":\"name\",\"message\":\"String must contain at least 1 character(s)\"},{\"field\":\"description\",\"message\":\"String must contain at least 1 character(s)\"}]");
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
        // await expect(async () => await CreateProductService.create(requestId, createProductRequest)).rejects.toThrow("[{\"field\":\"message\",\"message\":\"internal server error\"}]");
        // await expect(CreateProductService.create(requestId, createProductRequest)).resolves.toEqual({"description": "description", "name": "name", "stock": 1})
        const result = await CreateProductService.create(requestId, createProductRequest)
        // console.log("result:", result);
        // console.log("result.name:", result.name);
        // console.log("result.description:", result.description);
        // console.log("result.stock:", result.stock);
        
        const [rows] = await getDataProduct(poolConnection, 1)
        // console.log("rows:", rows);

        const resultElasticsearch = await ElasticsearchUtil.getClient().get({
            index: "products",
            id: "1"
        })
        // console.log("a:", a);
        expect(resultElasticsearch.found).toEqual(true)
        // console.log("source:", resultElasticsearch._source);
        

        expect(rows.length).toEqual(1)
        expect(rows[0].name).toEqual("name1")
        expect(rows[0].description).toEqual("description1")
        expect(rows[0].stock).toEqual(1)
        // console.log("rows.length:", rows.length);
        
        // await expect(result).resolves.toEqual({"description": "description", "name": "name", "stock": 1})
        expect(result.name).toEqual("name1")
        expect(result.description).toEqual("description1")
        expect(result.stock).toEqual(1)
        

        // await expect(CreateProductService.create(requestId, createProductRequest)).resolves.toEqual({"description": "description", "name": "name", "stock": 1})
        // const [rows] = await getDataProduct(poolConnection, 1)
        // expect(rows.length).toEqual(1)
    })
})