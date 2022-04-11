package logs

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

var (
	logger    zerolog.Logger
	errLogger zerolog.Logger
)

func InitLogger(level zerolog.Level) {

	//zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.SetGlobalLevel(level)

	stdout := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	logger = zerolog.New(stdout).With().Timestamp().Caller().Logger()

	stderr := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	errLogger = logger.Output(stderr)

	// multi output need
	/*multi := zerolog.MultiLevelWriter(stdout, stderr)
	errLogger = logger.Output(multi)*/
}

//panic/fatal/error/warn/info/debug/trace

func Info() *zerolog.Event {

	return logger.Info()
}

func Warn() *zerolog.Event {

	return logger.Warn()
}

func Debug() *zerolog.Event {

	return logger.Debug()
}

func Trace() *zerolog.Event {

	return logger.Trace()
}

func Error(err error) *zerolog.Event {

	evt := logger.Trace()
	if err != nil {
		evt = errLogger.Error().Err(err)
	}
	return evt
}

func Fatal(err error) *zerolog.Event {

	evt := logger.Trace()
	if err != nil {
		evt = errLogger.Fatal().Err(err)
	}
	return evt
}

func Panic(err error) *zerolog.Event {

	evt := logger.Trace()
	if err != nil {
		evt = errLogger.Panic().Err(err)
	}
	return evt
}
