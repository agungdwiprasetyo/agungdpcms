package resume

import (
	"github.com/agungdwiprasetyo/agungdpcms/config"
	"github.com/agungdwiprasetyo/agungdpcms/middleware"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/delivery"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/repository"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/usecase"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/validation"
)

// Module model
type Module struct {
	Handler    *delivery.GraphQLHandler
	Usecase    usecase.Resume
	Repository *repository.Repository
}

// New resume module constructor
func New(conf *config.Config, midd middleware.Middleware) *Module {
	v := validation.New()

	repo := repository.NewRepository(conf.DB)
	uc := usecase.NewResumeUsecase(repo)
	handler := delivery.NewGraphQLHandler(uc, midd, v)

	return &Module{
		Handler:    handler,
		Usecase:    uc,
		Repository: repo,
	}
}
