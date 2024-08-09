package routes

import (
	"project-user/commons/middlewares"
	"project-user/commons/utils"
	"project-user/features/logout/controllers"
	"project-user/features/logout/repositories"
	"project-user/features/logout/services"

	"github.com/labstack/echo/v4"
)

func LogoutRoute(e *echo.Echo, redisUtil utils.RedisUtil) {
	redisRepository := repositories.NewRedisRepository()
	logoutService := services.NewLogoutService(redisUtil, redisRepository)
	logoutController := controllers.NewLogoutController(logoutService)
	e.POST("/api/v1/users/logout", logoutController.Logout, middlewares.PrintRequestResponseLogWithNoRequestBody, middlewares.GetSessionIdUser)
}
