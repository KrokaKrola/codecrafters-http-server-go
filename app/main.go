package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/internal/server"
)

const PORT = 4221
const HOST = "0.0.0.0"

func main() {
	dirValue := flag.String("directory", "/tmp/data", "Directory for static files")
	flag.Parse()

	s := server.NewServer(HOST, PORT, *dirValue)
	if err := s.Listen(); err != nil {
		fmt.Println("err while trying to start server", err)
		os.Exit(1)
	}
}
