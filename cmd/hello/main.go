package main

import (
	"github.com/lycheng/gobjection/pkg/logger"
)

// Version comes from build time
var Version string

func main() {
	log := logger.NewLogger()
	log.WithField("project", "gobjection").WithField("version", Version).Info("Hello world")
}
