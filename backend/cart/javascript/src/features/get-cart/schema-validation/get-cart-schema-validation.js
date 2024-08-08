import Joi from "joi";

export const getCartValidation = Joi.object({
    userId: Joi.number().min(1).positive().required()
})