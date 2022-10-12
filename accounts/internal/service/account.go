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
	pubcase *biz.PubUsecase
	logger  *zap.Logger
}

func NewAccountService(usecase *biz.AccountUsecase, rsa *biz.PubUsecase, logger *zap.Logger) *AccountService {
	return &AccountService{usecase: usecase, pubcase: rsa, logger: logger}
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

func (g *AccountService) validateUniqueId(uniqueId string) error {
	if !g.pubcase.ValidateUniqueId(uniqueId) {
		return status.Errorf(codes.InvalidArgument, "Invalid unique: %s", uniqueId)
	}
	return nil
}

func (g *AccountService) Register(_ context.Context, req *acctspb.RegisterRequest) (*acctspb.RegisterReply, error) {
	if err := g.validateUniqueId(req.GetUnique()); err != nil {
		return nil, err
	}

	obj, err := g.pubcase.FetchObj("", biz.WithRsaNoGen)
	if err != nil {
		return nil, err
	}

	_ = obj.Private
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func (g *AccountService) Login(_ context.Context, req *acctspb.LoginRequest) (*acctspb.LoginReply, error) {
	if err := g.validateUniqueId(req.GetUnique()); err != nil {
		return nil, err
	}

	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
