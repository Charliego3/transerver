package httpx

import (
	"net"
	"net/http"
)

type HttpServer struct {
}

func (h *HttpServer) ListenAndServe(network, addr string) error {
	return nil
}

func (h *HttpServer) Serve(lis net.Listener) error {
	return nil
}

func (h *HttpServer) GracefulShutdown() {

}

func (h *HttpServer) Shutdown() error {
	return nil
}

func (h *HttpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
