package app

import (
	"net"
	"net/http"
)

type Server interface {
	Shutdown() error
	GracefulShutdown()
	Serve(lis net.Listener) error
	ListenAndServe(network, addr string) error
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
