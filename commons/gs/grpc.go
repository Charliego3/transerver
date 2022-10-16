package gs

import (
	"github.com/Charliego93/go-i18n/v2"
	"github.com/golang/glog"
	gm "github.com/grpc-ecosystem/go-grpc-middleware"
	ga "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	gz "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	gr "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	gc "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	gt "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	gp "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/transerver/commons/configs"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	*grpc.Server

	bootstrap  configs.IConfig
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

func NewGRPCServer(logger *zap.Logger, bootstrap configs.IConfig, services []Service, opts ...Option) (*Server, func()) {
	gs := &Server{bootstrap: bootstrap}
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
	l, err := net.Listen("tcp", s.bootstrap.Addr())
	if err != nil {
		return err
	}
	defer func() {
		s.GracefulStop()
		if err := l.Close(); err != nil {
			glog.Errorf("Failed to close %s %s: %v", "tcp", s.bootstrap.Addr(), err)
		}
	}()
	return s.Serve(l)
}
