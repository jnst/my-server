package logger

import (
	"os"

	"github.com/rs/zerolog"
)

func New() *zerolog.Logger {
	logger := zerolog.New(os.Stderr).
		With().
		Timestamp().
		Logger()
	return &logger
}
