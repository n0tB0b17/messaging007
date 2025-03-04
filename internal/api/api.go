package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type APIServer struct {
	Port int
	s    *http.Server
}

func NewApiServer(port int) *APIServer {
	return &APIServer{
		Port: port,
	}
}

func (s *APIServer) Start() error {
	addr := fmt.Sprintf("%s:%d", "", s.Port)
	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length"},
	})

	handler := c.Handler(router)
	s.s = &http.Server{Addr: addr, Handler: handler}

	fmt.Printf("Server is up and running %d \n", s.Port)
	return s.s.ListenAndServe()
}

func (s *APIServer) Shutdown(ctx context.Context) error {
	contxt, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	fmt.Println("Shutting down api server")
	return s.s.Shutdown(contxt)
}
