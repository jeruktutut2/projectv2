import { PoolConnection } from "mysql2/promise";
import { MysqlUtil } from "../../../commons/utils/mysql-utils";
import { errorHandler } from "../../../commons/exceptions/error-exception";
import { DeleteProductByIdRepository } from "../repositories/delete-product-by-id-repository";
import { ResponseException } from "../../../commons/exceptions/response-exception";
import { ElasticsearchUtil } from "../../../commons/utils/elasticsearch-util";
import { setErrorMessages } from "../../../commons/helpers/error-message";
import { DataMessage, setDataMessage } from "../../../commons/helpers/data-message";

export class DeleteProductByIdService {
    static async deleteProductById(requestId: string, id: number): Promise<DataMessage> {
        let poolConnection: PoolConnection | null = null
        try {
            poolConnection = await MysqlUtil.getPool().getConnection()
            await poolConnection.beginTransaction()

            const [resultSetHeader] = await DeleteProductByIdRepository.DeleteProductById(poolConnection, id)
            if (resultSetHeader.affectedRows !== 1) {
                const errorMessage = "number of affected rows when deleting product is not one:" + resultSetHeader.affectedRows.toString()
                throw new ResponseException(500, setErrorMessages(errorMessage), errorMessage)
            }

            await ElasticsearchUtil.getClient().delete({
                index: "products",
                id: id.toString()
            })

            await poolConnection.commit()

            const dataMessage = setDataMessage("successfully delete product")
            return Promise.resolve(dataMessage)
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