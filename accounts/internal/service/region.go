package service

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/transerver/accounts/internal/biz"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type RegionService struct {
	// TODO: Unimplemented pb
	usecase *biz.RegionUsecase
	logger  *zap.Logger
}

func NewRegionService(usecase *biz.RegionUsecase, logger *zap.Logger) *RegionService {
	return &RegionService{usecase: usecase, logger: logger}
}

func (g *RegionService) RegisterGRPC(s *grpc.Server) {

}

func (g *RegionService) RegisterHTTP(s *runtime.ServeMux) error {
	return nil
}
