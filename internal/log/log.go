package log

import (
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func Info(msg string) {
	log.Info().Caller(1).Msg(msg)
}

func Infof(format string, v ...interface{}) {
	log.Info().Caller(1).Msgf(format, v...)
}

func Debug(msg string) {
	log.Debug().Caller(1).Msg(msg)
}

func Debugf(format string, v ...interface{}) {
	log.Debug().Caller(1).Msgf(format, v...)
}

func Error(err error) {
	logger := log.Error().Caller(1).Err(err)
	if _, ok := err.(stackTracer); ok {
		logger.Msgf("%+v\n", err)
		return
	}
	logger.Send()
}

func Fatal(err error) {
	logger := log.Fatal().Caller(1).Err(err)
	if _, ok := err.(stackTracer); ok {
		logger.Msgf("%+v\n", err)
		return
	}
	logger.Send()
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}
