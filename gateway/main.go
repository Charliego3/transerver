package main

import (
	"github.com/google/wire"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/transerver/commons/configs"
	"github.com/transerver/commons/gw"
	"github.com/transerver/protos/acctspb"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/balancer/grpclb"
	_ "google.golang.org/grpc/balancer/rls"
	_ "google.golang.org/grpc/balancer/roundrobin"
	_ "google.golang.org/grpc/balancer/weightedroundrobin"
	_ "google.golang.org/grpc/balancer/weightedtarget"
	"google.golang.org/grpc/credentials/insecure"
	"io"
)

var providerSet = wire.NewSet(
	NewLoggerWriter,
	NewBootstrap,
	NewCfgOpts,
	NewGWOpts,
)

var accountDialer = []gw.DialerFunc{
	acctspb.RegisterAccountServiceHandler,
	acctspb.RegisterRsaServiceHandler,
	acctspb.RegisterRegionServiceHandler,
}

func main() {
	app, cleanup, err := wireApp()
	if err != nil {
		panic(err)
	}

	defer cleanup()
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func NewGWOpts() []gw.Option {
	return []gw.Option{
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
		),
	}
}

func NewLoggerWriter() io.Writer {
	return io.Discard
}

func NewBootstrap() any {
	return &configs.Bootstrap{}
}

func NewCfgOpts() []configs.Option {
	return []configs.Option{
		configs.WithPath("config.yaml"),
	}
}
