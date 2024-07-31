import { Client } from "@elastic/elasticsearch"
import { ElasticsearchUtil } from "../../../utils/elasticsearch-util"
import { SearchProductResponse } from "../models/search-product-response"
import { ResponseException } from "../../../exceptions/response-exception"
// import { setErrorMessages } from "../../../exceptions/exception"
import { errorHandler } from "../../../exceptions/error-exception"
import { setErrorMessages } from "../../../helpers/error-message"

export class SearchProductService {
    static async searchProduct(requestId: string, keyword: string): Promise<SearchProductResponse[]> {
        let client: Client | null = null
        try {
            client = ElasticsearchUtil.getClient()
            // const result = await client.msearch({
            //     searches: [
            //         { 
            //             index: "products" 
            //         },
            //         { 
            //             query: { 
            //                 match: { 
            //                     name: keyword 
            //                 } 
            //             } 
            //         },

            //         { 
            //             index: "products" 
            //         },
            //         { 
            //             query: { 
            //                 match: { 
            //                     description: keyword 
            //                 } 
            //             } 
            //         }
            //     ]
            // })
            // console.log("result:", result);
            // console.log("result.responses[0]:", result.responses[0]);
            // console.log("result.responses[1]:", result.responses[1]);
            // console.log("result.took:", result.took);
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
            // console.log("result:", result);
            let searchResponses: SearchProductResponse[] = []
            const hits = result.hits.hits
            if (hits.length < 1) {
                const errorMessage = "cannot find data with keyword: " + keyword
                throw new ResponseException(404, setErrorMessages(errorMessage), errorMessage)
            }
            // if (hits.length > 0) {
            //     // let searchResponses: SearchResponse[] = []
            //     hits.forEach((hit: any) => {
            //         // console.log("hit:", hit._source);
            //         const searchResponse: SearchResponse = {
            //             id: hit._source.id,
            //             userId: hit._source.userId,
            //             name: hit._source.name,
            //             description: hit._source.description,
            //             stock: hit._source.stock
            //         }
            //         searchResponses.push(searchResponse)
            //         // hit._source
            //         // if () {
                    
            //         // }
            //     })
            // }
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
        // finally {
        //     if (client) {
                
        //     }
        // }
    }
}