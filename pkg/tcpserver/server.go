package tcpserver

import (
	"context"
	"net"
	"sync"

	"github.com/lycheng/gobjection/pkg/logger"
	"github.com/lycheng/gobjection/pkg/signals"
)

// Server for TCP server
// reference: https://github.com/firstrow/tcp_server
type Server struct {
	network        string
	address        string
	handlerCreator HandlerCreator
}

// New returns server instance
func New(network string, address string, creater HandlerCreator) *Server {
	srv := &Server{
		address:        address,
		network:        network,
		handlerCreator: creater,
	}
	return srv
}

// Run starts the TCP server
func (s *Server) Run() {
	ln, err := net.Listen(s.network, s.address)
	if err != nil {
		logger.Logger.Fatal(err)
	}

	defer ln.Close()

	ctx, cancel := context.WithCancel(context.Background())

	wg := &sync.WaitGroup{}
	wg.Add(1) // For the waiting signals
	go func() {
		defer wg.Done()
		signals.WaitToStopped()
		cancel()
	}()

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				select {
				case <-ctx.Done():
					return
				default:
					logger.Logger.WithField("stage", "accepted").Warn(err)
				}
				continue
			}

			wg.Add(1)
			client := s.handlerCreator.new(ctx, conn, s)
			go func() {
				defer wg.Done()
				client.handleConn()
			}()
		}
	}()

	wg.Wait()
	logger.Logger.Info("EXIT")
}
