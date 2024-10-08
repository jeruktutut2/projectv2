import { FieldPacket, Pool, PoolConnection, ResultSetHeader } from "mysql2/promise";
import { Product } from "../models/create-product";

export class CreateProductRepository {
    static async create(pool: PoolConnection, product: Product): Promise<[ResultSetHeader, FieldPacket[]]> {
        const query = `INSERT INTO products (user_id, name, description, stock) VALUES (?, ?, ?, ?);`
        const result = await pool.execute<ResultSetHeader>(query, [product.userId, product.name, product.description, product.stock])
        return result
    }
}