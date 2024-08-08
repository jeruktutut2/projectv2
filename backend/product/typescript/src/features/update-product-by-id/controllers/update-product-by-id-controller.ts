import { NextFunction, Request, Response } from "express";
import { UpdateProductByIdService } from "../services/update-product-by-id-service";
import { UpdateProductByIdRequest } from "../models/update-product-by-id-request";
import { setResponse } from "../../../helpers/response-message";
import { errorHandlerResponse } from "../../../exceptions/error-exception";

export class UpdateProductByIdController {
    static async updateProductById(req: Request, res: Response, next: NextFunction) {
        const requestId = req.get("X-REQUEST-ID") ?? ""
        try {
            const updateProductByIdRequest: UpdateProductByIdRequest = req.body as UpdateProductByIdRequest
            const updateProductByIdResponse = await UpdateProductByIdService.updateProductById(requestId, updateProductByIdRequest)
            const response = setResponse(updateProductByIdResponse, null)
            return res.status(200).json(response)
        } catch(e: unknown) {
            return errorHandlerResponse(res, e, requestId)
        }
    }
}