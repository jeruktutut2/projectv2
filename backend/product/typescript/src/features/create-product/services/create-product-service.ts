import { errorHandler } from "../../../exceptions/error-exception";
import { Validation } from "../../../validation/validation";
import { CreateProductRequest } from "../models/create-product-request";
import { CreateProductResponse } from "../models/create-product-response";
import { CreateProductValidationSchema } from "../validation-schema/create-product-validation-schema";
import { PoolConnection } from 'mysql2/promise';
import { MysqlUtil } from "../../../utils/mysql-utils";
import { Product } from "../models/create-product";
import { CreateProductRepository } from "../repositories/create-product-repository";
import { ResponseException } from "../../../exceptions/response-exception";
import { ElasticsearchUtil } from "../../../utils/elasticsearch-util";
import { ProductElasticSearch } from "../models/create-product-elasticsearch";
import { setErrorMessages } from "../../../helpers/error-message";
// import { setErrorMessages } from "../../../exceptions/exception";

export class CreateProductService {
    static async create(requestId: string, createProductRequest: CreateProductRequest): Promise<CreateProductResponse> {
        // console.log("mantap");
        let poolConnection: PoolConnection | null = null
        try {
            // console.log("mantap1");
            createProductRequest = Validation.validate(CreateProductValidationSchema.CREATE, createProductRequest)
            // console.log("mantap2");
            
            poolConnection = await MysqlUtil.getPool().getConnection()
            await poolConnection.beginTransaction()

            const product: Product = {
                userId: createProductRequest.userId,
                name: createProductRequest.name,
                description: createProductRequest.description,
                stock: createProductRequest.stock
            }
            const [resultSetHeader] = await CreateProductRepository.create(poolConnection, product)
            // console.log("mantap:", resultSetHeader);
            if (resultSetHeader.affectedRows !== 1) {
                // throw new Error("number of affected rows when creating product is not one:" + resultSetHeader.affectedRows)                
                const errorMessage = "number of affected rows when creating product is not one:" + resultSetHeader.affectedRows.toString()
                throw new ResponseException(500, setErrorMessages(errorMessage), errorMessage)
            }

            // product.id = resultSetHeader.insertId
            // const productElasticsearch: ProductElasticSearch = {
            //     id: resultSetHeader.insertId.toString(),
            //     userId: product.userId.toString(),
            //     name: product.name,
            //     description: product.description
            // }
            // const result = await ElasticsearchUtil.getClient().index({
            //     index: "products",
            //     id: resultSetHeader.insertId.toString(),
            //     // document: productElasticsearch
            //     document: {
            //         id: resultSetHeader.insertId.toString(),
            //         userId: product.userId.toString(),
            //         name: product.name,
            //         description: product.description
            //     }
            // })

            await ElasticsearchUtil.getClient().index({
                index: "products",
                id: resultSetHeader.insertId.toString(),
                // document: productElasticsearch
                document: {
                    id: resultSetHeader.insertId.toString(),
                    userId: product.userId.toString(),
                    name: product.name,
                    description: product.description
                }
            })
            // console.log("result elasticsearch:", result);
            
            await poolConnection.commit()

            const response = {
                name: product.name,
                description: product.description,
                stock: product.stock
            }
            return Promise.resolve(response)
        } catch(e: unknown) {
            if (poolConnection) {
                await poolConnection.rollback()
            }
            errorHandler(e, requestId)
            // if (poolConnection) {
            //     await poolConnection.rollback()
            // }
            return Promise.reject(e)
        } finally {
            if (poolConnection) {
                poolConnection.release()
            }
        }

        // const response = {
        //     name: "name",
        //     description: "description",
        //     stock: 1
        // }
        // return response
    }
}