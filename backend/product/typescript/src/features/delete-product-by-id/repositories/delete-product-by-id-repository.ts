import { FieldPacket, PoolConnection, ResultSetHeader } from "mysql2/promise";

export class DeleteProductByIdRepository {
    static async DeleteProductById(poolConnection: PoolConnection, id: number): Promise<[ResultSetHeader, FieldPacket[]]> {
        const query = `DELETE FROM products WHERE id = ?;`
        const result = await poolConnection.execute<ResultSetHeader>(query, [id])
        return result
    } 
}