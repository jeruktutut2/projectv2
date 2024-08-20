package routes

import (
	"golang-postgres/commons/helpers"
	"golang-postgres/commons/middlewares"
	"golang-postgres/commons/utils"
	"golang-postgres/features/login/controllers"
	"golang-postgres/features/login/repositories"
	"golang-postgres/features/login/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func LoginRoute(e *echo.Echo, postgresUtil utils.PostgresUtil, redisUtil utils.RedisUtil, validate *validator.Validate, bcryptHelper helpers.BcryptHelper, uuidHelper helpers.UuidHelper) {
	userRepository := repositories.NewUserRepository()
	userPermissionRepository := repositories.NewUserPermissinoRepository()
	redisRepository := repositories.NewRedisRepository()
	loginService := services.NewLoginService(postgresUtil, redisUtil, validate, userRepository, userPermissionRepository, bcryptHelper, redisRepository, uuidHelper)
	loginController := controllers.NewLoginController(loginService)
	e.POST("/api/v1/users/login", loginController.Login, middlewares.PrintRequestResponseLog, middlewares.GetSessionIdUser)
}
