import { FieldPacket, ResultSetHeader, RowDataPacket } from "mysql2";
import { PoolConnection } from "mysql2/promise";

export class GetProductByIdRepository {
    static async getProductById(poolConnection: PoolConnection, id: number): Promise<[RowDataPacket[], FieldPacket[]]> {
        // console.log("masuk ke sini");
        
        const query = `SELECT id, user_id, name, description, stock FROM products WHERE id = ?`
        const result = await poolConnection.execute<RowDataPacket[]>(query, [id])
        // console.log("result:", result);
        return result
    }
}