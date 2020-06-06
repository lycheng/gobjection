package logger

import "github.com/sirupsen/logrus"

// Logger for default usage
var Logger *logrus.Logger

func init() {
	Logger = NewLogger()
}

// NewLogger returns a new logger instance
func NewLogger() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	return log
}
