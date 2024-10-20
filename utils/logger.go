package utils

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitServiceAdvancedLogger(loggerTag string) *zerolog.Logger {
	logger := log.With().Str("TAG", loggerTag).Logger()

	return &logger
}
