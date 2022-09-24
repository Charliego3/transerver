package service

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/transerver/cli/templates/internal/biz"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type GreeterService struct {
	// TODO: unimplement pb
	usecase *biz.GreeterUsecase
	logger  *zap.Logger
}

func NewGreeterSerivce(usecase *biz.GreeterUsecase, logger *zap.Logger) *GreeterService {
	return &GreeterService{usecase: usecase, logger: logger}
}

func (g *GreeterService) RegisterGRPC(s *grpc.Server) {

}

func (g *GreeterService) RegisterHTTP(s *runtime.ServeMux) error {
	return nil
}
