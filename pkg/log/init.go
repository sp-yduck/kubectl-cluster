package log

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(logLevel zapcore.Level) {
	level := zap.NewAtomicLevel()
	level.SetLevel(logLevel)
	config := zap.Config{
		Level: level,
		Encoding: "json",
    EncoderConfig: zapcore.EncoderConfig{
        TimeKey:        "Time",
        LevelKey:       "Level",
        NameKey:        "Name",
        CallerKey:      "Caller",
        MessageKey:     "Msg",
        StacktraceKey:  "St",
        EncodeLevel:    zapcore.CapitalLevelEncoder,
        EncodeTime:     zapcore.ISO8601TimeEncoder,
        EncodeDuration: zapcore.StringDurationEncoder,
        EncodeCaller:   zapcore.ShortCallerEncoder,
    },
    OutputPaths:      []string{"stdout"},
    ErrorOutputPaths: []string{"stderr"},
	}
	logger, err := config.Build()
	if err != nil {
		log.Fatalf("failed to create new logger: %v", err)
	}
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
}