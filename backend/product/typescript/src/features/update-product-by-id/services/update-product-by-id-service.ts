import { PoolConnection } from "mysql2/promise";
import { UpdateProductByIdRequest } from "../models/update-product-by-id-request";
import { UpdateProductByIdResponse } from "../models/update-product-by-id-response";
import { MysqlUtil } from "../../../utils/mysql-utils";
import { errorHandler } from "../../../exceptions/error-exception";
import { UpdateProductByIdRepository } from "../repositories/update-product-by-id-repository";
import { ResponseException } from "../../../exceptions/response-exception";
import { Product } from "../models/product";
import { ElasticsearchUtil } from "../../../utils/elasticsearch-util";
import { Validation } from "../../../validation/validation";
import { UpdateProductByIdValidationSchema } from "../validation-schema/update-product-by-id-validation-schema";
import { setErrorMessages } from "../../../helpers/error-message";

export class UpdateProductByIdService {

    static async updateProductById(requestId: string, updateProductByIdRequest: UpdateProductByIdRequest): Promise<UpdateProductByIdResponse> {

        let poolConnection: PoolConnection | null = null
        try {

            updateProductByIdRequest = Validation.validate(UpdateProductByIdValidationSchema.UPDATE, updateProductByIdRequest)

            poolConnection = await MysqlUtil.getPool().getConnection()
            await poolConnection.beginTransaction()

            const [rows] = await UpdateProductByIdRepository.getById(poolConnection, updateProductByIdRequest.id)
            if (rows.length !== 1) {
                const errorMessage = "cannot find product with id:" + updateProductByIdRequest.id.toString()
                throw new ResponseException(400, setErrorMessages(errorMessage), errorMessage)
            }

            const product: Product = {
                id: updateProductByIdRequest.id,
                name: updateProductByIdRequest.name,
                description: updateProductByIdRequest.description
            }
            const [resultSetHeader] = await UpdateProductByIdRepository.updateNameAndDescriptionById(poolConnection, product)
            if (resultSetHeader.affectedRows !== 1) {
                const errorMessage = "number of affected rows when creating product is not one:" + resultSetHeader.affectedRows.toString()
                throw new ResponseException(500, setErrorMessages(errorMessage), errorMessage)
            }

            await ElasticsearchUtil.getClient().update({
                index: "products",
                id: updateProductByIdRequest.id.toString(),
                doc: {
                    name: updateProductByIdRequest.name,
                    description: updateProductByIdRequest.description
                }
            })

            await poolConnection.commit()
            const response = {
                name: updateProductByIdRequest.name,
                description: updateProductByIdRequest.description
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