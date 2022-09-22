package commons

import (
	"net"
	"net/http"

	"github.com/soheilhy/cmux"
	"github.com/transerver/commons/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	hs *http.Server
	gs *grpc.Server
	bs config.Bootstrap
	lg *zap.Logger
}

func NewApp(
	bs config.Bootstrap,
	lg *zap.Logger,
	hs *http.Server,
	gs *grpc.Server,
) *App {
	return &App{bs: bs, lg: lg, hs: hs, gs: gs}
}

func (app *App) Run() {
	listener, err := net.Listen("tcp", app.bs.Address())
	if err != nil {
		app.lg.Panic("can not listen", zap.Error(err))
	}

	mux := cmux.New(listener)
	grpcL := mux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := mux.Match(cmux.Any())

	go app.gs.Serve(grpcL)
	go app.hs.Serve(httpL)

	if err := mux.Serve(); err != nil {
		app.lg.Panic("app serve error", zap.Error(err))
	}
}
