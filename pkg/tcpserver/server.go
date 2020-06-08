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
	go s.loop(ctx, wg, ln)
	go func() {
		defer wg.Done()
		signals.WaitToStopped()
		logger.Logger.Info("exit")
		cancel()
	}()
	wg.Wait()
	logger.Logger.Info("EXIT")
}

func (s *Server) loop(ctx context.Context, wg *sync.WaitGroup, ln net.Listener) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.Logger.WithField("stage", "accepted").Warn(err)
			continue
		}

		wg.Add(1)
		client := s.handlerCreator.new(ctx, conn, s)
		go func() {
			defer wg.Done()
			client.handleConn()
		}()
	}
}
