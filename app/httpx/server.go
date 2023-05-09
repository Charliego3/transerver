package httpx

import (
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/transerver/app/service"
)

type Server struct {
	*mux.Router
}

func NewServer() *Server {
	return &Server{
		Router: mux.NewRouter(),
	}
}

func (h *Server) RegisterService(service ...service.Service) {
	h.Path("").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func (h *Server) ListenAndServe(network, addr string) error {
	return nil
}

func (h *Server) Serve(lis net.Listener) error {
	return nil
}

func (h *Server) GracefulShutdown() {

}

func (h *Server) Shutdown() error {
	return nil
}

func (h *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
