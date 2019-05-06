package main

import (
	"github.com/agungdwiprasetyo/agungdpcms/config"
	"github.com/agungdwiprasetyo/agungdpcms/middleware"
	jwtToken "github.com/agungdwiprasetyo/agungdpcms/shared/token"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat"
	cd "github.com/agungdwiprasetyo/agungdpcms/src/chat/delivery"
	"github.com/agungdwiprasetyo/agungdpcms/src/master"
	md "github.com/agungdwiprasetyo/agungdpcms/src/master/delivery"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume"
	rd "github.com/agungdwiprasetyo/agungdpcms/src/resume/delivery"
	"github.com/agungdwiprasetyo/agungdpcms/src/user"
	ud "github.com/agungdwiprasetyo/agungdpcms/src/user/delivery"
	"github.com/agungdwiprasetyo/agungdpcms/websocket"
)

type service struct {
	conf      *config.Config
	handler   *handler
	websocket struct {
		server  *websocket.Server
		handler *cd.WsHandler
	}
}

type handler struct {
	Resume *rd.GraphQLHandler
	Chat   *cd.GraphQLHandler
	User   *ud.GraphQLHandler
	Master *md.GraphQLHandler
}

func newService(conf *config.Config) *service {
	// init middleware
	// midd := middleware.NewBasicAuth(conf)
	token := jwtToken.New(conf.PrivateKey, conf.PublicKey, conf.Env.TokenAge)
	midd := middleware.NewBearer(conf, token)

	// init master module
	masterModule := master.New(conf, midd)
	// init user module
	userModule := user.New(conf, midd, token)
	// init resume module
	resumeModule := resume.New(conf, midd)
	// init chat module
	chatModule := chat.New(conf, midd)

	srv := new(service)
	srv.conf = conf
	srv.websocket.server = websocket.NewServer()
	srv.websocket.handler = cd.NewWebsocketHandler(srv.websocket.server, chatModule.Usecase)
	srv.handler = &handler{
		Resume: resumeModule.Handler,
		Chat:   chatModule.Handler,
		User:   userModule.Handler,
		Master: masterModule.Handler,
	}

	return srv
}
