# USER  

## application  
this service is to provide user management including login, register and logout  

## stack  
- go  
- echo: http framework  
- mysql: store user data  
- redis: store user session  

## run application  
go run main.go  

## Run in docker  
## crate network  
docker network create project  
docker network ls  
if you want to use bridge network (default), you no need to create network  

## create mysql with network  
docker run --name project-mysql-network --network project -e MYSQL_ROOT_PASSWORD=12345 -p 3310:3306 -d mysql:8.1.0  
if you want to use bridge network (default), you no need to put --network project  

## create redis with network
docker run --name project-redis-network --network project -p 6380:6379 -d redis:latest
if you want to use bridge network (default), you no need to put --network project


## change host mysql  
MYSQL_HOST=project-mysql-network  
DATABASE_URL="project-mysql-network://username:password@host:3306/databasename"  
REDIS_HOST=project-redis-network  

## build image  
docker build -t project-backend-user:1.0.0 .  

## run project  
docker run --name project-user-container --network project -p 10001:10001 project-user:1.0.0
docker run --name project-user-container -p 10001:10001 project-user:1.0.0  
if you want to use bridge network (default), you no need to put --network project  

## run test
go test -v tests/integration_tests/features/register/services/register_service_test.go  
go test -v tests/integration_tests/features/login/services/login_service_test.go  
go test -v tests/integration_tests/features/logout/services/logout_service_test.go  

go test -v tests/unit_tests/features/register/services/register_service_test.go  
go test -v tests/unit_tests/features/login/services/login_service_test.go  
go test -v tests/unit_tests/features/logout/services/logout_service_test.go  

go test -v tests/api_tests/features/register/register_test.go  
go test -v tests/api_tests/features/login/login_test.go  
go test -v tests/api_tests/features/logout/logout_test.go   

## test grpc golang
https://github.com/grpc/grpc-go/issues/1786  

## run curl test
chmod +x register_curl.sh
./register_curl.sh

chmod +x login_curl.sh
./login_curl.sh

chmod +x logout_curl.sh
./logout_curl.sh