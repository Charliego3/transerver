package hs

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/transerver/commons/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"net/http"
)

type Server struct {
	*http.Server

	logger   *zap.SugaredLogger
	muxOpts  []runtime.ServeMuxOption
	handlers []Handler
	authFunc func(*http.Request) error
}

func (s *Server) Check(_ context.Context, _ *grpc_health_v1.HealthCheckRequest, _ ...grpc.CallOption) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (s *Server) Watch(_ context.Context, _ *grpc_health_v1.HealthCheckRequest, _ ...grpc.CallOption) (grpc_health_v1.Health_WatchClient, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func NewHTTPServer(logger *zap.Logger, services []service.Service, opts ...Option) (*Server, error) {
	hs := &Server{logger: logger.Sugar()}
	hs.muxOpts = []runtime.ServeMuxOption{
		runtime.WithMarshalerOption("application/json", NewJSONMarshaller()),
		runtime.WithMarshalerOption("application/json+pretty", NewJSONMarshaller(true)),
		runtime.WithErrorHandler(DefaultErrorHandler),
		runtime.WithRoutingErrorHandler(DefaultRoutingErrorHandler),
		runtime.WithHealthzEndpoint(hs),
	}

	for _, opt := range opts {
		opt(hs)
	}

	gwmux := runtime.NewServeMux(hs.muxOpts...)
	for _, handler := range hs.handlers {
		if err := gwmux.HandlePath(handler.Method, handler.Path, handler.route); err != nil {
			return nil, err
		}
	}

	routes := make(map[string]struct{})
	for _, s := range services {
		if err := s.RegisterHTTP(gwmux); err != nil {
			return nil, err
		}
		_, rs := s.Routers()
		for _, r := range rs {
			routes[r] = struct{}{}
		}
	}

	var handler http.Handler = gwmux
	handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if hs.auth(gwmux, routes, w, r) {
			return
		}

		gwmux.ServeHTTP(w, r)
	})
	hs.Server = &http.Server{Handler: handler}
	return hs, nil
}

func (s *Server) auth(
	mux *runtime.ServeMux,
	routers map[string]struct{},
	w http.ResponseWriter,
	r *http.Request,
) bool {
	if s.authFunc == nil {
		return false
	}

	_, outbound := runtime.MarshalerForRequest(mux, r)
	if _, ok := routers[r.URL.Path]; !ok {
		return false
	}

	err := s.authFunc(r)
	if err == nil {
		return false
	}

	buf, err := outbound.Marshal(err)
	if err != nil {
		buf = []byte(fallback)
	}
	_, _ = w.Write(buf)
	return true
}
