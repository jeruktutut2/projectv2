import { Client } from "@elastic/elasticsearch";
import { SearchProductService } from "../../../../../src/features/search-product/services/search-product-service";
import { ElasticsearchUtil } from "../../../../../src/utils/elasticsearch-util";
import { createDataProductsElasticsearch } from "../../../../initialize/products";

describe("", () => {

    const requestId = "requestId"
    
    beforeAll( async () => {
        await ElasticsearchUtil.getInstance()
    })

    beforeEach(() => {

    })

    afterEach(() => {
        jest.resetAllMocks();
    })

    afterAll( async () => {
        await ElasticsearchUtil.closeClient()
    })

    it("should return not found when data don't exists with certain keyword", async () => {
        const client: Client = ElasticsearchUtil.getClient()
        await createDataProductsElasticsearch(client)
        const keyword = "nam"
        await expect(async () => await SearchProductService.searchProduct(requestId, keyword)).rejects.toThrow("cannot find data with keyword: nam")
    })

    it("should success", async () => {
        const keyword = "name1"
        const client: Client = ElasticsearchUtil.getClient()
        await createDataProductsElasticsearch(client)
        const result = await SearchProductService.searchProduct(requestId, keyword)
        expect(result.length).toEqual(1)
        expect(result[0].id).toEqual("1")
        expect(result[0].name).toEqual("name1")
        expect(result[0].description).toEqual("description1")
        expect(result[0].stock).toEqual(1)
        expect(result[0].userId).toEqual("1")
        
    })
})