package chat

import (
	"github.com/agungdwiprasetyo/agungdpcms/config"
	"github.com/agungdwiprasetyo/agungdpcms/middleware"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/delivery"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/repository"
	"github.com/agungdwiprasetyo/agungdpcms/src/chat/usecase"
)

// Module model
type Module struct {
	Handler    *delivery.GraphQLHandler
	Usecase    usecase.Chat
	Repository *repository.Repository
}

// New chat module constructor
func New(conf *config.Config, midd middleware.Middleware) *Module {
	repo := repository.NewRepository(conf.DB)
	uc := usecase.New(repo)
	handler := delivery.NewGraphQLHandler(uc, midd)

	return &Module{
		Handler:    handler,
		Usecase:    uc,
		Repository: repo,
	}
}
