package routes

import (
	"project-user/features/login/controllers"
	"project-user/features/login/repositories"
	"project-user/features/login/services"
	"project-user/helpers"
	"project-user/middlewares"
	"project-user/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func LoginRoute(e *echo.Echo, mysqlUtil utils.MysqlUtil, redisUtil utils.RedisUtil, validate *validator.Validate, bcryptHelper helpers.BcryptHelper, uuidHelper helpers.UuidHelper) {
	userRepository := repositories.NewUserRepository()
	userPermissionRepository := repositories.NewUserPermissinoRepository()
	redisRepository := repositories.NewRedisRepository()
	loginService := services.NewLoginService(mysqlUtil, redisUtil, validate, userRepository, userPermissionRepository, bcryptHelper, redisRepository, uuidHelper)
	loginController := controllers.NewLoginController(loginService)
	e.POST("/api/v1/users/login", loginController.Login, middlewares.PrintRequestResponseLog, middlewares.GetSessionIdUser)
}
