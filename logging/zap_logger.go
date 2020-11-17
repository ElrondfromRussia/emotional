package emotional

import (
	"go.uber.org/zap"
)

var ZapZapLogger = NewZapZapLogger()

func NewZapZapLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return logger
}
