package main

import (
	"github.com/agungdwiprasetyo/agungdpcms/config"
	"github.com/agungdwiprasetyo/agungdpcms/middleware"
	jwtToken "github.com/agungdwiprasetyo/agungdpcms/shared/token"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat"
	"github.com/agungdwiprasetyo/agungdpcms/src/master"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume"
	"github.com/agungdwiprasetyo/agungdpcms/src/user"
	"github.com/agungdwiprasetyo/agungdpcms/websocket"
)

type service struct {
	conf *config.Config

	graphql struct {
		resolver *graphqlResolver
		handler  *graphqlHandler
	}

	websocket struct {
		server  *websocket.Server
		handler *websocket.Handler
	}
}

func newService(conf *config.Config) *service {
	// init middleware
	// midd := middleware.NewBasicAuth(conf)
	token := jwtToken.New(conf.PrivateKey, conf.PublicKey, conf.Env.TokenAge)
	midd := middleware.NewBearer(conf, token)
	wsServer := websocket.NewServer(&conf.Env)

	// init master module
	masterModule := master.New(conf, midd)
	// init user module
	userModule := user.New(conf, midd, token)
	// init resume module
	resumeModule := resume.New(conf, midd)
	// init chat module
	chatModule := chat.New(conf, midd)

	gqlResolver := &graphqlResolver{
		Resume: resumeModule.Handler,
		Chat:   chatModule.Handler,
		User:   userModule.Handler,
		Master: masterModule.Handler,
	}

	srv := new(service)
	srv.conf = conf
	srv.websocket.server = wsServer
	srv.websocket.handler = websocket.NewWebsocketHandler(wsServer, chatModule.Usecase)
	srv.graphql.resolver = gqlResolver
	srv.graphql.handler = newGraphQLHandler(&conf.Env, gqlResolver)

	return srv
}
