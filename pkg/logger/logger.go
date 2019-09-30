package logger

import "github.com/sirupsen/logrus"

// NewLogger returns a new logger instance
func NewLogger() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	return log
}
