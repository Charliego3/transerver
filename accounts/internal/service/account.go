package service

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/transerver/accounts/internal/biz"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type AccountService struct {
	// TODO: Unimplemented pb
	usecase *biz.AccountUsecase
	logger  *zap.Logger
}

func NewAccountService(usecase *biz.AccountUsecase, logger *zap.Logger) *AccountService {
	return &AccountService{usecase: usecase, logger: logger}
}

func (g *AccountService) RegisterGRPC(s *grpc.Server) {

}

func (g *AccountService) RegisterHTTP(s *runtime.ServeMux) error {
	return nil
}
