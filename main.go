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
	"github.com/bob17/msg/internal/config"
	"github.com/bob17/msg/internal/db"
	"github.com/bob17/msg/internal/nats"
	"github.com/bob17/msg/internal/service"
)

func main() {
	cfg := config.GetConfig()
	s := api.NewApiServer(cfg.Port)
	db := db.NewDB()
	natClient, err := nats.GetNatsClient(cfg.NATURL)
	if err != nil {
		fmt.Printf("unable to fetch NAT client: %v \n", err)
	}

	service.InitializeService(db, natClient)

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
