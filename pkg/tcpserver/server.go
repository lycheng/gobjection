package tcpserver

import (
	"bufio"
	"net"

	"github.com/lycheng/gobjection/pkg/logger"
)

// Server for TCP server
// reference: https://github.com/firstrow/tcp_server
type Server struct {
	network string
	address string
}

// Client holds info about connection
type Client struct {
	conn net.Conn
	srv  *Server
}

// New returns server instance
func New(network string, address string) *Server {
	srv := &Server{
		address: address,
		network: network,
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
		client := &Client{
			conn: conn,
			srv:  s,
		}

		if err != nil {
			logger.Logger.WithField("stage", "accept").Warn(err)
		}
		go client.Work()
	}
}

// Work handles TCP traffic
func (c *Client) Work() {
	logger.Logger.WithField("stage", "connected").Info("accept", c.conn.RemoteAddr())
	reader := bufio.NewReader(c.conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			c.conn.Close()
			return
		}
		logger.Logger.WithField("stage", "read").Info(message)
	}
}
