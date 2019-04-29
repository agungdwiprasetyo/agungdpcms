package main

/*
	Custom graphql handler
*/

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/agungdwiprasetyo/agungdpcms/config"
	"github.com/agungdwiprasetyo/agungdpcms/shared"
	"github.com/graph-gophers/graphql-go"
)

type customHandler struct {
	schema *graphql.Schema
	conf   *config.Config
}

func newCustomHandler(schema *graphql.Schema, conf *config.Config) *customHandler {
	return &customHandler{
		schema: schema, conf: conf,
	}
}

func (h *customHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// handle cors
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", h.conf.Env.CORSWhitelist)
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
