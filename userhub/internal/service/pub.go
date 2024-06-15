package service

import (
	"context"
	"fmt"
	"github.com/transerver/accounts/internal/biz"
	"github.com/transerver/pkg/errors"
	"github.com/transerver/protos/acctspb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"time"
)

type PubService struct {
	acctspb.UnimplementedRsaServiceServer

	usecase *biz.PubUsecase
}

func NewRsaService(usecase *biz.PubUsecase) *PubService {
	return &PubService{usecase: usecase}
}

func (g *PubService) RegisterGRPC(s *grpc.Server) {
	acctspb.RegisterRsaServiceServer(s, g)
}

func (g *PubService) PublicKey(ctx context.Context, req *acctspb.RsaRequest) (*acctspb.RsaReplay, error) {
	if req.G {
		uniqueId, err := g.Unique(ctx, nil)
		if err != nil {
			return nil, err
		}
		req.Unique = uniqueId.Value
	}

	err := req.Validate()
	if err != nil {
		return nil, errors.NewValidate(ctx, err)
	}

	if !req.G {
		if err = g.usecase.ValidateUniqueId(ctx, req.GetUnique()); err != nil {
			return nil, err
		}
	}

	requestId := fmt.Sprintf("%s:%s", req.GetAction(), req.GetUnique())
	obj, err := g.usecase.FetchRsaObj(ctx, requestId)
	if err != nil {
		return nil, err
	}

	return &acctspb.RsaReplay{
		Unique: req.Unique,
		Key:    obj.Public,
	}, nil
}

func (g *PubService) Unique(ctx context.Context, _ *emptypb.Empty) (*wrapperspb.StringValue, error) {
	uniqueId, err := g.usecase.FetchUniqueId(ctx, time.Minute*10)
	if err != nil {
		return nil, err
	}
	return &wrapperspb.StringValue{Value: uniqueId}, nil
}
