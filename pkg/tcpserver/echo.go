package tcpserver

import (
	"bufio"
	"net"

	"github.com/lycheng/gobjection/pkg/logger"
)

// EchoServerCreater uses to create EchoServer
type EchoServerCreater struct {
}

// EchoServer for echo TCP server
type EchoServer struct {
	conn net.Conn
	srv  *Server
}

func (esc *EchoServerCreater) new(conn net.Conn, srv *Server) Clienter {
	es := &EchoServer{
		conn: conn,
		srv:  srv,
	}
	return es
}

func (es *EchoServer) handleConn() {
	logger.Logger.WithField("stage", "connected").Info("accept ", es.conn.RemoteAddr())
	reader := bufio.NewReader(es.conn)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			es.conn.Close()
			logger.Logger.WithField("stage", "exit").Warn(err)
			break
		}

		es.conn.Write(line)
	}
	logger.Logger.WithField("stage", "end").Info("connection closed")
}
