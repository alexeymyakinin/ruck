package app

import (
	"github.com/alexeymyakinin/ruck/app/src/api/http/handler"
	"github.com/alexeymyakinin/ruck/app/src/core/env"
	"github.com/alexeymyakinin/ruck/app/src/core/service"
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

	e.POST("/api/v1/auth", handler.SignIn)
	e.POST("/api/v1/user", handler.CreateUser)

	r := e.Group("", middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &service.JWTClaims{},
		SigningKey: []byte(env.JWTSecret),
	}))

	{
		r.POST("/api/v1/chat", handler.CreateChat)
		r.GET("/api/v1/chat/:id", handler.GetChat)
	}
	{
		r.GET("/api/v1/user/:id", handler.GetUser)
	}

	return e
}
