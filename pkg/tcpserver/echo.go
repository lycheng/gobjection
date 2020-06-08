package tcpserver

import (
	"bufio"
	"context"
	"net"
	"time"

	"github.com/lycheng/gobjection/pkg/logger"
)

var (
	goodbyeWords       = []byte("Goodbye\n")
	defaultReadTimeout = 2 * time.Second
)

// EchoServerCreater uses to create EchoServer
type EchoServerCreater struct {
}

// EchoServer for echo TCP server
type EchoServer struct {
	ctx  context.Context
	conn net.Conn
	srv  *Server
}

func (esc *EchoServerCreater) new(ctx context.Context, conn net.Conn, srv *Server) Handler {
	es := &EchoServer{
		ctx:  ctx,
		conn: conn,
		srv:  srv,
	}
	return es
}

func (es *EchoServer) handleConn() {
	logger.Logger.WithField("client", es.conn.RemoteAddr()).Info("accept")
	reader := bufio.NewReader(es.conn)
	for {
		select {
		case <-es.ctx.Done():
			logger.Logger.WithField("client", es.conn.RemoteAddr()).Info("cancel")
			es.conn.Write(goodbyeWords)
			return
		default:
			es.conn.SetReadDeadline(time.Now().Add(defaultReadTimeout))
			line, err := reader.ReadBytes('\n')

			// TODO: it can be more elegant
			if e, ok := err.(net.Error); ok && e.Timeout() {
				break
			}

			if err != nil {
				es.conn.Close()
				logger.Logger.WithField("client", es.conn.RemoteAddr()).Warn(err)
				return
			}
			es.conn.Write(line)
		}
	}
}
