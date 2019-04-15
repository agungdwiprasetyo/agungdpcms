package main

import (
	"time"

	"github.com/agungdwiprasetyo/agungdpcms/config"
	"github.com/agungdwiprasetyo/agungdpcms/middleware"
	jwtToken "github.com/agungdwiprasetyo/agungdpcms/shared/token"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat"
	cd "github.com/agungdwiprasetyo/agungdpcms/src/chat/delivery"
	cu "github.com/agungdwiprasetyo/agungdpcms/src/chat/usecase"
	rd "github.com/agungdwiprasetyo/agungdpcms/src/resume/delivery"
	ru "github.com/agungdwiprasetyo/agungdpcms/src/resume/usecase"
	ud "github.com/agungdwiprasetyo/agungdpcms/src/user/delivery"
	uu "github.com/agungdwiprasetyo/agungdpcms/src/user/usecase"
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
	User   *ud.GraphqlHandler
}

func newService(conf *config.Config) *service {
	// init middleware
	// midd := middleware.NewBasicAuth(conf)
	token := jwtToken.New(conf.PrivateKey, conf.PublicKey, 24*time.Hour)
	midd := middleware.NewBearer(conf, token)

	resumeUsecase := ru.NewResumeUsecase(conf)
	resumeHandler := rd.New(resumeUsecase, midd)

	chatUsecase := cu.New(conf)
	chatGqlHandler := cd.NewGraphqlHandler(chatUsecase, midd)

	userUsecase := uu.NewUserUsecase(conf, token)
	userGqlHandler := ud.NewGraphqlHandler(userUsecase, midd)

	srv := new(service)
	srv.conf = conf

	srv.websocket.server = chat.NewServer()
	srv.websocket.handler = cd.NewWebsocketHandler(srv.websocket.server, chatUsecase)

	srv.handler = &handler{
		Resume: resumeHandler,
		Chat:   chatGqlHandler,
		User:   userGqlHandler,
	}

	return srv
}
