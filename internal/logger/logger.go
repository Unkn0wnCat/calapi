package logger

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

func StartLogger() {
	Logger, _ = zap.NewProduction()

	if viper.GetBool("development") {
		Logger, _ = zap.NewDevelopment()
	}

	Sugar = Logger.Sugar()
}
