package biz

import (
	"context"
	"github.com/Charliego93/go-i18n/v2"
	"github.com/gookit/goutil/strutil"
	"github.com/transerver/accounts/internal/ent"
	"github.com/transerver/commons/errors"
	"github.com/transerver/protos/acctspb"
	"github.com/transerver/utils"
	"strings"
)

type AccountRepo interface {
	FindById(ctx context.Context, id int64) (*ent.Account, error)
}

type AccountUsecase struct {
	repo       AccountRepo
	regionRepo RegionRepo
}

type AccountHelper struct {
	repo AccountRepo
	err  error
}

func NewAccountUsecase(repo AccountRepo, regionRepo RegionRepo) *AccountUsecase {
	return &AccountUsecase{repo: repo, regionRepo: regionRepo}
}

func (g *AccountUsecase) Helper() *AccountHelper {
	return &AccountHelper{repo: g.repo}
}

func (h *AccountHelper) Err() error {
	return h.err
}

func (g *AccountUsecase) Register(ctx context.Context, req *acctspb.RegisterRequest, obj *RsaObj) error {
	if utils.Blanks(req.Phone, req.Email) {
		return errors.NewArgumentf(ctx, "手机和邮箱不能同时为空")
	}

	if strutil.IsNotBlank(req.Phone) {
		req.Region = strings.ToUpper(req.Region)
		region, err := g.regionRepo.FindByCode(context.Background(), req.Region)
		if err != nil {
			return errors.NewArgumentf(ctx, &i18n.Localized{
				MessageID:    "RegionNotFound",
				TemplateData: req.Region,
			})
		}

		_ = region.Name.En
	} else {

	}
	return nil
}

func (g *AccountUsecase) Login(req *acctspb.LoginRequest, obj *RsaObj) {

}
