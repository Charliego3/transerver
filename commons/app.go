package commons

import (
	"github.com/soheilhy/cmux"
	"github.com/transerver/commons/configs"
	"github.com/transerver/commons/gs"
	"github.com/transerver/commons/hs"
	"go.uber.org/zap"
	"net"
)

type App struct {
	hs *hs.Server
	gs *gs.Server
	bs configs.IConfig
	lg *zap.Logger
}

func NewApp(
	bs configs.IConfig,
	lg *zap.Logger,
	hs *hs.Server,
	gs *gs.Server,
) *App {
	return &App{bs: bs, lg: lg, hs: hs, gs: gs}
}

func (app *App) Run() {
	listener, err := net.Listen("tcp", app.bs.Addr())
	if err != nil {
		app.lg.Panic("can not listen", zap.Error(err))
	}

	mux := cmux.New(listener)
	grpcL := mux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := mux.Match(cmux.Any())

	go app.serve(app.gs.Serve, grpcL)
	go app.serve(app.hs.Serve, httpL)

	app.lg.Info("Server listen on", zap.String("address", listener.Addr().String()))
	if err := mux.Serve(); err != nil {
		app.lg.Panic("app serve error", zap.Error(err))
	}
}

func (app *App) serve(fn func(lis net.Listener) error, lis net.Listener) {
	err := fn(lis)
	if err != nil {
		app.lg.Panic("serve error", zap.Error(err))
	}
}
