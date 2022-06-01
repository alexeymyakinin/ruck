package log

import (
	"io"
	"time"

	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Logger(name string) zerolog.Logger {
	return log.Logger.With().Str("logger", name).Logger()
}

func AccessLogger() gin.HandlerFunc {
	return logger.SetLogger(logger.WithLogger(
		func(c *gin.Context, w io.Writer, duration time.Duration) zerolog.Logger {
			return zerolog.New(w).With().
				Timestamp().
				Str("req_id", requestid.Get(c)).
				Str("req_ip", c.ClientIP()).
				Str("req_path", c.Request.URL.Path).
				Str("req_method", c.Request.Method).
				Dur("req_duration", duration).
				Str("req_user_agent", c.Request.UserAgent()).
				Int("res_status", c.Writer.Status()).
				Logger()
		}))
}
