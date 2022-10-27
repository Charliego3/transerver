package biz

import (
	"context"
	"github.com/Charliego93/go-i18n/v2"
	"github.com/gookit/goutil/strutil"
	nanoid "github.com/matoous/go-nanoid/v2"
	"github.com/nyaruka/phonenumbers"
	"github.com/transerver/accounts/internal/ent"
	"github.com/transerver/accounts/internal/ent/region"
	"github.com/transerver/commons/errors"
	"github.com/transerver/commons/types/enums"
	"github.com/transerver/protos/acctspb"
	"github.com/transerver/utils"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"unicode"
)

const (
	minPasswordLength = 8
	maxPasswordLength = 32
)

type AccountRepo interface {
	FindById(context.Context, int64, ...string) (*ent.Account, error)
	Save(context.Context, *ent.Account) (*ent.Account, error)
	CheckPhoneExists(context.Context, string) bool
	CheckEmailExists(context.Context, string) bool
}

type AccountUsecase struct {
	repo       AccountRepo
	regionRepo RegionRepo
}

func NewAccountUsecase(repo AccountRepo, regionRepo RegionRepo) *AccountUsecase {
	return &AccountUsecase{repo: repo, regionRepo: regionRepo}
}

func (g *AccountUsecase) Register(ctx context.Context, req *acctspb.RegisterRequest, obj *RsaObj) error {
	if utils.Blanks(req.Phone, req.Email) {
		return errors.NewArgumentf(ctx, "手机和邮箱不能同时为空")
	}

	var reg *ent.Region
	if strutil.IsNotBlank(req.Phone) {
		req.Region = strings.ToUpper(req.Region)
		var err error
		reg, err = g.regionRepo.FindByCode(context.Background(), req.Region, region.FieldCode, region.FieldArea)
		if err != nil {
			return errors.NewArgumentf(ctx, &i18n.Localized{
				MessageID:    "RegionNotFound",
				TemplateData: req.Region,
			})
		}

		number, err := phonenumbers.Parse(req.Phone, reg.Code)
		if err != nil {
			return errors.NewArgumentf(ctx, "手机号码和地区不匹配")
		}

		if !phonenumbers.IsValidNumberForRegion(number, req.Region) {
			return errors.NewArgumentf(ctx, "手机号码和地区不匹配")
		}

		if g.repo.CheckPhoneExists(ctx, req.Phone) {
			return errors.NewArgumentf(ctx, "手机号已经存在")
		}
	} else if g.repo.CheckEmailExists(ctx, req.Email) {
		return errors.NewArgumentf(ctx, "邮箱已经存在")
	}

	password := utils.B64Decode(req.Password)
	pwd, err := obj.Decrypt(ctx, password)
	if err != nil {
		return errors.NewInternal(ctx, "请求失败, 请刷新页面重试!")
	}

	pwdBuf, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return errors.NewInternal(ctx, "注册失败, 请尝试修改密码")
	}

	pwdLevel, err := g.passwordLevel(ctx, pwd)
	if err != nil {
		return err
	}

	account := &ent.Account{
		UserID:   nanoid.Must(),
		Username: req.Uname,
		Region:   req.Region,
		Area:     reg.Area,
		Phone:    req.Phone,
		Email:    req.Email,
		Password: pwdBuf,
		PwdLevel: pwdLevel,
		Platform: "p",
		State:    enums.UserUnverified,
	}

	account, err = g.repo.Save(ctx, account)
	_ = account
	return err
}

func (g *AccountUsecase) Login(ctx context.Context, req *acctspb.LoginRequest, obj *RsaObj) {

}

// passwordLevel returns level
// if it has a special rune level++ and special count > 5 then level++
// if it has a number level++ and number count > 5 then level++
// if it has an upperCase character level++ and upperCase character count > 5 then level++
// if it has a lower character level++
// if it has more than 5 space count then level++
func (g *AccountUsecase) passwordLevel(ctx context.Context, pwd []byte) (level uint8, err error) {
	password := []rune(utils.String(pwd))
	pwdLength := len(password)
	if pwdLength < minPasswordLength {
		return 0, errors.NewArgumentf(ctx, "密码强度过低, 不得少于?字符", minPasswordLength)
	} else if pwdLength > maxPasswordLength {
		return 0, errors.NewArgumentf(ctx, "密码过长，最长不超过?字符", maxPasswordLength)
	}

	var sc, nc, uc, lc, ec int // specialCount, numberCount, upperCount, lowerCount, spaceCount
	for _, r := range password {
		if unicode.IsControl(r) {
			return 0, errors.NewArgumentf(ctx, "密码包含非法字符")
		}

		if unicode.IsUpper(r) {
			uc++
		}
		if unicode.IsLower(r) {
			lc++
		}
		if unicode.IsSymbol(r) || unicode.IsPunct(r) || unicode.IsLetter(r) {
			sc++
		}
		if unicode.IsSpace(r) {
			ec++
		}
		if unicode.IsNumber(r) {
			nc++
		}
	}

	if nc == 0 {
		return 0, errors.NewArgumentf(ctx, "密码必须包含数字")
	} else {
		level++
	}
	if uc == 0 {
		return 0, errors.NewArgumentf(ctx, "密码必须包含大写字母")
	} else {
		level++
	}
	if lc == 0 {
		return 0, errors.NewArgumentf(ctx, "密码必须包含小写字母")
	} else {
		level++
	}

	if sc > 0 {
		level++
		if sc > 5 {
			level++
		}
	}
	if nc > 5 {
		level++
	}
	if uc > 5 {
		level++
	}
	if ec > 0 {
		level++
	}
	return level, nil
}
