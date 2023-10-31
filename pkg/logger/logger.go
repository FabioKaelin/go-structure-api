package logger

import (
	"api1-new/pkg/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Log *zap.Logger
)

func InitializeZapCustomLogger() {
	loglevel := zapcore.DebugLevel
	if config.GinMode == "release" {
		loglevel = zapcore.ErrorLevel
	}
	conf := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(loglevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			MessageKey:   "msg",
			CallerKey:    "file",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	foo, err := conf.Build()
	if err != nil {
		Log.Error(err.Error())
	}
	Log = foo
}
