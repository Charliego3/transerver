package grpcx

import (
	"net"
	"net/http"

	"github.com/transerver/app/logger"
	"github.com/transerver/app/opts"
	"github.com/transerver/app/service"
	"google.golang.org/grpc"
)

type Server struct {
	addr    string
	server  *grpc.Server
	srvOpts []grpc.ServerOption
}

func NewServer(opts ...opts.Option[Server]) *Server {
	srv := &Server{}
	srv.getOpts(opts...)
	return srv
}

func (g *Server) getOpts(opts ...opts.Option[Server]) {
	for _, opt := range opts {
		opt.Apply(g)
	}

	g.server = grpc.NewServer(g.srvOpts...)
}

func (g *Server) RegisterService(services ...service.Service) {
	for _, srv := range services {
		g.server.RegisterService(srv.ServiceDesc(), srv)
	}
}

func (g *Server) ListenAndServe(network, addr string) error {
	lis, err := net.Listen(network, addr)
	if err != nil {
		return err
	}
	return g.Serve(lis)
}

func (g *Server) Serve(lis net.Listener) error {
	logger.Info("Grpc.Server on", "address", lis.Addr().String())
	return g.server.Serve(lis)
}

func (g *Server) GracefulShutdown() {

}

func (g *Server) Shutdown() error {
	return nil
}

func (g *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.server.ServeHTTP(w, r)
}
