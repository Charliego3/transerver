package gs

import (
	"github.com/Charliego93/go-i18n/v2"
	gm "github.com/grpc-ecosystem/go-grpc-middleware"
	ga "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	gz "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	gr "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	gc "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	gt "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	gp "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/transerver/commons/configs"
	"github.com/transerver/commons/logger"
	"google.golang.org/grpc"
	"net"
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
	i18nOpts   []i18n.Option
	authFunc   ga.AuthFunc
}

func NewGRPCServer(services []Service, opts ...Option) (*Server, func()) {
	gs := &Server{}
	for _, opt := range opts {
		opt(gs)
	}

	gs.streamOpts = append(gs.streamOpts,
		gp.StreamServerInterceptor,
		gc.StreamServerInterceptor(gs.ctxTagOpts...),
		gt.StreamServerInterceptor(gs.tracingOpt...),
		gz.StreamServerInterceptor(logger.Standard(), gs.loggerOpts...),
		gr.StreamServerInterceptor(gs.recoverOpt...),
	)

	gs.unaryOpts = append(gs.unaryOpts,
		gp.UnaryServerInterceptor,
		gc.UnaryServerInterceptor(gs.ctxTagOpts...),
		gt.UnaryServerInterceptor(gs.tracingOpt...),
		gz.UnaryServerInterceptor(logger.Standard(), gs.loggerOpts...),
		gr.UnaryServerInterceptor(gs.recoverOpt...),
	)

	if gs.authFunc != nil {
		gs.streamOpts = append(gs.streamOpts, ga.StreamServerInterceptor(gs.authFunc))
		gs.unaryOpts = append(gs.unaryOpts, ga.UnaryServerInterceptor(gs.authFunc))
	}

	gs.serverOpts = append(gs.serverOpts,
		grpc.StreamInterceptor(gm.ChainStreamServer(gs.streamOpts...)),
		grpc.UnaryInterceptor(gm.ChainUnaryServer(gs.unaryOpts...)),
	)

	gs.Server = grpc.NewServer(gs.serverOpts...)
	for _, s := range services {
		s.RegisterGRPC(gs.Server)
	}

	i18n.Initialize(gs.i18nOpts...)
	return gs, gs.GracefulStop
}

func (s *Server) Run() error {
	l, err := net.Listen(configs.Bootstrap.Root().Network, configs.Bootstrap.Root().Address)
	if err != nil {
		return err
	}
	defer func() {
		s.GracefulStop()
		if err := l.Close(); err != nil {
			logger.Sugar().Errorf("Failed to close %s %s: %v",
				configs.Bootstrap.Root().Network, configs.Bootstrap.Root().Address, err)
		}
	}()
	logger.Sugar().Infof(`Starting listening at "%s:%s"`,
		configs.Bootstrap.Root().Network, configs.Bootstrap.Root().Address)
	return s.Serve(l)
}
