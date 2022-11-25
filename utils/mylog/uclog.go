package mylog

import (
	"github.com/rs/zerolog"
)

func Error() *zerolog.Event {
	return ZeroLogger.Error()
}

func SubError() *zerolog.Event {
	return SubZeroLogger.Error()
}

func Info() *zerolog.Event {
	return ZeroLogger.Info()
}
