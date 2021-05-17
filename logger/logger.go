package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",   // e.g. "Level": "info"
			TimeKey:      "time",    // e.g. "Time": "2021-01-01T15:30:00Z"
			MessageKey:   "message", // e.g. "msg": "this is a message"
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	if logger, err = logConfig.Build(); err != nil {
		panic(fmt.Sprintf("Cannot start logger: %s", err))
	}
}

func GetLogger() *zap.Logger {
	return logger
}

func Info(msg string, tags ...zap.Field) {
	logger.Info(msg, tags...)
	_ = logger.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("Error", err))

	logger.Error(msg, tags...)
	_ = logger.Sync()
}
