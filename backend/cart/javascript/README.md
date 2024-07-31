# Cart

## joi
npm i joi  
https://www.npmjs.com/package/joi  

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
add this on package.json:  
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
create file babel.config.json, add:
{
  "presets": ["@babel/preset-env"]
}

## supertest
npm install --save-dev supertest @types/supertest  
https://www.npmjs.com/package/supertest  

## mysql2
npm install --save mysql2  
https://sidorares.github.io/node-mysql2/docs  

## integration test  
npx jest tests/integration-tests/features/create-cart/services/create-cart-service.test.js --runInBand --detectOpenHandles  
npx jest tests/integration-tests/features/update-quantity/services/update-quantity-service.test.js --runInBand --detectOpenHandles  
npx jest tests/integration-tests/features/delete-cart/services/delete-cart-service.test.js --runInBand --detectOpenHandles  
npx jest tests/integration-tests/features/get-cart/services/get-cart-service.test.js --runInBand --detectOpenHandles  