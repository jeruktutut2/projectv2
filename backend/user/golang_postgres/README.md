# USER  

## application  
this service provide login, register, logout  

## stack  
- go  
- echo: http framework  
- postgres: store user data  
- redis: store user session  

## run application  
go run main.go  

## install postgres
go get github.com/jackc/pgx/v5  
go get github.com/jackc/pgx/v5/pgxpool  

## install redis
go get github.com/redis/go-redis/v9  

## install validator
go get github.com/go-playground/validator/v10  

## install testify
go get github.com/stretchr/testify  

## install uuid
go get github.com/google/uuid  

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

## install echo
go get github.com/labstack/echo/v4  

## curl test
chmod +x register_curl.sh
./register_curl.sh

chmod +x login_curl.sh
./login_curl.sh

chmod +x logout_curl.sh
./logout_curl.sh