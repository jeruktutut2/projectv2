import { Client } from "@elastic/elasticsearch"
// import { GetResponse } from "@elastic/elasticsearch/lib/api/types"
import { Pool } from "mysql2"
import { FieldPacket, PoolConnection, ResultSetHeader, RowDataPacket } from "mysql2/promise"

export async function createTableProducts(poolConnection: PoolConnection) {
    try {
        const query = `CREATE TABLE products (
  	        id int NOT NULL AUTO_INCREMENT,
	        user_id int NOT NULL,
	        name varchar(255) NOT NULL,
	        description TEXT NOT NULL,
	        stock MEDIUMINT NOT NULL,
  	        PRIMARY KEY (id),
  	        KEY fk_user_permission_user (user_id),
  	        CONSTRAINT user_permission_ibfk_1 FOREIGN KEY (user_id) REFERENCES user (id)
        ) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`
        const [result] = await poolConnection.execute<ResultSetHeader>(query)
    } catch(e) {
        console.log("error when creating table products:", e);
    }
    // if (result.affectedRows !== 1) {
    //     console.log("error when creating table products");
    // }
    console.log("create table products succedded");
}

export async function createDataProducts(poolConnection: PoolConnection) {
    try {
        const query = `INSERT INTO products (id, user_id, name, description, stock) VALUES (1, 1, "name1", "description1", 1), (2, 1, "name2", "description2", 1), (3, 1, "name3", "description3", 1);`
        const [result] = await poolConnection.execute<ResultSetHeader>(query)
    } catch(e) {
        console.log("error when creating data products");
    }
    console.log("create data products succedded");
    
}

export async function getDataProduct(poolConnection: PoolConnection, id: number): Promise<[RowDataPacket[], FieldPacket[]]> {
    try {
        const query = `SELECT id, user_id, name, description, stock FROM products WHERE id = ?;`
        const result = await poolConnection.execute<RowDataPacket[]>(query, [id])
        // console.log("result:", result);
        
        console.log("get data product succedded");
        return result
    } catch(e) {
        console.log("error when get data product");
        return Promise.reject(e)
    }
    // console.log("get data product succedded");
}

export async function deleteTableProducts (poolConnection: PoolConnection) {
    try {
        const query = `DROP TABLE IF EXISTS products;`
        const [result] = await poolConnection.execute<ResultSetHeader>(query)
    } catch(e) {
        console.log("error when deleting table products:", e);
    }
    // if (result.affectedRows !== 1) {
    //     console.log("error when deleting table products");
    // }

    console.log("delete table products succedded");
    
}

export async function createDataProductsElasticsearch(client: Client) {
    try {
        await client.index({
            index: "products",
            id: "1",
            document: {
                id: "1",
                userId: "1",
                name: "name1",
                description: "description1",
                stock: 1
            }
        })
        await client.index({
            index: "products",
            id: "2",
            document: {
                id: "2",
                userId: "1",
                name: "name2",
                description: "description2",
                stock: 1
            }
        })
        await client.index({
            index: "products",
            id: "3",
            document: {
                id: "3",
                userId: "1",
                name: "name3",
                description: "description3",
                stock: 1
            }
        })
    } catch(e: unknown) {
        console.log("error when creating data product elasticsearch:", e);
        
    }
}

interface Product {
    id: string,
    userId: string,
    name: string,
    description: string
}

// why don't use GetResponse, GetResponse doesn't exist in elasticsearch library
export async function getDataProductsElasticsearch(client: Client): Promise<any>  {
    try {
        const result = await client.get({
            index: "products",
            id: "1"
        })
        return result
    } catch(e: unknown) {
        console.log("error when getting data product elasticsearch:", e);
    }
}