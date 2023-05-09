package grpcx

import (
	"context"
	"testing"

	"github.com/charmbracelet/log"
	"github.com/stretchr/testify/require"
	"github.com/transerver/protos/acctspb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type AccountService struct {
	acctspb.UnimplementedAccountServiceServer
}

func NewAccountService() *AccountService {
	return &AccountService{}
}

func (g *AccountService) ServiceDesc() *grpc.ServiceDesc {
	return &acctspb.AccountService_ServiceDesc
}

func (g *AccountService) Register(
	ctx context.Context,
	req *acctspb.RegisterRequest,
) (*acctspb.RegisterReply, error) {
	log.Info("Register request", "body", req)
	return &acctspb.RegisterReply{}, nil
}

func (g *AccountService) Login(
	ctx context.Context,
	req *acctspb.LoginRequest,
) (*acctspb.LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func TestService(t *testing.T) {
	app := NewServer(WithAddr("tcp", "0.0.0.0:8081"))
	app.RegisterService(NewAccountService())
	err := app.Run()
	require.NoError(t, err)
}

func TestClient(t *testing.T) {
	conn, err := grpc.Dial("0.0.0.0:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	defer conn.Close()

	c := acctspb.NewAccountServiceClient(conn)
	reply, err := c.Register(context.Background(), &acctspb.RegisterRequest{Uname: "Charlie", Email: "charlie.go.3@outlook.com"})
	require.NoError(t, err)

	t.Logf("Response: %+v", reply)
}
