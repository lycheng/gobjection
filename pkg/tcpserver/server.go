package tcpserver

import (
	"net"

	"github.com/lycheng/gobjection/pkg/logger"
)

// Server for TCP server
// reference: https://github.com/firstrow/tcp_server
type Server struct {
	network string
	address string
	creater ClientCreator
}

// ClientCreator is the tcp client creator interface
type ClientCreator interface {
	new(conn net.Conn, srv *Server) Clienter
}

// Clienter use for tcp client
type Clienter interface {
	handleConn()
}

// New returns server instance
func New(network string, address string, creater ClientCreator) *Server {
	srv := &Server{
		address: address,
		network: network,
		creater: creater,
	}
	return srv
}

// Listen starts network server
func (s *Server) Listen() {
	ln, err := net.Listen(s.network, s.address)
	if err != nil {
		logger.Logger.Fatal(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			logger.Logger.WithField("stage", "accepted").Warn(err)
			continue
		}

		client := s.creater.new(conn, s)
		go client.handleConn()
	}
}
