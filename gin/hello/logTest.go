package main

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"io"
	"os"
	"strings"
	"time"
)

func basicTest() {

	// panic and fatal level log break program flow.

	//log.Panic().Msg("panic level log")
	//log.Fatal().Msg("fatal level log")
	log.Log().Msg("no level log")

	log.Error().Msg("error level log")
	log.Warn().Msg("warn level log")
	log.Info().Msg("info level log")
	log.Debug().Msg("debug level log")
	log.Trace().Msg("trace level log")
}

func logObjTest() {

	logger := zerolog.New(os.Stderr)
	logger.Info().Str("foo", "bar").Msg("logger obj info level output")

	sublogger := logger.With().Caller().Timestamp().Str("sublog", "poll").Logger()

	sublogger = sublogger.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	sublogger.Info().Msg("sublogger obj info level output")

	sublogger = sublogger.Output(customLogFormat())
	sublogger.Info().Str("format", "custom").Msg("Custom log output format")

}

func customLogFormat() io.Writer {
	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("***%s****", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}

	return output
}

func errLogTest() {

	err := errors.New("there have an error here")
	log.Error().Err(err).Msg("Error level log")

	// log.Fatal().Err(err).Msg("Error level log")

	//zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// zerolog.ErrorStackMarshaler must be set in order for the stack to output anything.
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	errstack := outer()
	log.Error().Stack().Err(errstack).Msg("Error stack log output")
}

func inner() error {
	return errors.New("seems we have an error here")
}

func middle() error {
	err := inner()
	if err != nil {
		return err
	}
	return nil
}

func outer() error {
	err := middle()
	if err != nil {
		return err
	}
	return nil
}

func multipleOutputTest() {
	mo := zerolog.ConsoleWriter{Out: os.Stdout}

	multi := zerolog.MultiLevelWriter(mo, os.Stdout)
	logger := zerolog.New(multi).With().Timestamp().Logger()

	logger.Info().Msg("Hello Multiple Log Output")

}

func main() {
	/**
	全局的Logger使用比较简单，不需要额外创建;
	zerolog有panic/fatal/error/warn/info/debug/trace这几种级别。我们可以调用SetGlobalLevel()设置全局Logger的日志级别;
	调用Msg()或Send()之后，日志会被输出;
	The default log level for log.Print is debug
	*/

	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	//basicTest()
	//fmt.Println("------------------------------------------------")

	logObjTest()

	//fmt.Println("------------------------------------------------")

	//errLogTest()

	//fmt.Println("------------------------------------------------")

	//multipleOutputTest()

}
