package logger

import "github.com/rs/zerolog/log"

func Info(msg string) {
  log.Info().Msg(msg)
}

func Error(msg string) {
  log.Error().Msg(msg)
}
