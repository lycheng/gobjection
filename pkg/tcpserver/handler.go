package tcpserver

import (
	"context"
	"net"
)

// HandlerCreator is the TCP client creator interface
type HandlerCreator interface {
	New(ctx context.Context, conn net.Conn, srv *Server) Handler
}

// Handler use for TCP client
type Handler interface {
	HandleConn()
}
