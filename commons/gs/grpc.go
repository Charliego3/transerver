package gs

import (
	"context"
	gm "github.com/grpc-ecosystem/go-grpc-middleware"
	ga "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	gz "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	gr "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	gc "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	gt "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	gp "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/transerver/commons/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server

	streamOpts []grpc.StreamServerInterceptor
	unaryOpts  []grpc.UnaryServerInterceptor
	serverOpts []grpc.ServerOption
	ctxTagOpts []gc.Option
	tracingOpt []gt.Option
	loggerOpts []gz.Option
	recoverOpt []gr.Option
	authFunc   ga.AuthFunc
}

func NewGRPCServer(logger *zap.Logger, services []service.Service, opts ...Option) (*Server, func()) {
	gs := &Server{}
	for _, opt := range opts {
		opt(gs)
	}

	gs.streamOpts = append(gs.streamOpts,
		gp.StreamServerInterceptor,
		gc.StreamServerInterceptor(gs.ctxTagOpts...),
		gt.StreamServerInterceptor(gs.tracingOpt...),
		gz.StreamServerInterceptor(logger, gs.loggerOpts...),
		gr.StreamServerInterceptor(gs.recoverOpt...),
	)

	gs.unaryOpts = append(gs.unaryOpts,
		gp.UnaryServerInterceptor,
		gc.UnaryServerInterceptor(gs.ctxTagOpts...),
		gt.UnaryServerInterceptor(gs.tracingOpt...),
		gz.UnaryServerInterceptor(logger, gs.loggerOpts...),
		gr.UnaryServerInterceptor(gs.recoverOpt...),
	)

	if gs.authFunc != nil {
		routes := make(map[string]struct{})
		for _, s := range services {
			rs, _ := s.Routers()
			for _, r := range rs {
				routes[r] = struct{}{}
			}
		}

		fn := ga.AuthFunc(func(ctx context.Context) (context.Context, error) {
			if method, ok := grpc.Method(ctx); !ok {
				return ctx, nil
			} else if _, ok := routes[method]; !ok {
				return ctx, nil
			}

			return gs.authFunc(ctx)
		})

		gs.streamOpts = append(gs.streamOpts, ga.StreamServerInterceptor(fn))
		gs.unaryOpts = append(gs.unaryOpts, ga.UnaryServerInterceptor(fn))
	}

	gs.serverOpts = append(gs.serverOpts,
		grpc.StreamInterceptor(gm.ChainStreamServer(gs.streamOpts...)),
		grpc.UnaryInterceptor(gm.ChainUnaryServer(gs.unaryOpts...)),
	)

	gs.Server = grpc.NewServer(gs.serverOpts...)
	for _, s := range services {
		s.RegisterGRPC(gs.Server)
	}
	return gs, gs.GracefulStop
}
