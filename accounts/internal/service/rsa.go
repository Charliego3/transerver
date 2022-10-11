package service

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/transerver/accounts/internal/biz"
	"github.com/transerver/protos/acctspb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type RsaService struct {
	acctspb.UnimplementedRsaServiceServer

	usecase *biz.RsaUsecase
	logger  *zap.Logger
}

func NewRsaService(usecase *biz.RsaUsecase, logger *zap.Logger) *RsaService {
	return &RsaService{usecase: usecase, logger: logger}
}

func (g *RsaService) RegisterGRPC(s *grpc.Server) {
	acctspb.RegisterRsaServiceServer(s, g)
}

func (g *RsaService) RegisterHTTP(s *runtime.ServeMux) error {
	return acctspb.RegisterRsaServiceHandlerServer(context.Background(), s, g)
}

func (g *RsaService) Routers() ([]string, []string) {
	return nil, nil
}

func (g *RsaService) PublicKey(context.Context, *emptypb.Empty) (*wrapperspb.BytesValue, error) {
	obj, err := g.usecase.FetchObj(":testRequestId") // TODO: real requestId
	if err != nil {
		return nil, err
	}

	return &wrapperspb.BytesValue{Value: obj.Public}, nil
}
