package service

import (
	"context"
	"fmt"
	"github.com/transerver/accounts/internal/biz"
	"github.com/transerver/commons/errors"
	"github.com/transerver/protos/acctspb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccountService struct {
	acctspb.UnimplementedAccountServiceServer

	usecase *biz.AccountUsecase
	pubcase *biz.PubUsecase
}

func NewAccountService(usecase *biz.AccountUsecase, rsa *biz.PubUsecase) *AccountService {
	return &AccountService{usecase: usecase, pubcase: rsa}
}

func (g *AccountService) RegisterGRPC(s *grpc.Server) {
	acctspb.RegisterAccountServiceServer(s, g)
}

func (g *AccountService) Register(ctx context.Context, req *acctspb.RegisterRequest) (*acctspb.RegisterReply, error) {
	obj, err := getRsaObj(ctx, g, req)
	err = g.usecase.Register(ctx, req, obj)
	if err != nil {
		return nil, err
	}
	return &acctspb.RegisterReply{}, nil
}

func (g *AccountService) Login(ctx context.Context, req *acctspb.LoginRequest) (*acctspb.LoginReply, error) {
	obj, err := getRsaObj(ctx, g, req)
	if err != nil {
		return nil, err
	}

	g.usecase.Login(ctx, req, obj)
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func getRsaObj[T interface {
	Validate() error
	GetUnique() string
}](ctx context.Context, g *AccountService, req T) (*biz.RsaObj, error) {
	err := req.Validate()
	if err != nil {
		return nil, errors.NewArgument(ctx, err)
	}

	if err = g.pubcase.ValidateUniqueId(ctx, req.GetUnique()); err != nil {
		return nil, err
	}

	return g.pubcase.FetchRsaObj(ctx, fmt.Sprintf("register:%s", req.GetUnique()), biz.WithRsaNoGen)
}
