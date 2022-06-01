package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Logger(name string) zerolog.Logger {
	return log.Logger.With().Str("logger", name).Logger()
}
