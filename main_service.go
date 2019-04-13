package main

import (
	"github.com/agungdwiprasetyo/agungdpcms/config"
	"github.com/agungdwiprasetyo/agungdpcms/middleware"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat"
	cd "github.com/agungdwiprasetyo/agungdpcms/src/chat/delivery"
	cu "github.com/agungdwiprasetyo/agungdpcms/src/chat/usecase"
	rd "github.com/agungdwiprasetyo/agungdpcms/src/resume/delivery"
	ru "github.com/agungdwiprasetyo/agungdpcms/src/resume/usecase"
)

type service struct {
	conf      *config.Config
	handler   *handler
	websocket struct {
		server  *chat.Server
		handler *cd.WsHandler
	}
}

type handler struct {
	Resume *rd.ResumeHandler
	Chat   *cd.GraphqlHandler
}

func newService(conf *config.Config) *service {
	// init middleware
	midd := middleware.NewBasicAuth(conf)

	resumeUsecase := ru.NewResumeUsecase(conf)
	resumeHandler := rd.New(resumeUsecase, midd)

	chatUsecase := cu.New(conf)
	chatGqlHandler := cd.NewGraphqlHandler(chatUsecase)

	srv := new(service)
	srv.conf = conf

	srv.websocket.server = chat.NewServer()
	srv.websocket.handler = cd.NewWebsocketHandler(srv.websocket.server, chatUsecase)

	srv.handler = &handler{resumeHandler, chatGqlHandler}

	return srv
}
