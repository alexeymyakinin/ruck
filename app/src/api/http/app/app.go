package app

import (
	"app/src/api/http/handler"
	"app/src/core/cfg"
	"app/src/core/log"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func NewApplication() *gin.Engine {
	gin.SetMode(cfg.Config.AppMode)

	srv := gin.New()
	srv.Use(
		requestid.New(),
		gin.Recovery(),
		log.AccessLogger(),
	)
	if err := srv.SetTrustedProxies(cfg.Config.AppCORS); err != nil {
		panic(err)
	}

	{
		srv.GET("/api/v1/chat/:id", handler.GetChat)
	}
	{
		srv.POST("/api/v1/user", handler.CreateUser)
		srv.GET("/api/v1/user/:id", handler.GetUserByID)
	}

	return srv
}
