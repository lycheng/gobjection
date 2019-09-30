package main

import (
	"github.com/lycheng/gobjection/pkg/logger"
)

func main() {
	log := logger.NewLogger()
	log.WithField("project", "gobjection").Info("Hello world")
}
