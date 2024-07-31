import { NextFunction, Request, Response } from "express";
import { DeleteProductByIdService } from "../services/delete-product-by-id-service";
import { DeleteProductByIdRequest } from "../models/delete-product-by-id-request";
import { setResponse } from "../../../helpers/response-message";
import { errorHandlerResponse } from "../../../exceptions/error-exception";
export class DeleteProductByIdController {
    static async deleteProductById(req: Request, res: Response, next: NextFunction) {
        const requestId = req.get("X-REQUEST-ID") ?? ""
        try {
            // const requestId = req.get("X-REQUEST-ID") ?? ""
            const deleteProductByIdRequest: DeleteProductByIdRequest = req.body as DeleteProductByIdRequest
            const deleteProductByIdResponse = await DeleteProductByIdService.deleteProductById(requestId, deleteProductByIdRequest.id)
            const response = setResponse(deleteProductByIdResponse, null)
            return res.status(200).json(response)
        } catch(e: unknown) {
            return errorHandlerResponse(res, e, requestId)
        }
    }
}