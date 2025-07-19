package logger

import (
	"fmt"
	"go-skeleton/config"
	"log"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	globalLogger *zap.Logger
	sugar        *zap.SugaredLogger
)

// Init initializes the global logger with the provided configuration
func Init(logConfig config.LoggerConfig) {
	// Parse log level
	var level zapcore.Level
	if err := level.UnmarshalText([]byte(strings.ToLower(logConfig.Level))); err != nil {
		level = zapcore.InfoLevel
	}

	// Configure encoder
	var encoderConfig zapcore.EncoderConfig
	if logConfig.Development {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	}

	// Choose an encoder type
	var encoder zapcore.Encoder
	if logConfig.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	// writeSyncer always defaults to stdout for now
	writeSyncer := zapcore.AddSync(os.Stdout)

	// Create core
	core := zapcore.NewCore(encoder, writeSyncer, level)

	// Create a logger with options
	opts := []zap.Option{
		zap.AddStacktrace(zapcore.ErrorLevel),
	}

	if !logConfig.DisableCaller {
		opts = append(opts, zap.AddCaller())
	}

	if !logConfig.DisableStacktrace {
		opts = append(opts, zap.AddStacktrace(zapcore.ErrorLevel))
	}

	if logConfig.Development {
		opts = append(opts, zap.Development())
	}

	globalLogger = zap.New(core, opts...)
	sugar = globalLogger.Sugar()
}

// GetLogger returns the global zap logger
func GetLogger() *zap.Logger {
	if globalLogger == nil {
		// Fallback to nop logger if not initialized
		return zap.NewNop()
	}
	return globalLogger
}

// GetSugar returns the global sugared logger
func GetSugar() *zap.SugaredLogger {
	if sugar == nil {
		// Fallback to nop logger if not initialized
		return zap.NewNop().Sugar()
	}
	return sugar
}

// Sync flushes any buffered log entries
func Sync() {
	if globalLogger != nil {
		if err := globalLogger.Sync(); err != nil {
			log.Fatal(fmt.Errorf("failed to flushing log: %w", err))
		}
	}
	// no explicit return needed
}

// Info Convenience functions for common logging patterns
func Info(msg string, fields ...zap.Field) {
	GetLogger().Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	GetLogger().Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	GetLogger().Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	GetLogger().Warn(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	GetLogger().Fatal(msg, fields...)
}

// Info Sugared convenience functions
func Infof(template string, args ...any) {
	GetSugar().Infof(template, args...)
}

func Errorf(template string, args ...any) {
	GetSugar().Errorf(template, args...)
}

func Debugf(template string, args ...any) {
	GetSugar().Debugf(template, args...)
}

func Warnf(template string, args ...any) {
	GetSugar().Warnf(template, args...)
}

func Fatalf(template string, args ...any) {
	GetSugar().Fatalf(template, args...)
}
