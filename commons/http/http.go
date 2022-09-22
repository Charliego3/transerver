package http

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func NewHTTPServer(muxOpts []runtime.ServeMuxOption) *http.Server {
	muxOpts = append(muxOpts, runtime.WithMarshalerOption("application/json", NewJSONMarshaler()))
	mux := runtime.NewServeMux(muxOpts...)
	server := &http.Server{
		Handler: mux,
	}
	return server
}
