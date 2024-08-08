import { Client } from "@elastic/elasticsearch"
import { ElasticsearchUtil } from "../../../commons/utils/elasticsearch-util"
import { SearchProductResponse } from "../models/search-product-response"
import { ResponseException } from "../../../commons/exceptions/response-exception"
import { errorHandler } from "../../../commons/exceptions/error-exception"
import { setErrorMessages } from "../../../commons/helpers/error-message"

export class SearchProductService {
    static async searchProduct(requestId: string, keyword: string): Promise<SearchProductResponse[]> {
        let client: Client | null = null
        try {
            client = ElasticsearchUtil.getClient()
            const result = await client.search({
                index: "products",
                query: {
                    bool: {
                        should: [
                            { match: { name: keyword } },
                            { match: { description: keyword } }
                        ]
                    }
                }
            })
            let searchResponses: SearchProductResponse[] = []
            const hits = result.hits.hits
            if (hits.length < 1) {
                const errorMessage = "cannot find data with keyword: " + keyword
                throw new ResponseException(404, setErrorMessages(errorMessage), errorMessage)
            }
            hits.forEach((hit: any) => {
                const searchResponse: SearchProductResponse = {
                    id: hit._source.id,
                    userId: hit._source.userId,
                    name: hit._source.name,
                    description: hit._source.description,
                    stock: hit._source.stock
                }
                searchResponses.push(searchResponse)
            })
            return Promise.resolve(searchResponses)
        } catch(e: unknown) {
            errorHandler(e, requestId)
            return Promise.reject(e)
        } 
    }
}