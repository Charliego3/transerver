package service

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/transerver/accounts/internal/biz"
	"github.com/transerver/protos/acctspb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccountService struct {
	acctspb.UnimplementedAccountServiceServer

	usecase *biz.AccountUsecase
	logger  *zap.Logger
}

func NewAccountService(usecase *biz.AccountUsecase, logger *zap.Logger) *AccountService {
	return &AccountService{usecase: usecase, logger: logger}
}

func (g *AccountService) RegisterGRPC(s *grpc.Server) {
	acctspb.RegisterAccountServiceServer(s, g)
}

func (g *AccountService) RegisterHTTP(s *runtime.ServeMux) error {
	return acctspb.RegisterAccountServiceHandlerServer(context.Background(), s, g)
}

func (g *AccountService) Register(context.Context, *acctspb.RegisterRequest) (*acctspb.RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func (g *AccountService) Login(context.Context, *acctspb.LoginRequest) (*acctspb.LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}