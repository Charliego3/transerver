package db

import (
	"context"
	nanoid "github.com/matoous/go-nanoid/v2"
	"github.com/stretchr/testify/require"
	"github.com/transerver/utils"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	account, err := testQueries.AccountCreate(context.Background(), AccountCreateParams{
		UserID:   nanoid.Must(),
		Username: "Charlie",
		Region:   "CN",
		Area:     "86",
		Phone:    utils.SQLString("15293441412"),
		Password: []byte("password"),
		PwdLevel: 5,
		Platform: "transerver",
	})
	require.NoError(t, err)

	t.Log(account)
}

func TestGetAccount(t *testing.T) {
	account, err := testQueries.AccountById(context.Background(), 1)
	require.NoError(t, err)
	require.NotEmpty(t, account)
}
