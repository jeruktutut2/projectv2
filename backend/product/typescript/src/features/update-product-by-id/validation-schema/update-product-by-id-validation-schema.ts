import { ZodType, z } from "zod";

export class UpdateProductByIdValidationSchema {
    static readonly UPDATE: ZodType = z.object({
        id: z.number({required_error: 'id is required'}).positive(),
        name: z.string({required_error: 'name is required'}).min(1).max(255),
        description: z.string({required_error: 'description is required'}).min(1).max(1000)
    })
}