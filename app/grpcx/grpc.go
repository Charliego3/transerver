package grpcx

import (
	"github.com/transerver/app"
	"github.com/transerver/app/logger"
	"github.com/transerver/app/service"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

type GrpcServer struct {
	addr    string
	server  *grpc.Server
	srvOpts []grpc.ServerOption
}

func NewGrpcServer(opts ...Option) *GrpcServer {
	srv := &GrpcServer{}
	srv.getOpts(opts...)
	return srv
}

func (g *GrpcServer) getOpts(opts ...Option) {
	for _, opt := range opts {
		opt(g)
	}

	g.server = grpc.NewServer(g.srvOpts...)
}

func (g *GrpcServer) RegisterService(services ...service.Service) app.Server {
	for _, srv := range services {
		g.server.RegisterService(srv.ServiceDesc(), srv)
	}
	return g
}

func (g *GrpcServer) ListenAndServe(network, addr string) error {
	lis, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	return g.Serve(lis)
}

func (g *GrpcServer) Serve(lis net.Listener) error {
	logger.Info("Grpc.Server on", "address", lis.Addr().String())
	return g.server.Serve(lis)
}

func (g *GrpcServer) GracefulShutdown() {

}

func (g *GrpcServer) Shutdown() error {
	return nil
}

func (g *GrpcServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.server.ServeHTTP(w, r)
}
