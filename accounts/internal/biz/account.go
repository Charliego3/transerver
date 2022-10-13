package biz

import (
	"context"
	"github.com/gookit/goutil/strutil"
	"github.com/transerver/accounts/internal/ent"
	"github.com/transerver/protos/acctspb"
	"github.com/transerver/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type AccountRepo interface {
	FindById(ctx context.Context, id int64) (*ent.Account, error)
}

type AccountUsecase struct {
	repo       AccountRepo
	regionRepo RegionRepo
	logger     *zap.Logger
}

type AccountHelper struct {
	repo   AccountRepo
	logger *zap.Logger
	err    error
}

func NewAccountUsecase(repo AccountRepo, regionRepo RegionRepo, logger *zap.Logger) *AccountUsecase {
	return &AccountUsecase{repo: repo, regionRepo: regionRepo, logger: logger}
}

func (g *AccountUsecase) Helper() *AccountHelper {
	return &AccountHelper{repo: g.repo, logger: g.logger}
}

func (h *AccountHelper) Err() error {
	return h.err
}

func (g *AccountUsecase) Register(req *acctspb.RegisterRequest, obj *RsaObj) error {
	if utils.Blanks(req.Phone, req.Email) {
		return status.Error(codes.InvalidArgument, "手机和邮箱不能同时为空")
	}

	if strutil.IsNotBlank(req.Phone) {
		req.Region = strings.ToUpper(req.Region)
		region, err := g.regionRepo.FindByCode(context.Background(), req.Region)
		if err != nil {
			return utils.ErrorArgumentf("Not found region with: [%s]", req.Region)
		}

		_ = region.Name.En
	} else {

	}
	return nil
}

func (g *AccountUsecase) Login(req *acctspb.LoginRequest, obj *RsaObj) {

}
