package hs

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/transerver/commons"
)

func NewHTTPServerWithoutOpts(
	services []commons.Service,
) (*http.Server, error) {
	return NewHTTPServerWithOptions(services, nil)
}

func NewHTTPServerWithoutHandlers(
	services []commons.Service,
	muxOpts ...runtime.ServeMuxOption,
) (*http.Server, error) {
	return NewHTTPServerWithOptions(services, nil, muxOpts...)
}

func NewHTTPServerWithoutMuxOpts(
	services []commons.Service,
	handlers []Handler,
) (*http.Server, error) {
	return NewHTTPServerWithOptions(services, handlers)
}

func NewHTTPServerWithOptions(
	services []commons.Service,
	handlers []Handler,
	muxOpts ...runtime.ServeMuxOption,
) (*http.Server, error) {
	muxOpts = append(muxOpts, runtime.WithMarshalerOption("application/json", NewJSONMarshaller()))
	muxOpts = append(muxOpts, runtime.WithMarshalerOption("application/json+pretty", NewJSONMarshaller(true)))
	mux := runtime.NewServeMux(muxOpts...)
	for _, handler := range handlers {
		if err := mux.HandlePath(handler.Method, handler.Path, handler.route); err != nil {
			return nil, err
		}
	}

	for _, service := range services {
		if err := service.RegisterHTTP(mux); err != nil {
			return nil, err
		}
	}
	return &http.Server{Handler: handlerFunc(mux)}, nil
}

func handlerFunc(mux *runtime.ServeMux) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		if strings.HasSuffix(r.URL.Path, "register") {
			_, _ = w.Write([]byte("you must be authed..."))
			return
		}
		mux.ServeHTTP(w, r)
	}
}

func DefaultServeMuxOpts() []runtime.ServeMuxOption {
	return []runtime.ServeMuxOption{
		runtime.WithErrorHandler(DefaultErrorHandler),
		runtime.WithRoutingErrorHandler(DefaultRoutingErrorHandler),
	}
}
