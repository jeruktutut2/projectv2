package routes

import (
	"project-user/features/register/controllers"
	"project-user/features/register/repositories"
	"project-user/features/register/services"
	"project-user/helpers"
	"project-user/middlewares"
	"project-user/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(e *echo.Echo, mysqlUtil utils.MysqlUtil, validate *validator.Validate, bcryptHelper helpers.BcryptHelper) {
	userRepository := repositories.NewUserRepository()
	registerService := services.NewRegisterService(mysqlUtil, validate, bcryptHelper, userRepository)
	registerController := controllers.NewRegisterController(registerService)
	// e.POST("/api/v1/users/register", registerController.Register, middlewares.GetRequestId, middlewares.GetSessionIdUser)
	e.POST("/api/v1/users/register", registerController.Register, middlewares.PrintRequestResponseLog)
}
