package biz

import (
	"context"
	"database/sql"
	"github.com/Charliego93/go-i18n/v2"
	"github.com/gookit/goutil/strutil"
	nanoid "github.com/matoous/go-nanoid/v2"
	"github.com/nyaruka/phonenumbers"
	db "github.com/transerver/accounts/internal/db/sqlc"
	"github.com/transerver/pkg/errors"
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

type AccountUsecase struct{}

func NewAccountUsecase() *AccountUsecase {
	return &AccountUsecase{}
}

func (g *AccountUsecase) Register(ctx context.Context, req *acctspb.RegisterRequest, obj *RsaObj) error {
	if utils.Blanks(req.Phone, req.Email) {
		return errors.NewArgumentf(ctx, "手机和邮箱不能同时为空")
	}

	var reg db.Region
	var phone, email sql.NullString
	if strutil.IsNotBlank(req.Phone) {
		req.Region = strings.ToUpper(req.Region)
		var err error
		reg, err = db.Query().RegionByCode(context.Background(), req.Region)
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

		phone = utils.SQLString(req.Phone)
		exists, err := db.Query().AccountExistByPhone(ctx, phone)
		if err != nil && exists {
			return errors.NewArgumentf(ctx, "手机号已经存在")
		}
	} else {
		email = utils.SQLString(req.Email)
		exits, err := db.Query().AccountExistsByEmail(ctx, email)
		if err != nil && exits {
			return errors.NewArgumentf(ctx, "邮箱已经存在")
		}
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

	account, err := db.Query().AccountCreate(ctx, db.AccountCreateParams{
		UserID:   nanoid.Must(),
		Username: req.Uname,
		Region:   req.Region,
		Area:     reg.Area,
		Phone:    phone,
		Email:    email,
		Password: pwdBuf,
		PwdLevel: int16(pwdLevel),
		Platform: "p",
	})
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
