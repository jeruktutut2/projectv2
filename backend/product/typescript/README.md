# Product

## init project
npm init  

## zod  
npm i zod  
https://www.npmjs.com/package/zod  

## express  
npm install express  
npm install --save-dev @types/express    
https://www.npmjs.com/package/express  

## jest  
npm install --save-dev jest @types/jest  
https://www.npmjs.com/package/jest  

## babel  
npm install --save-dev babel-jest @babel/preset-env  
https://babeljs.io/setup#installation  
```
{
  "scripts": {
    "test": "jest"
  },
  "jest": {
    "transform": {
      "^.+\\.[t|j]sx?$": "babel-jest"
    }
  }
}
```
create file: babel.config.json
```
{
  "presets": ["@babel/preset-env"]
}
```
npm install --save-dev @babel/preset-typescript  
npm install --save-dev @jest/globals  
https://jestjs.io/docs/getting-started#using-typescript  
add "@babel/preset-typescript" to babel.config.json  

## typescript  
npm install --save-dev typescript  
https://www.npmjs.com/package/typescript  
"main": "index.js",

## init typescript project  
npx tsc --init  
"target": "es2016"  
"module": "commonjs"  
"moduleResolution": "Node"  
"include": [
    "src/**/*"
]  
"outDir": "./dist"

## mysql2
npm install --save mysql2  
npm install --save-dev @types/node  

## test  
npx jest path/to/your/test-file.js --runInBand  
npx jest tests/integration-test/services/user-service.test.js --detectOpenHandles  
npx jest tests/integration-tests/features/create-product/services/create-product-service.test.ts --runInBand --detectOpenHandles  
npx jest -t "should return internal server error when there is no table products"   

npx jest tests/integration-tests/features/get-product-by-id/services/get-product-by-id-service.test.ts --runInBand --detectOpenHandles  

<!-- cannot do unit testing -->
<!-- npx jest tests/unit-tests/features/create-product/services/create-product-service.test.ts --runInBand --detectOpenHandles   -->
npx jest tests/integration-tests/features/update-product-by-id/services/update-product-by-id.test.ts  --runInBand --detectOpenHandles  
npx jest tests/integration-tests/features/delete-product-by-id/services/delete-product-by-id-service.test.ts --runInBand --detectOpenHandles  
npx jest tests/integration-tests/features/search-product/services/search-product-service.test.ts  --runInBand --detectOpenHandles  
npx jest tests/api-tests/features/create-product/create-product.test.ts --runInBand --detectOpenHandles  
npx jest tests/api-tests/features/delete-product/delete-product.test.ts --runInBand --detectOpenHandles  
npx jest tests/api-tests/features/get-product-by-id/get-product-by-id.test.ts --runInBand --detectOpenHandles  
npx jest tests/api-tests/features/search-product/search-product.test.ts --runInBand --detectOpenHandles  
npx jest tests/api-tests/features/update-product-by-id/update-product-by-id.test.ts --runInBand --detectOpenHandles  
https://stackoverflow.com/questions/62214949/testing-grpc-functions  

## elasticsearch  
docker pull elasticsearch  
docker pull elasticsearch:7.17.22  
docker run -d --name project-elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.17.22  
https://www.elastic.co/guide/en/elasticsearch/client/javascript-api/current/getting-started-js.html  
npm install @elastic/elasticsearch  

## redis
https://www.npmjs.com/package/ioredis  
npm i ioredis  

## supertest
npm install --save-dev supertest @types/supertest  
https://www.npmjs.com/package/supertest  

## run
npm run start  

## run curl test
chmod +x create-product-curl.sh  
./create-product-curl.sh  

chmod +x get-product-by-id-curl.sh
./get-product-by-id-curl.sh

chmod +x search-product-curl.sh
./search-product-curl.sh

chmod +x update-product-by-id.sh
./update-product-by-id.sh

chmod +x delete-product-by-id.sh
./delete-product-by-id.sh