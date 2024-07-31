package routes

import (
	"project-user/features/logout/controllers"
	"project-user/features/logout/repositories"
	"project-user/features/logout/services"
	"project-user/middlewares"
	"project-user/utils"

	"github.com/labstack/echo/v4"
)

func LogoutRoute(e *echo.Echo, redisUtil utils.RedisUtil) {
	redisRepository := repositories.NewRedisRepository()
	logoutService := services.NewLogoutService(redisUtil, redisRepository)
	logoutController := controllers.NewLogoutController(logoutService)
	// e.POST("/api/v1/users/logout", logoutController.Logout, middlewares.GetRequestId, middlewares.GetSessionIdUser)
	e.POST("/api/v1/users/logout", logoutController.Logout, middlewares.PrintRequestResponseLogWithNoRequestBody, middlewares.GetSessionIdUser)
}
