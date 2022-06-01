package app

import (
	"app/src/api/http/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewApplication() *echo.Echo {
	e := echo.New()
	e.Use(
		middleware.RequestID(),
		middleware.Recover(),
		middleware.Logger(),
	)

	{
		e.GET("/api/v1/chat/:id", handler.GetChat)
	}
	{
		e.POST("/api/v1/user", handler.CreateUser)
		e.GET("/api/v1/user/:id", handler.GetUserByID)
	}

	return e
}
