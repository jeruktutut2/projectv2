import { NextFunction, Request, Response } from "express";
import { GetProductByIdService } from "../services/get-product-by-id-service";
import { GetProductByIdRequest } from "../models/get-product-by-id-request";
import { setResponse } from "../../../commons/helpers/response-message";
import { errorHandlerResponse } from "../../../commons/exceptions/error-exception";
export class GetProductByIdController {
    static async getProductById(req: Request, res: Response, next: NextFunction) {
        const requestId = req.get("X-REQUEST-ID") ?? ""
        try {
            const getProductByIdRequest: GetProductByIdRequest = req.body as GetProductByIdRequest
            const getProductByIdResponse = await GetProductByIdService.getProductById(requestId, Number(getProductByIdRequest.id))
            const response = setResponse(getProductByIdResponse, null)
            return res.status(200).json(response)
        } catch(e: unknown) {
            return errorHandlerResponse(res, e, requestId)
        }
    }
}