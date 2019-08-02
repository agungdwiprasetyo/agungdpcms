package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func (s *service) ServeHTTP() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(fmt.Sprintf("%s/static", os.Getenv("APP_PATH")))))
	mux.Handle("/graphql", s.graphql.handler)
	mux.Handle("/ws", s.websocket.handler)

	httpPort := fmt.Sprintf(":%d", s.conf.Env.HTTPPort)
	fmt.Println("HTTP Server running on port", httpPort)
	err := http.ListenAndServe(httpPort, mux)
	if err != nil {
		log.Fatal(err)
	}
}
