import { errorHandler } from "../../../commons/exceptions/error-exception";
import { Validation } from "../../../commons/validations/validation";
import { CreateProductRequest } from "../models/create-product-request";
import { CreateProductResponse } from "../models/create-product-response";
import { CreateProductValidationSchema } from "../validation-schema/create-product-validation-schema";
import { PoolConnection } from 'mysql2/promise';
import { MysqlUtil } from "../../../commons/utils/mysql-utils";
import { Product } from "../models/create-product";
import { CreateProductRepository } from "../repositories/create-product-repository";
import { ResponseException } from "../../../commons/exceptions/response-exception";
import { ElasticsearchUtil } from "../../../commons/utils/elasticsearch-util";
import { ProductElasticSearch } from "../models/create-product-elasticsearch";
import { setErrorMessages } from "../../../commons/helpers/error-message";

export class CreateProductService {
    static async create(requestId: string, createProductRequest: CreateProductRequest): Promise<CreateProductResponse> {

        let poolConnection: PoolConnection | null = null
        try {
            createProductRequest = Validation.validate(CreateProductValidationSchema.CREATE, createProductRequest)

            poolConnection = await MysqlUtil.getPool().getConnection()
            await poolConnection.beginTransaction()

            const product: Product = {
                userId: createProductRequest.userId,
                name: createProductRequest.name,
                description: createProductRequest.description,
                stock: createProductRequest.stock
            }
            const [resultSetHeader] = await CreateProductRepository.create(poolConnection, product)
            if (resultSetHeader.affectedRows !== 1) {
                const errorMessage = "number of affected rows when creating product is not one:" + resultSetHeader.affectedRows.toString()
                throw new ResponseException(500, setErrorMessages(errorMessage), errorMessage)
            }

            await ElasticsearchUtil.getClient().index({
                index: "products",
                id: resultSetHeader.insertId.toString(),
                document: {
                    id: resultSetHeader.insertId.toString(),
                    userId: product.userId.toString(),
                    name: product.name,
                    description: product.description
                }
            })

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
            return Promise.reject(e)
        } finally {
            if (poolConnection) {
                poolConnection.release()
            }
        }
    }
}