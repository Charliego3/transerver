package gs

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/transerver/commons/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server

	streamOpts []grpc.StreamServerInterceptor
	unaryOpts  []grpc.UnaryServerInterceptor
	serverOpts []grpc.ServerOption
	ctxTagOpts []grpc_ctxtags.Option
	tracingOpt []grpc_opentracing.Option
	loggerOpts []grpc_zap.Option
	recoverOpt []grpc_recovery.Option
	authFunc   grpc_auth.AuthFunc
}

func NewGRPCServer(logger *zap.Logger, services []service.Service, opts ...Option) (*Server, func()) {
	gs := &Server{}
	for _, opt := range opts {
		opt(gs)
	}

	gs.streamOpts = append(gs.streamOpts,
		grpc_prometheus.StreamServerInterceptor,
		grpc_ctxtags.StreamServerInterceptor(gs.ctxTagOpts...),
		grpc_opentracing.StreamServerInterceptor(gs.tracingOpt...),
		grpc_zap.StreamServerInterceptor(logger, gs.loggerOpts...),
		grpc_recovery.StreamServerInterceptor(gs.recoverOpt...),
	)

	gs.unaryOpts = append(gs.unaryOpts,
		grpc_prometheus.UnaryServerInterceptor,
		grpc_ctxtags.UnaryServerInterceptor(gs.ctxTagOpts...),
		grpc_opentracing.UnaryServerInterceptor(gs.tracingOpt...),
		grpc_zap.UnaryServerInterceptor(logger, gs.loggerOpts...),
		grpc_recovery.UnaryServerInterceptor(gs.recoverOpt...),
	)

	if gs.authFunc != nil {
		var prvRoutes []string
		for _, s := range services {
			routes, _ := s.Routers()
			prvRoutes = append(prvRoutes, routes...)
		}

		fn := grpc_auth.AuthFunc(func(ctx context.Context) (context.Context, error) {
			if method, ok := grpc.Method(ctx); ok {
				for _, route := range prvRoutes {
					if route != method {
						continue
					}

					ctx2, err := gs.authFunc(ctx)
					if err != nil {
						return nil, err
					}
					return ctx2, nil
				}
			}
			return ctx, nil
		})

		gs.streamOpts = append(gs.streamOpts, grpc_auth.StreamServerInterceptor(fn))
		gs.unaryOpts = append(gs.unaryOpts, grpc_auth.UnaryServerInterceptor(fn))
	}

	gs.serverOpts = append(gs.serverOpts,
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(gs.streamOpts...)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(gs.unaryOpts...)),
	)

	gs.Server = grpc.NewServer(gs.serverOpts...)
	for _, s := range services {
		s.RegisterGRPC(gs.Server)
	}
	return gs, gs.GracefulStop
}
