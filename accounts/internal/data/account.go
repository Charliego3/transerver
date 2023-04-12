package data

import (
	"context"
	"database/sql"
	"github.com/gookit/goutil/strutil"
	"github.com/transerver/accounts/internal/biz"
	db "github.com/transerver/accounts/internal/data/sqlc"
	"github.com/transerver/pkg1/errors"
	"github.com/transerver/utils"
)

type accountRepo struct{}

func NewAccountRepo() biz.AccountRepo {
	return &accountRepo{}
}

func (r *accountRepo) ExistsByPhone(ctx context.Context, phone sql.NullString) bool {
	exist, err := queries.AccountExistByPhone(ctx, phone)
	if err != nil {
		return false
	}
	return exist
}

func (r *accountRepo) ExistsByEmail(ctx context.Context, email sql.NullString) bool {
	exist, err := queries.AccountExistsByEmail(ctx, email)
	if err != nil {
		return false
	}
	return exist
}

func (r *accountRepo) Create(ctx context.Context, params db.AccountCreateParams) (db.Account, error) {
	return queries.AccountCreate(ctx, params)
}

func (r *accountRepo) ByUname(ctx context.Context, uname string) (db.Account, error) {
	var account db.Account
	if strutil.IsBlank(uname) {
		return account, errors.NewArgumentf(ctx, "用户名不能为空")
	}

	var err error
	if utils.IsEmail(uname) {
		account, err = queries.AccountByEmail(ctx, utils.SQLString(uname))
	} else {
		account, err = queries.AccountByPhone(ctx, utils.SQLString(uname))
	}
	if err == sql.ErrNoRows {
		return account, errors.NewArgumentf(ctx, "用户不存在")
	} else if err != nil {
		err = errors.NewInternal(ctx, "操作失败, 请重试")
	}
	return account, err
}
