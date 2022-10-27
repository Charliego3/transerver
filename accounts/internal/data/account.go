package data

import (
	"context"
	"github.com/transerver/accounts/internal/biz"
	"github.com/transerver/accounts/internal/ent"
	"github.com/transerver/accounts/internal/ent/account"
)

var _ biz.AccountRepo = (*accountRepo)(nil)

type accountRepo struct {
	data *Data
}

func NewAccountRepo(data *Data) biz.AccountRepo {
	return &accountRepo{data: data}
}

func (g *accountRepo) FindById(ctx context.Context, uid int64, fields ...string) (*ent.Account, error) {
	return g.data.ent.Account.Query().Select(fields...).Where(account.ID(uid)).First(ctx)
}

func (g *accountRepo) CheckPhoneExists(ctx context.Context, phone string) bool {
	exist, err := g.data.ent.Account.Query().Where(account.Phone(phone)).Exist(ctx)
	if err != nil {
		return true
	}
	return exist
}

func (g *accountRepo) CheckEmailExists(ctx context.Context, email string) bool {
	exist, err := g.data.ent.Account.Query().Where(account.Email(email)).Exist(ctx)
	if err != nil {
		return true
	}
	return exist
}

func (g *accountRepo) Save(ctx context.Context, a *ent.Account) (*ent.Account, error) {
	return g.data.ent.Account.Create().
		SetUserID(a.UserID).
		SetUsername(a.Username).
		SetRegion(a.Region).
		SetArea(a.Area).
		SetPhone(a.Phone).
		SetEmail(a.Email).
		SetAvatar(a.Avatar).
		SetPassword(a.Password).
		SetPwdLevel(a.PwdLevel).
		SetPlatform(a.Platform).
		SetCreateAt(a.CreateAt).
		Save(ctx)
}
