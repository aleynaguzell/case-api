package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

func Init() {
	log.Logger = logger
}

func Error(msg string, err error) {
	log.Logger.Error().Err(err).Msg(msg)
}

func CustomError(err error) {
	log.Logger.Error().Err(err)
}

func Info(msg string) {
	log.Logger.Info().Msg(msg)
}
func Fatal(msg string, err error) {
	log.Logger.Fatal().Err(err).Msg(msg)
}

func Warn(msg string) {
	log.Logger.Warn().Msg(msg)
}
