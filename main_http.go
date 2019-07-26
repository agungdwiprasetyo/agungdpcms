package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/agungdwiprasetyo/agungdpcms/schema"
	"github.com/graph-gophers/graphql-go"
)

func (s *service) ServeHTTP() {
	// init graphql
	gqlSchema := schema.LoadSchema()
	schema := graphql.MustParseSchema(gqlSchema, s.handler, graphql.UseStringDescriptions(), graphql.UseFieldResolvers())

	gqlHandler := newGraphQLHandler(schema, s.conf)

	// open host in browser for tool for writing, validating, and testing GraphQL queries.
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(fmt.Sprintf("%s/static", os.Getenv("APP_PATH")))))
	mux.Handle("/graphql", gqlHandler)

	// mux.Handle("/", http.FileServer(http.Dir(fmt.Sprintf("%s/static/ws", os.Getenv("APP_PATH")))))
	mux.HandleFunc("/ws", s.websocket.handler.Socket)

	httpPort := fmt.Sprintf(":%d", s.conf.Env.HTTPPort)
	fmt.Println("HTTP Server running on port", httpPort)
	err := http.ListenAndServe(httpPort, mux)
	if err != nil {
		log.Fatal(err)
	}
}
