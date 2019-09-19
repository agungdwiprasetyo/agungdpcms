package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func (s *service) ServeHTTP() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(fmt.Sprintf("%s/static", os.Getenv("APP_PATH")))))
	mux.Handle("/graphql", s.graphql.handler)
	mux.Handle("/ws", s.websocket.handler)

	httpPort := fmt.Sprintf(":%d", s.conf.Env.HTTPPort)

	s.httpServer = &http.Server{
		Addr:    httpPort,
		Handler: mux,
	}

	fmt.Println("HTTP Server running on port", httpPort)
	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func (s *service) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to Shutdown HTTP Server :%+v", err)
	}
	log.Print("Server Exited Properly")
}
