import { errorHandler } from "../../../commons/exceptions/error-exception";
import { MysqlUtil } from "../../../commons/utils/mysql-utils";
import { GetProductByIdResponse } from "../models/get-product-by-id-response";
import { PoolConnection } from "mysql2/promise";
import { GetProductByIdRepository } from "../repositories/get-product-by-id-repository";
import { ResponseException } from "../../../commons/exceptions/response-exception";
import { setErrorMessages } from "../../../commons/helpers/error-message";

export class GetProductByIdService {
    static async getProductById(requestId: string, id: number): Promise<GetProductByIdResponse> {
        let poolConnection: PoolConnection | null = null
        try {
            poolConnection = await MysqlUtil.getPool().getConnection()
            await poolConnection.beginTransaction()
            const [rows] = await GetProductByIdRepository.getProductById(poolConnection, id)
            if (rows.length !== 1) {
                const errorMessage = "cannot find product with id:" + id.toString()
                throw new ResponseException(400, setErrorMessages(errorMessage), errorMessage)
            }

            await poolConnection.commit()
            const response = {
                id: rows[0].id,
                name: rows[0].name,
                description: rows[0].description,
                stoct: rows[0].stock
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