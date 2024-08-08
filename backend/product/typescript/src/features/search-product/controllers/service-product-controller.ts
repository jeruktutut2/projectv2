import { NextFunction, Request, Response } from "express";
import { SearchProductService } from "../services/search-product-service";
import { SearchProductRequest } from "../models/search-product-request";
import { setResponse } from "../../../helpers/response-message";
import { errorHandlerResponse } from "../../../exceptions/error-exception";

export class SearchProductController {
    static async searchProduct(req: Request, res: Response, next: NextFunction) {
        const requestId = req.get("X-REQUEST-ID") ?? ""
        try {
            const searchProductRequest: SearchProductRequest = req.body as SearchProductRequest
            const searchProductResponse = await SearchProductService.searchProduct(requestId, searchProductRequest.keyword)
            const response = setResponse(searchProductResponse, null)
            return res.status(200).json(response)
        } catch(e: unknown) {
            return errorHandlerResponse(res, e, requestId)
        }
    }
}