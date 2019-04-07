package main

import (
	"log"
	"net/http"

	"github.com/agungdwiprasetyo/agungdpcms/schema"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func (s *service) ServeHTTP() {
	gqlSchema := schema.LoadSchema()
	schema := graphql.MustParseSchema(gqlSchema, s.handler, graphql.UseStringDescriptions(), graphql.UseFieldResolvers())

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("./static")))
	mux.Handle("/graphiql", &relay.Handler{Schema: schema}) // open host in browser for tool for writing, validating, and testing GraphQL queries.
	mux.Handle("/graphql", &customHandler{schema: schema})

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
