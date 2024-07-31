import Joi from "joi";

export const deleteCartValidation = Joi.object({
    id: Joi.number().min(1).positive().required(),
    userId: Joi.number().min(1).positive().required(),
    productId: Joi.number().min(1).positive().required()
})