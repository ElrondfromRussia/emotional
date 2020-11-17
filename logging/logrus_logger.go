package emotional

import (
	"github.com/sirupsen/logrus"
	"os"
)

var SirLogger = NewSirLogger()

func NewSirLogger() *logrus.Logger {
	logger := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &logrus.TextFormatter{
			FullTimestamp:          true,
			TimestampFormat:        "2006-01-02 15:04:05",
			ForceColors:            true,
			DisableLevelTruncation: true,
		},
	}
	logger.SetReportCaller(true)
	return logger
}
