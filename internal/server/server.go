package server

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/http-server-starter-go/internal/connection"
)

type Server struct {
	port     int
	host     string
	dirValue string
}

func NewServer(host string, port int, dirValue string) *Server {
	return &Server{
		port:     port,
		host:     host,
		dirValue: dirValue,
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

		c := connection.NewConnection(conn, s.dirValue)
		go c.Handle()
	}

	err = listener.Close()
	if err != nil {
		return err
	}

	return nil
}
