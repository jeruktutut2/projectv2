import { FieldPacket, ResultSetHeader, RowDataPacket } from "mysql2";
import { PoolConnection } from "mysql2/promise";

export class GetProductByIdRepository {
    static async getProductById(poolConnection: PoolConnection, id: number): Promise<[RowDataPacket[], FieldPacket[]]> {
        const query = `SELECT id, user_id, name, description, stock FROM products WHERE id = ?`
        const result = await poolConnection.execute<RowDataPacket[]>(query, [id])
        return result
    }
}