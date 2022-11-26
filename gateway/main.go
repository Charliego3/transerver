package main

import (
	"github.com/Charliego93/go-i18n/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/transerver/pkg/gw"
	"github.com/transerver/pkg/logger"
	"github.com/transerver/pkg/resolver"
	"google.golang.org/grpc"
	//_ "google.golang.org/grpc/balancer/grpclb"
	//_ "google.golang.org/grpc/balancer/rls"
	//_ "google.golang.org/grpc/balancer/roundrobin"
	//_ "google.golang.org/grpc/balancer/weightedroundrobin"
	//_ "google.golang.org/grpc/balancer/weightedtarget"
	"google.golang.org/grpc/credentials/insecure"
	"path/filepath"
)

func main() {
	opts := []gw.Option{
		gw.WithServeMuxOpts(runtime.WithIncomingHeaderMatcher(func(key string) (string, bool) {
			switch key {
			case "Accept-Language":
				return key, true
			default:
				return "", false
			}
		})),
	}

	if len(Bootstrap.Services) == 0 {
		logger.Sugar().Fatal("there is no registered service.")
	}

	for _, service := range Bootstrap.Services {
		if ds, ok := dialers[filepath.Base(service.Target)]; ok {
			opts = append(opts, gw.WithDialer(
				gw.NewDialer(service.Target, ds,
					grpc.WithTransportCredentials(insecure.NewCredentials()),
					grpc.WithResolvers(resolver.NewBuilder(service.Config)),
				),
			))
			continue
		}

		logger.Sugar().Fatalf("doesn't declare dialer for: %q", service.Target)
	}

	i18n.Initialize()
	app, err := gw.NewGatewayServer(opts...)
	if err != nil {
		logger.Sugar().Fatal("create gateway fail", err)
	}

	if err := app.Run(); err != nil {
		logger.Sugar().Fatal("gateway running error", err)
	}
}
