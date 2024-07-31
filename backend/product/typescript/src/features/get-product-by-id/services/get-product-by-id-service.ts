import { errorHandler } from "../../../exceptions/error-exception";
import { MysqlUtil } from "../../../utils/mysql-utils";
import { GetProductByIdResponse } from "../models/get-product-by-id-response";
import { PoolConnection } from "mysql2/promise";
import { GetProductByIdRepository } from "../repositories/get-product-by-id-repository";
import { ResponseException } from "../../../exceptions/response-exception";
import { setErrorMessages } from "../../../helpers/error-message";
// import { setErrorMessages } from "../../../exceptions/exception";

export class GetProductByIdService {
    static async getProductById(requestId: string, id: number): Promise<GetProductByIdResponse> {
        let poolConnection: PoolConnection | null = null
        try {
            poolConnection = await MysqlUtil.getPool().getConnection()
            await poolConnection.beginTransaction()
            // console.log("lewat sini 1");
            const [rows] = await GetProductByIdRepository.getProductById(poolConnection, id)
            if (rows.length !== 1) {
                const errorMessage = "cannot find product with id:" + id.toString()
                throw new ResponseException(400, setErrorMessages(errorMessage), errorMessage)
            }
            // console.log("lewat sini 2");
            // console.log("rows:", rows);
            
            await poolConnection.commit()
            const response = {
                id: rows[0].id,
                name: rows[0].name,
                description: rows[0].description,
                stoct: rows[0].stock
            }
            // console.log("response:", response);
            
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
    }
}