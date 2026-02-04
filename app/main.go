package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/internal/server"
)

const PORT = 4221
const HOST = "0.0.0.0"

func main() {
	s := server.NewServer(HOST, PORT)
	if err := s.Listen(); err != nil {
		fmt.Println("err while trying to start server", err)
		os.Exit(1)
	}
}
