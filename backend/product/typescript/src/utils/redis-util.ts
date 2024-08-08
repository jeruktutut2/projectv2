import { Redis } from "ioredis";

export class RedisUtil {
    private static instance: RedisUtil
    private static client: Redis

    private constructor() {
        console.log(new Date().toISOString(), " redis: connecting to redis: " + process.env.PROJECT_USER_REDIS_HOST + ":" + process.env.PROJECT_USER_REDIS_PORT);
        RedisUtil.client = new Redis({
            host: process.env.PROJECT_USER_REDIS_HOST ?? "",
            port: Number(process.env.PROJECT_USER_REDIS_PORT),
            db: Number(process.env.PROJECT_USER_REDIS_DATABASE)
        })
        console.log(new Date().toISOString(), " redis: connecteed to redis: " + process.env.PROJECT_USER_REDIS_HOST + ":" + process.env.PROJECT_USER_REDIS_PORT);
    }

    public static async getInstance(): Promise<RedisUtil> {
        try {
            if (!RedisUtil) {
                RedisUtil.instance = new RedisUtil()
            } else {
                RedisUtil.instance
            }
            console.log(new Date().toISOString(), " redis: pinging to redis: " + process.env.PROJECT_USER_REDIS_HOST + ":" + process.env.PROJECT_USER_REDIS_PORT);
            await RedisUtil.client.ping()
            console.log(new Date().toISOString(), " redis: pinged to redis: " + process.env.PROJECT_USER_REDIS_HOST + ":" + process.env.PROJECT_USER_REDIS_PORT);
            return Promise.resolve(RedisUtil.instance)
        } catch(e: unknown) {
            console.log("error when creating connection to redis: " + process.env.PROJECT_USER_REDIS_HOST + ":" + process.env.PROJECT_USER_REDIS_PORT);
            return Promise.reject(e)
        }
    }

    public static getClient() {
        return this.client
    }

    public static async disconnect() {
        console.log(new Date().toISOString(), " redis: disconnecting to redis:" + process.env.PROJECT_USER_REDIS_HOST + ":" + process.env.PROJECT_USER_REDIS_PORT);
        this.client.disconnect()
        console.log(new Date().toISOString(), " redis: disconnected to redis:" + process.env.PROJECT_USER_REDIS_HOST + ":" + process.env.PROJECT_USER_REDIS_PORT);
    }
}