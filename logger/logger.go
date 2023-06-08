package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	defaultLevel = zapcore.InfoLevel
)

var (
	globalLogger *zap.Logger
	level        zap.AtomicLevel
)

type flusher struct {
	core zapcore.Core
}

func NewEnvironmentLogger(envLevel string) *zap.Logger {

	level := strings.ToUpper(envLevel)

	switch level {
	case "DEBUG", "Debug", "debug":
		return Initialize(zap.DebugLevel)
	case "INFO", "Info", "info":
		return Initialize(zap.InfoLevel)
	case "WARN", "Warn", "warn":
		return Initialize(zap.WarnLevel)
	case "ERROR", "Error", "error":
		return Initialize(zap.ErrorLevel)
	}
	return Initialize(defaultLevel)
}

func Initialize(l zapcore.Level) *zap.Logger {
	level.SetLevel(l)
	level = zap.NewAtomicLevelAt(l)

	logConfig := zap.NewProductionConfig()
	logConfig.Level = level
	logConfig.EncoderConfig.MessageKey = "message"
	logConfig.EncoderConfig.TimeKey = "@timestamp"
	logConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	//globalLogger.WithOptions(zap.WithFatalHook(&flusher{core: globalLogger.Core()}))
	globalLogger = setServiceName(globalLogger)
	globalLogger = setServiceEnvironment(globalLogger)
	globalLogger.Debug("logger initialized")

	return globalLogger
}

func setServiceName(logger *zap.Logger) *zap.Logger {
	serviceName := strings.ToUpper(os.Getenv("apex-network-api"))
	if serviceName != "" {
		return logger.With(zap.String("service_name", serviceName))
	}
	return logger
}

func setServiceEnvironment(logger *zap.Logger) *zap.Logger {
	appEnv := strings.ToUpper(os.Getenv("local"))
	if appEnv != "" {
		return logger.With(zap.String("env", appEnv))
	}
	return logger
}
