import Joi from "joi";

export const createCartValidation = Joi.object({
    userId: Joi.number().min(1).positive().required(),
    productId: Joi.number().min(1).positive().required(),
    quantity: Joi.number().min(1).positive().required()
})