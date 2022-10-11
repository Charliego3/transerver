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
	rsa     *biz.RsaUsecase
	logger  *zap.Logger
}

func NewAccountService(usecase *biz.AccountUsecase, rsa *biz.RsaUsecase, logger *zap.Logger) *AccountService {
	return &AccountService{usecase: usecase, rsa: rsa, logger: logger}
}

func (g *AccountService) RegisterGRPC(s *grpc.Server) {
	acctspb.RegisterAccountServiceServer(s, g)
}

func (g *AccountService) RegisterHTTP(s *runtime.ServeMux) error {
	return acctspb.RegisterAccountServiceHandlerServer(context.Background(), s, g)
}

func (g *AccountService) Routers() ([]string, []string) {
	return nil, nil
}

func (g *AccountService) Register(context.Context, *acctspb.RegisterRequest) (*acctspb.RegisterReply, error) {
	obj, err := g.rsa.FetchObj("", biz.WithRsaNoGen)
	if err != nil {
		return nil, err
	}

	_ = obj.Private
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func (g *AccountService) Login(context.Context, *acctspb.LoginRequest) (*acctspb.LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
