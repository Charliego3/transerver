package grpcx

import (
	"net"

	"github.com/transerver/app/logger"
	"github.com/transerver/app/opts"
	"github.com/transerver/app/service"
	"google.golang.org/grpc"
)

type Server struct {
	listener net.Listener
	server   *grpc.Server
	srvOpts  []grpc.ServerOption
}

// NewServer returns grpc server instance
func NewServer(opts ...opts.Option[Server]) *Server {
	srv := &Server{}
	srv.init(opts...)
	return srv
}

// init initialize server properties
func (g *Server) init(opts ...opts.Option[Server]) {
	for _, opt := range opts {
		opt.Apply(g)
	}

	if g.listener == nil {
		logger.Fatal("grpc server has no address specified, use WithAddr or WithListener to specify")
	}
	g.server = grpc.NewServer(g.srvOpts...)
}

// Address returns grpc listener addr
func (g *Server) Address() net.Addr {
	return g.listener.Addr()
}

// RegisterService register server to grpc servser
func (g *Server) RegisterService(services ...service.Service) {
	for _, srv := range services {
		g.server.RegisterService(srv.ServiceDesc(), srv)
	}
}

func (g *Server) Run() error {
	return nil
}
