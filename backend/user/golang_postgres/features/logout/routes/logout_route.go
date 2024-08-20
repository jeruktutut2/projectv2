package routes

import (
	"golang-postgres/commons/middlewares"
	"golang-postgres/commons/utils"
	"golang-postgres/features/logout/controllers"
	"golang-postgres/features/logout/repositories"
	"golang-postgres/features/logout/services"

	"github.com/labstack/echo/v4"
)

func LogoutRoute(e *echo.Echo, redisUtil utils.RedisUtil) {
	redisRepository := repositories.NewRedisRepository()
	logoutService := services.NewLogoutService(redisUtil, redisRepository)
	logoutController := controllers.NewLogoutController(logoutService)
	e.POST("/api/v1/users/logout", logoutController.Logout, middlewares.PrintRequestResponseLogWithNoRequestBody, middlewares.GetSessionIdUser)
}
