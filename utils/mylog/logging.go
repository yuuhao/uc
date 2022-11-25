package mylog

import (
	"io"
	"os"
	"path"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var ZeroLogger *zeroLogging

var SubZeroLogger *zeroLogging

type Config struct {
	ConsoleLoggingEnabled bool `json:"console_logging_enabled"`
	// EncodeLogsAsJson makes the log framework log JSON
	EncodeLogsAsJson bool
	// FileLoggingEnabled makes the framework log to a file
	// the fields below can be skipped if this value is false!
	FileLoggingEnabled bool
	// Directory to log to to when filelogging is enabled
	Directory string
	// Filename is the name of the logfile which will be placed inside the directory
	Filename string
	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int
	// MaxBackups the max number of rolled files to keep
	MaxBackups int
	// MaxAge the max age in days to keep a logfile
	MaxAge int
}

type zeroLogging struct {
	*zerolog.Logger
}

func InitZeroLogger(config Config) {
	var writers []io.Writer

	if config.ConsoleLoggingEnabled {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}
	if config.FileLoggingEnabled {
		writers = append(writers, newRollingFile(config))
	}

	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	mw := io.MultiWriter(writers...)

	logger := zerolog.New(mw).With().Timestamp().Caller()
	slugger := logger.Logger()
	ZeroLogger = &zeroLogging{&slugger}
	subLogger := logger.Str("aa", "sub").Logger()
	SubZeroLogger = &zeroLogging{&subLogger}
}

func newRollingFile(config Config) io.Writer {

	if err := os.MkdirAll(config.Directory, 0744); err != nil {
		log.Error().Err(err).Str("path", config.Directory).Msg("cnt create dir")
	}

	return &lumberjack.Logger{
		Filename:   path.Join(config.Directory, config.Filename),
		MaxBackups: config.MaxBackups,
		MaxSize:    config.MaxSize,
		MaxAge:     config.MaxAge,
	}
}
