package main

import (
	"net/http"
	"os"

	"github.com/agungdwiprasetyo/agungdpcms/config"
	"github.com/agungdwiprasetyo/agungdpcms/middleware"
	"github.com/agungdwiprasetyo/agungdpcms/schema/jsonschema"
	jwtToken "github.com/agungdwiprasetyo/agungdpcms/shared/token"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat"
	"github.com/agungdwiprasetyo/agungdpcms/src/master"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume"
	"github.com/agungdwiprasetyo/agungdpcms/src/user"
	"github.com/agungdwiprasetyo/agungdpcms/websocket"
)

type service struct {
	conf       *config.Config
	httpServer *http.Server

	websocket struct {
		server  *websocket.Server
		handler *websocket.Handler
	}

	resumeModule *resume.Module
	masterModule *master.Module
	chatModule   *chat.Module
	userModule   *user.Module
}

func newService(conf *config.Config) *service {
	// init middleware
	// midd := middleware.NewBasicAuth(conf)
	token := jwtToken.New(conf.PrivateKey, conf.PublicKey, config.GlobalEnv.TokenAge)
	midd := middleware.NewBearer(conf, token)
	wsServer := websocket.NewServer()

	// load json schema
	if err := jsonschema.Load(os.Getenv("APP_PATH") + "/schema/jsonschema/"); err != nil {
		panic(err)
	}

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
	srv.websocket.server = wsServer
	srv.websocket.handler = websocket.NewWebsocketHandler(wsServer, chatModule.Usecase)

	srv.resumeModule = resumeModule
	srv.masterModule = masterModule
	srv.chatModule = chatModule
	srv.userModule = userModule
	return srv
}
