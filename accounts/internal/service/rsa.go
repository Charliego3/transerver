package service

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/transerver/accounts/internal/biz"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type RsaService struct {
	// TODO: Unimplemented pb
	usecase *biz.RsaUsecase
	logger  *zap.Logger
}

func NewRsaService(usecase *biz.RsaUsecase, logger *zap.Logger) *RsaService {
	return &RsaService{usecase: usecase, logger: logger}
}

func (g *RsaService) RegisterGRPC(s *grpc.Server) {

}

func (g *RsaService) RegisterHTTP(s *runtime.ServeMux) error {
	return nil
}
