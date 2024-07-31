import { FieldPacket, PoolConnection, ResultSetHeader, RowDataPacket } from "mysql2/promise";
import { UpdateProductByIdRequest } from "../models/update-product-by-id-request";
import { Product } from "../models/product";

export class UpdateProductByIdRepository {
    static async getById(poolConnection: PoolConnection, id: number): Promise<[RowDataPacket[], FieldPacket[]]> {
        const query = `SELECT id, user_id, name, description FROM products WHERE id = ?`
        const result = await poolConnection.execute<RowDataPacket[]>(query, [id])
        return result
    }

    static async updateNameAndDescriptionById(poolConnection: PoolConnection, product: Product): Promise<[ResultSetHeader, FieldPacket[]]> {
        const query = `UPDATE products SET name = ?, description = ? WHERE id = ?`
        const result = await poolConnection.execute<ResultSetHeader>(query, [product.name, product.description, product.id])
        return result
    }
}