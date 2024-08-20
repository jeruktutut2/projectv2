package routes

import (
	"golang-postgres/commons/helpers"
	"golang-postgres/commons/middlewares"
	"golang-postgres/commons/utils"
	"golang-postgres/features/register/controllers"
	"golang-postgres/features/register/repositories"
	"golang-postgres/features/register/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(e *echo.Echo, postgresUtil utils.PostgresUtil, validate *validator.Validate, bcryptHelper helpers.BcryptHelper) {
	userRepository := repositories.NewUserRepository()
	registerService := services.NewRegisterService(postgresUtil, validate, bcryptHelper, userRepository)
	registerController := controllers.NewRegisterController(registerService)
	e.POST("/api/v1/users/register", registerController.Register, middlewares.PrintRequestResponseLog)
}
