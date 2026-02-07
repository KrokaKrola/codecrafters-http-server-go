package server

import (
	"context"
	"log/slog"
	"net"

	"github.com/codecrafters-io/http-server-starter-go/internal/connection"
	"github.com/codecrafters-io/http-server-starter-go/internal/http"
	"github.com/codecrafters-io/http-server-starter-go/internal/router"
)

type Server struct {
	port     string
	host     string
	router   *router.Router
	listener net.Listener
}

func NewServer(host string, port string) *Server {
	return &Server{
		port:   port,
		host:   host,
		router: router.NewRouter(),
	}
}

func (s *Server) Listen(ctx context.Context) error {
	addr := net.JoinHostPort(s.host, s.port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	s.listener = listener

	slog.Info("started server on", "addr", addr)

	go func() {
		<-ctx.Done()
		slog.Info("gracefully shutting down server")
		if err := listener.Close(); err != nil {
			slog.Error("err while trying to gracefully shut down server", "err", err)
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			if ctx.Err() != nil {
				return nil
			}
			slog.Error("err while trying to close listener", "err", err)

			break
		}

		slog.Info("new connection", "conn addr", conn.RemoteAddr())

		c := connection.NewConnection(conn, s.router)
		go c.Handle(ctx)
	}

	return nil
}

func (s *Server) Handle(method http.Method, pattern string, handler router.HandlerFunc) {
	s.router.Handle(method, pattern, handler)
}
