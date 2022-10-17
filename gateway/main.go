package main

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/transerver/commons/gw"
	"github.com/transerver/commons/logger"
	"github.com/transerver/protos/acctspb"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/balancer/grpclb"
	_ "google.golang.org/grpc/balancer/rls"
	_ "google.golang.org/grpc/balancer/roundrobin"
	_ "google.golang.org/grpc/balancer/weightedroundrobin"
	_ "google.golang.org/grpc/balancer/weightedtarget"
	"google.golang.org/grpc/credentials/insecure"
)

var accountDialer = []gw.DialerFunc{
	acctspb.RegisterAccountServiceHandler,
	acctspb.RegisterRsaServiceHandler,
	acctspb.RegisterRegionServiceHandler,
}

func main() {
	app, err := gw.NewGatewayServer(
		gw.WithServeMuxOpts(runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
			switch key {
			case "Accept-Language":
				return key, true
			default:
				return "", false
			}
		})),
		gw.WithDialer(
			gw.NewDialer(
				"discovery:///accounts",
				accountDialer,
				grpc.WithTransportCredentials(insecure.NewCredentials()),
				grpc.WithResolvers(gw.NewResolver(":9091", ":9092")),
			),
		))
	if err != nil {
		logger.Sugar().Fatal("create gateway fail", err)
	}

	if err := app.Run(); err != nil {
		logger.Sugar().Fatal("gateway running error", err)
	}
}
