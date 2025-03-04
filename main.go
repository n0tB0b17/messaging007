package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bob17/msg/internal/api"
)

func main() {
	port := 8585
	s := api.NewApiServer(port)
	go func(s *api.APIServer) {
		if err := s.Start(); err != nil && err != http.ErrServerClosed {
			log.Panic("unable to start the web server...")
		}
	}(s)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		fmt.Println("error while shutting down the server")
	}

	fmt.Printf("server shutdown successfully")
}
