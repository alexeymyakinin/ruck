package app

import (
	"app/src/api/http/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewApplication() *echo.Echo {
	e := echo.New()
	e.HidePort = true
	e.HideBanner = true
	e.Use(
		middleware.RequestID(),
		middleware.Recover(),
		middleware.Logger(),
	)

	{
		e.POST("/api/v1/chat", handler.CreateChat)
		e.GET("/api/v1/chat/:id", handler.GetChat)
	}
	{
		e.POST("/api/v1/user", handler.CreateUser)
		e.GET("/api/v1/user/:id", handler.GetUser)
	}

	return e
}
