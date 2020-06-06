package main

import (
	"github.com/lycheng/gobjection/pkg/envs"
	"github.com/lycheng/gobjection/pkg/logger"
	"github.com/lycheng/gobjection/pkg/tcpserver"
)

// Version comes from build time
var Version string

func main() {
	logger.Logger.WithField("project", "gotcpserver").WithField("version", Version).Info("start tcp server")

	envs.InitWithPrefix("TCPSERVER_")
	network := envs.GetEnvWithDefault("ADDRESS", "tcp")
	address := envs.GetEnvWithDefault("ADDRESS", ":8080")

	server := tcpserver.New(network, address)
	server.Listen()
}
