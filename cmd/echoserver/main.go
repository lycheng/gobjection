package main

import (
	"github.com/lycheng/gobjection/pkg/envs"
	"github.com/lycheng/gobjection/pkg/logger"
	"github.com/lycheng/gobjection/pkg/tcpserver"
)

// Version comes from build time
var Version string

func main() {

	envs.InitWithPrefix("TCPSERVER_")
	network := envs.GetEnvWithDefault("NETWORK", "tcp")
	address := envs.GetEnvWithDefault("ADDRESS", ":8080")

	ecs := &EchoServerCreater{}
	logger.Logger.
		WithField("project", "tcpserver").
		WithField("network", network).
		WithField("address", address).
		Info("start")

	server := tcpserver.New(network, address, ecs)
	server.Run()
}
