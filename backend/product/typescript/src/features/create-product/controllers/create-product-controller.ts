import { Request, Response, NextFunction } from "express";
import { CreateProductService } from "../services/create-product-service";
import { CreateProductRequest } from "../models/create-product-request";
import { setResponse } from "../../../helpers/response-message";
import { errorHandlerResponse } from "../../../exceptions/error-exception";

export class CreateProductController {
    static async createProduct(req: Request, res: Response, next: NextFunction) {
        const requestId = req.get("X-REQUEST-ID") ?? ""
        try {
            const createProductRequest: CreateProductRequest = req.body as CreateProductRequest
            const createProductResponse = await CreateProductService.create(requestId, createProductRequest)
            const response = setResponse(createProductResponse, null)
            return res.status(201).json(response)
        } catch(e: unknown) {
            return errorHandlerResponse(res, e, requestId)
        }
    }
}