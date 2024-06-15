package service

import (
	"context"
	"github.com/gookit/goutil/testutil/assert"
	"github.com/transerver/protos/acctspb"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
)

func TestUnique(t *testing.T) {
	ctx := context.Background()
	client := acctspb.NewRsaServiceClient(conn)
	value, err := client.Unique(ctx, &emptypb.Empty{})
	assert.NoErr(t, err, "fetch uniqueId failed")
	assert.NotEmpty(t, value.Value, "returns uniqueId is empty")
}

func TestRegister(t *testing.T) {

}
