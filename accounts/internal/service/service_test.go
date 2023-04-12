package service

import (
	"context"
	"github.com/transerver/accounts/internal/conf"
	"github.com/transerver/pkg1/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
	"time"
)

var (
	conn *grpc.ClientConn
)

func TestMain(m *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	clientConn, err := grpc.DialContext(
		ctx, conf.Bootstrap.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Sugar().Fatalf("dial grpc server error: %v", err)
	}
	conn = clientConn
	m.Run()
}
