package utils

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// SetUpDefaultLogger ... is used to bootstrap logging since some logging configurations are in the app config
func SetUpDefaultLogger() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.Kitchen})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

// SetUpLoggerFromConfig ...
func SetUpLoggerFromConfig() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.Kitchen})
	level := GetLevel(UseEnvOrDefault("LOG_LEVEL", "INFO"))
	log.Info().Msgf("Setting log level to %v", level)
	zerolog.SetGlobalLevel(level)
}

// GetLevel ...
func GetLevel(l string) zerolog.Level {
	log.Info().Msgf("debug level is %s", l)
	switch l {
	case "TRACE":
		return zerolog.TraceLevel
	case "DEBUG":
		return zerolog.DebugLevel
	case "INFO":
		return zerolog.InfoLevel
	case "WARN":
		return zerolog.WarnLevel
	case "ERROR":
		return zerolog.ErrorLevel
	case "FATAL":
		return zerolog.FatalLevel
	case "PANIC":
		return zerolog.PanicLevel
	case "NONE":
		return zerolog.NoLevel
	case "DISABLED":
		return zerolog.Disabled
	default:
		log.Fatal().Msg("Invalid environment variable 'LOG_LEVEL' passed. Valid values are DISABLED, DEBUG, INFO, ERROR etc.")
		return zerolog.NoLevel
	}
}
