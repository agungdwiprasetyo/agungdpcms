package main

/*
	Custom graphql handler
*/

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	graphqlschema "github.com/agungdwiprasetyo/agungdpcms/schema/graphql"
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/agungdwiprasetyo/agungdpcms/shared/logger"
	cd "github.com/agungdwiprasetyo/agungdpcms/src/chat/delivery"
	md "github.com/agungdwiprasetyo/agungdpcms/src/master/delivery"
	rd "github.com/agungdwiprasetyo/agungdpcms/src/resume/delivery"
	ud "github.com/agungdwiprasetyo/agungdpcms/src/user/delivery"
	"github.com/graph-gophers/graphql-go"
)

func (s *service) graphQLHandler() *graphqlHandler {
	var resolver = struct {
		Resume *rd.GraphQLHandler
		Chat   *cd.GraphQLHandler
		User   *ud.GraphQLHandler
		Master *md.GraphQLHandler
	}{
		s.resumeModule.Handler,
		s.chatModule.Handler,
		s.userModule.Handler,
		s.masterModule.Handler,
	}
	gqlSchema := graphqlschema.LoadSchema()
	return &graphqlHandler{
		schema: graphql.MustParseSchema(gqlSchema, &resolver,
			graphql.UseStringDescriptions(),
			graphql.UseFieldResolvers(),
			graphql.Logger(&logger.PanicLogger{}),
			graphql.Tracer(&logger.NoopTracer{})),
	}
}

type graphqlHandler struct {
	schema *graphql.Schema
}

func (h *graphqlHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle cors
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE")
	if r.Method == http.MethodOptions {
		return
	}

	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := json.Unmarshal(body, &params); err != nil {
		params.Query = string(body)
	}

	ctx := context.WithValue(r.Context(), shared.ContextKey("headers"), r.Header)
	response := h.schema.Exec(ctx, params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
