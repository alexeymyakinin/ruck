package app

import (
	"io"
	"time"

	"app/src/core/cfg"
	"app/src/web/http/handler/chat"
	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewApplication() *gin.Engine {
	gin.SetMode(cfg.Config.AppMode)
	srv := gin.New()
	srv.Use(
		requestid.New(),
		gin.Recovery(),
		logger.SetLogger(
			logger.WithLogger(func(context *gin.Context, writer io.Writer, duration time.Duration) zerolog.Logger {
				return zerolog.New(writer).With().
					Str("id", requestid.Get(context)).
					Str("path", context.Request.URL.Path).
					Dur("duration", duration).
					Logger()
			})),
	)
	err := srv.SetTrustedProxies(cfg.Config.AppCORS)

	if err != nil {
		log.Fatal().Err(err)
	}

	{
		srv.GET("/api/v1/chat/:id", chat.GetChat)
		srv.GET("/api/v1/chat/:id/messages", chat.GetChatMessages)
	}

	return srv
}
