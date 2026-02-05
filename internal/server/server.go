package server

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/http-server-starter-go/internal/connection"
)

type Server struct {
	port int
	host string
}

func NewServer(host string, port int) *Server {
	return &Server{
		port: port,
		host: host,
	}
}

func (s *Server) Listen() error {
	addr := net.JoinHostPort(s.host, fmt.Sprintf("%d", s.port))
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	fmt.Printf("server listening on %s\n", addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection", err)
			break
		}

		fmt.Printf("new connection from %s\n", conn.RemoteAddr())

		c := connection.NewConnection(conn)
		go c.Handle()
	}

	err = listener.Close()
	if err != nil {
		return err
	}

	return nil
}
