package logger

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New creates a zap logger for console and also setup an output for syslog.
func New(level string) *zap.Logger {
	var lvl zapcore.Level
	if err := lvl.Set(level); err != nil {
		log.Printf("cannot parse log level %s: %s", level, err)

		lvl = zapcore.WarnLevel
	}

	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	defaultCore := zapcore.NewCore(encoder, zapcore.Lock(zapcore.AddSync(os.Stderr)), lvl)
	cores := []zapcore.Core{
		defaultCore,
	}

	core := zapcore.NewTee(cores...)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))

	return logger
}
