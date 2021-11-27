package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func Init() (err error) {
	var config zap.Config

	if os.Getenv("ENVIRONMENT") != "production" || os.Getenv("DEBUG") == "true" {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}

	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logger, err := config.Build(zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.FatalLevel))

	if err != nil {
		return
	}

	log = logger

	return
}

func Error(msg interface{}) {
	log.Error(msg.(error).Error())
}

func Info(msg interface{}) {
	log.Info(msg.(string))
}
