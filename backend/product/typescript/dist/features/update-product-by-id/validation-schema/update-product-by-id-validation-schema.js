"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.UpdateProductByIdValidationSchema = void 0;
const zod_1 = require("zod");
class UpdateProductByIdValidationSchema {
}
exports.UpdateProductByIdValidationSchema = UpdateProductByIdValidationSchema;
UpdateProductByIdValidationSchema.UPDATE = zod_1.z.object({
    id: zod_1.z.number({ required_error: 'id is required' }).positive(),
    name: zod_1.z.string({ required_error: 'name is required' }).min(1).max(255),
    description: zod_1.z.string({ required_error: 'description is required' }).min(1).max(1000)
});
