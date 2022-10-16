package gw

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/transerver/commons/configs"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"net/http"
)

type HandlerFunc func(*mux.Router)

type Server struct {
	gwmux http.Handler
	ctx   context.Context

	bootstrap  configs.IConfig
	logger     *zap.SugaredLogger
	muxOpts    []runtime.ServeMuxOption
	handler    HandlerFunc
	authFunc   func(*http.Request) error
	dialers    []Dialer
	middleware []mux.MiddlewareFunc

	conns []*grpc.ClientConn
}

func (gs *Server) Check(_ context.Context, _ *grpc_health_v1.HealthCheckRequest, _ ...grpc.CallOption) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (gs *Server) Watch(_ context.Context, _ *grpc_health_v1.HealthCheckRequest, _ ...grpc.CallOption) (grpc_health_v1.Health_WatchClient, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func NewGatewayServer(logger *zap.Logger, bootstrap configs.IConfig, opts ...Option) (*Server, error) {
	gs := &Server{logger: logger.Sugar(), bootstrap: bootstrap, ctx: context.Background()}
	gs.muxOpts = []runtime.ServeMuxOption{
		runtime.WithMarshalerOption("application/json", NewJSONMarshaller()),
		runtime.WithMarshalerOption("application/json+pretty", NewJSONMarshaller(true)),
		runtime.WithErrorHandler(DefaultErrorHandler),
		runtime.WithRoutingErrorHandler(DefaultRoutingErrorHandler),
		runtime.WithHealthzEndpoint(gs),
	}

	for _, opt := range opts {
		opt(gs)
	}
	return gs, nil
}

func (gs *Server) Run() error {
	ctx, cancel := context.WithCancel(gs.ctx)
	defer cancel()

	err := gs.newgw(ctx)
	if err != nil {
		return err
	}

	router := mux.NewRouter()
	router.Use(gs.middleware...)

	if gs.handler != nil {
		gs.handler(router)
	}

	router.PathPrefix("/").Handler(gs.gwmux)

	s := &http.Server{
		Addr:    gs.bootstrap.Addr(),
		Handler: router,
	}

	defer func() {
		for _, conn := range gs.conns {
			if err := conn.Close(); err != nil {
				gs.logger.Errorf("Failed to close a client connection to the gRPC[%s] server: %v", conn.Target(), err)
			}
		}

		if err := s.Shutdown(ctx); err != nil {
			gs.logger.Errorf("Failed to shutdown http[%s] server: %v", gs.bootstrap.Addr(), err)
		}
	}()

	gs.logger.Infof("Starting listening at %s", gs.bootstrap.Addr())
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		gs.logger.Errorf("Failed to listen and serve: %v", err)
		return err
	}
	return nil
}

func (gs *Server) newgw(ctx context.Context) error {
	gwmux := runtime.NewServeMux(gs.muxOpts...)
	for _, dialer := range gs.dialers {
		conn, err := gs.dial(ctx, dialer.Target, dialer.Opts...)
		if err != nil {
			return err
		}

		for _, fn := range dialer.Servers {
			err := fn(ctx, gwmux, conn)
			if err != nil {
				return err
			}
		}
	}
	gs.gwmux = gwmux
	return nil
}

func (gs *Server) dial(ctx context.Context, target string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	if len(opts) == 0 {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	conn, err := grpc.DialContext(ctx, target, opts...)
	if err != nil {
		return nil, err
	}

	gs.conns = append(gs.conns, conn)
	return conn, nil
}

func (gs *Server) auth(
	mux *runtime.ServeMux,
	routers map[string]struct{},
	w http.ResponseWriter,
	r *http.Request,
) bool {
	if gs.authFunc == nil {
		return false
	}

	if _, ok := routers[r.URL.Path]; ok {
		return false
	}

	err := gs.authFunc(r)
	if err == nil {
		return false
	}

	_, outbound := runtime.MarshalerForRequest(mux, r)
	buf, err := outbound.Marshal(err)
	if err != nil {
		buf = []byte(fallback)
	}
	_, _ = w.Write(buf)
	return true
}
