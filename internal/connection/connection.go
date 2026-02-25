package connection

import (
	"bufio"
	"context"
	"errors"
	"io"
	"log/slog"
	"net"

	"github.com/codecrafters-io/http-server-starter-go/internal/request"
	"github.com/codecrafters-io/http-server-starter-go/internal/router"
)

type Connection struct {
	conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
	router *router.Router
}

func NewConnection(conn net.Conn, router *router.Router) *Connection {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	return &Connection{
		conn:   conn,
		reader: reader,
		writer: writer,
		router: router,
	}
}

func (c *Connection) Handle(ctx context.Context) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			slog.Error("err while trying to close connection", "err", err)
		}
	}(c.conn)

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		req := request.NewRequest(c.reader, ctx)
		if err := req.Parse(); err != nil {
			if errors.Is(err, io.EOF) {
				return
			}

			slog.Error("error while parsing request", "err", err)
			return
		}

		res := c.router.Match(req)

		if _, err := c.writer.WriteString(res.String()); err != nil {
			slog.Error("error while writing response", "err", err)
			return
		}

		if err := c.writer.Flush(); err != nil {
			slog.Error("error trying to flush writer", "err", err)
			return
		}
	}
}
