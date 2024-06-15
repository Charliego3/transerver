package service

import (
	"context"
	"github.com/stretchr/objx"
	"github.com/transerver/accounts/internal/biz"
	"github.com/transerver/protos/acctspb"
	"github.com/transerver/utils"
	"golang.org/x/text/language"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RegionService struct {
	acctspb.UnimplementedRegionServiceServer

	usecase *biz.RegionUsecase
}

func NewRegionService(usecase *biz.RegionUsecase) *RegionService {
	return &RegionService{usecase: usecase}
}

func (g *RegionService) RegisterGRPC(s *grpc.Server) {
	acctspb.RegisterRegionServiceServer(s, g)
}

func (g *RegionService) Regions(ctx context.Context, _ *emptypb.Empty) (*acctspb.RegionReply, error) {
	regions, err := g.usecase.Regions(ctx)
	if err != nil {
		return nil, err
	}
	lang := utils.Language(ctx)
	var reply []*acctspb.RegionReply_Region
	for _, r := range regions {
		region := &acctspb.RegionReply_Region{
			Code: r.Code,
			Area: r.Area,
			Img:  r.Img,
		}

		obj := objx.MustFromJSON(string(r.Name))
		switch lang {
		case language.English:
			region.Name = obj.Get("en").String()
		default:
			region.Name = obj.Get("zh").String()
		}
		reply = append(reply, region)
	}
	return &acctspb.RegionReply{Regions: reply}, nil
}
