package delivery

import (
	"context"
	"fmt"

	resumegrpc "github.com/agungdwiprasetyo/agungdpcms/schema/proto/resume"
	"github.com/agungdwiprasetyo/agungdpcms/src/resume/usecase"
	"google.golang.org/grpc"
)

// GRPCHandler model
type GRPCHandler struct {
	resumeUsecase usecase.Resume
}

// NewGRPCHandler create new grpc handler
func NewGRPCHandler(resumeUsecase usecase.Resume) *GRPCHandler {
	return &GRPCHandler{
		resumeUsecase: resumeUsecase,
	}
}

// Register grpc server
func (h *GRPCHandler) Register(server *grpc.Server) {
	resumegrpc.RegisterResumeHandlerServer(server, h)
}

// GetAllResume rpc
func (h *GRPCHandler) GetAllResume(ctx context.Context, filter *resumegrpc.Filter) (*resumegrpc.Resume, error) {
	fmt.Println(filter)
	return &resumegrpc.Resume{
		Name: "TESTING GRPC",
	}, nil
}
