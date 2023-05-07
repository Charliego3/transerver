package httpx

import (
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/transerver/app/service"
)

type HttpServer struct {
	*mux.Router
}

func NewServer() *HttpServer {
	return &HttpServer{
		Router: mux.NewRouter(),
	}
}

func (h *HttpServer) RegisterService(service service.Service) {

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
