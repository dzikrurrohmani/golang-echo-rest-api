package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func LoadRoutes(e *echo.Echo, handler *handler) {
	authMiddleware := GetAuthMiddleware(handler.restoUsecase)

	menuGroup := e.Group("/menu")
	menuGroup.GET("", handler.GetMenu)

	orderGroup := e.Group("/order")
	orderGroup.POST("", handler.Order,
		authMiddleware.CheckAuth)
	orderGroup.GET("/:orderID", handler.GetOrderInfo,
		authMiddleware.CheckAuth)

	userGroup := e.Group("/user")
	userGroup.POST("/register", handler.RegisterUser)
	userGroup.POST("/login", handler.Login)
}

func LoadMiddlewares(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// the safety comes from using https, since Origin header is controlled by the browser
		AllowOrigins: []string{"https://restoku.com"},
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogLevel: log.ERROR,
	}))
}
