package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/codecrafters-io/http-server-starter-go/internal/handlers"
	"github.com/codecrafters-io/http-server-starter-go/internal/http"
	"github.com/codecrafters-io/http-server-starter-go/internal/server"
)

const Port = "4221"
const Host = "0.0.0.0"

func main() {
	dirValue := flag.String("directory", "/tmp/data", "Directory for static files")
	flag.Parse()

	s := server.NewServer(Host, Port)

	s.Handle(http.GET, `^/($|index\.html$)`, handlers.HandleHomepage())
	s.Handle(http.GET, `^/echo/(\S+)$`, handlers.HandleEcho())
	s.Handle(http.GET, `^/user-agent(/)?$`, handlers.HandleUserAgent())
	s.Handle(http.GET, `^/files/(\S+)$`, handlers.HandleFilesGet(*dirValue))
	s.Handle(http.POST, `^/files/(\S+)$`, handlers.HandleFilesPost(*dirValue))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := s.Listen(ctx); err != nil {
		slog.Error("err while trying to start server", "err", err)
		os.Exit(1)
	}
}
