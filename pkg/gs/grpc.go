package gs

import (
	"context"
	"fmt"
	"github.com/Charliego93/go-i18n/v2"
	gm "github.com/grpc-ecosystem/go-grpc-middleware"
	ga "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	gz "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	gr "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	gc "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	gt "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	gp "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/transerver/pkg/app"
	"github.com/transerver/pkg/configs"
	"github.com/transerver/pkg/es"
	"github.com/transerver/pkg/logger"
	"github.com/transerver/utils"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	*grpc.Server
	em endpoints.Manager

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
	manager, err := endpoints.NewManager(es.C().Client, "/"+app.Name)
	if err != nil {
		logger.Sugar().Fatal("endpoint manager initialize error", err)
	}

	gs := &Server{em: manager}
	for _, opt := range opts {
		opt(gs)
	}

	gs.streamOpts = append(gs.streamOpts,
		gp.StreamServerInterceptor,
		gc.StreamServerInterceptor(gs.ctxTagOpts...),
		gt.StreamServerInterceptor(gs.tracingOpt...),
		gz.StreamServerInterceptor(logger.Standard(), gs.loggerOpts...),
		gr.StreamServerInterceptor(gs.recoverOpt...),
		StreamI18nInterceptor,
	)

	gs.unaryOpts = append(gs.unaryOpts,
		gp.UnaryServerInterceptor,
		gc.UnaryServerInterceptor(gs.ctxTagOpts...),
		gt.UnaryServerInterceptor(gs.tracingOpt...),
		gz.UnaryServerInterceptor(logger.Standard(), gs.loggerOpts...),
		gr.UnaryServerInterceptor(gs.recoverOpt...),
		UnaryI18nInterceptor,
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

	if len(gs.i18nOpts) == 0 {
		gs.i18nOpts = []i18n.Option{
			i18n.NewLoaderWithPath("i18n"),
		}
	}

	gs.i18nOpts = append([]i18n.Option{i18n.WithDefaultLanguage(utils.DefaultLanguage)}, gs.i18nOpts...)
	i18n.Initialize(gs.i18nOpts...)
	return gs, gs.GracefulStop
}

func (s *Server) Run() error {
	l, err := net.Listen(configs.Bootstrap.Root().Network, configs.Bootstrap.Root().Address)
	if err != nil {
		return err
	}

	target := fmt.Sprintf("/%s/%s", app.Name, configs.Bootstrap.Root().Address)

	defer func() {
		if err := l.Close(); err != nil {
			logger.Sugar().Errorf("Failed to close %s %s: %v",
				configs.Bootstrap.Root().Network, configs.Bootstrap.Root().Address, err)
		}
		s.GracefulStop()
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*3)
		err := s.em.DeleteEndpoint(ctx, target)
		if err != nil {
			logger.Sugar().Errorf("Delete grpc endpoint error", err)
		}
		cancelFunc()
	}()
	logger.Sugar().Infof(`Starting listening at "%s:%s"`,
		configs.Bootstrap.Root().Network, configs.Bootstrap.Root().Address)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = s.em.AddEndpoint(ctx, target, endpoints.Endpoint{
		Addr: configs.Bootstrap.Root().Address,
	})
	if err != nil {
		logger.Sugar().Errorf("Add endpoint error: %v", err)
		return err
	}

	go func() {
		err := s.Serve(l)
		if err != nil {
			logger.Sugar().Fatalf("Server serve error: %v", err)
		}
	}()

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGABRT, syscall.SIGIOT, syscall.SIGTERM)
	<-ch

	return nil
}
