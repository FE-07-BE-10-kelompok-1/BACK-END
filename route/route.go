package route

import (
	"bookstore/config"
	"bookstore/domain"

	"bookstore/feature/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteUser(e *echo.Echo, ud domain.UserHandler) {
	e.POST("/users", ud.InsertUser())
	e.POST("/login", ud.LoginHandler())
	e.PUT("/users", ud.UpdateUser(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.GET("/users", ud.GetProfile(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.DELETE("/users", ud.DeleteUser(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}
