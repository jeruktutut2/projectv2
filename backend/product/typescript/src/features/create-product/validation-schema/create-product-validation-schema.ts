import { ZodType, z } from "zod";

export class CreateProductValidationSchema {
    static readonly CREATE: ZodType = z.object({
        userId: z.number({required_error: 'userId is required'}).positive(),
        name: z.string({required_error: 'name is required'}).min(1).max(255),
        description: z.string({required_error: 'description is required'}).min(1).max(1000),
        stock: z.number({required_error: 'stock is required'}).min(0)
    })
}