package app

import (
	"net"
	"net/http"

	"github.com/pkg/errors"
	"github.com/soheilhy/cmux"
	"github.com/transerver/app/grpcx"
	"github.com/transerver/app/httpx"
	"github.com/transerver/app/logger"
	"github.com/transerver/app/opts"
	"github.com/transerver/app/service"
	"github.com/transerver/app/utils"
)

var Name string

type Application struct {
	// http server is httpx.Server using mux.Router
	http *httpx.Server

	// grpc server is grpcx.Server using grpc
	grpc *grpcx.Server

	// listener to accept http and grpc
	// if cfg.glis and cfg.hlis both nil else is nil
	listener cmux.CMux

	// Application config properties
	cfg *Config
}

// NewApp returns Application
func NewApp(opts ...opts.Option[Config]) *Application {
	app := &Application{}
	app.init(opts...)
	return app
}

// init handling and aggregation options
func (app *Application) init(opts ...opts.Option[Config]) {
	app.cfg = &Config{}
	for _, opt := range opts {
		opt.Apply(app.cfg)
	}

	if utils.Nils(app.cfg.glis, app.cfg.hlis) {
		if app.cfg.lis == nil {
			app.cfg.lis = app.dynamicListener("Application")
		}

		app.listener = cmux.New(app.cfg.lis)
		contentType := http.CanonicalHeaderKey("content-type")
		matcher := cmux.HTTP2MatchHeaderFieldPrefixSendSettings(contentType, "application/grpc")
		app.cfg.glis = app.listener.MatchWithWriters(matcher)
		app.cfg.hlis = app.listener.Match(cmux.Any())
	} else if app.cfg.glis == nil {
		app.cfg.glis = app.dynamicListener("Grpc")
	} else if app.cfg.hlis == nil {
		app.cfg.hlis = app.dynamicListener("Http")
	}

	app.http = httpx.NewServer(httpx.WithListener(app.cfg.hlis))
	app.grpc = grpcx.NewServer(grpcx.WithListener(app.cfg.glis))
}

// dynamicListener if app without any listener specifies then create a dynamic listener
func (app *Application) dynamicListener(server string) net.Listener {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		logger.Fatal("failed to listen", "server", server, "err", err)
	}

	logger.Warn("No address is specified so dynamic addresses are used",
		"server", server, "address", app.cfg.lis.Addr().String())
	return listener
}

// Address returns application listen address
// this address is http and grpc both
func (app *Application) Address() net.Addr {
	if app.cfg.lis == nil {
		return nil
	}
	return app.cfg.lis.Addr()
}

// RegisterService add service to http and grpc server
func (app *Application) RegisterService(services ...service.Service) {
	app.http.RegisterService(services...)
	app.grpc.RegisterService(services...)
}

// Run start the server until terminate
func (app *Application) Run() (err error) {
	go func() {
		herr := app.grpc.Run()
		if herr != nil {
			if err == nil {
				err = errors.Wrap(herr, "grpc server got an error")
			} else {
				err = errors.Wrap(err, errors.Wrap(herr, "grpc server got an error").Error())
			}
		}
	}()
	go func() {
		err = app.http.Run()
		if err != nil {
			errors.Wrap(err, "http server got an error")
		}
	}()

	if app.listener != nil {
		err = app.listener.Serve()
	}
	return nil
}
