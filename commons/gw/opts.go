package gw

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/transerver/commons/errors"
	"io"
	"net/http"
	"net/textproto"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
)

// return Internal when Marshal failed
const fallback = `{"code": 13, "message": "failed to marshal error message"}`

type Option func(*Server)

func WithServeMuxOpts(opts ...runtime.ServeMuxOption) Option {
	return func(gw *Server) {
		gw.muxOpts = append(gw.muxOpts, opts...)
	}
}

func WithHandlerFunc(fn HandlerFunc) Option {
	return func(gw *Server) {
		gw.handler = fn
	}
}

func WithAuthFunc(fn func(*http.Request) error) Option {
	return func(gw *Server) {
		gw.authFunc = fn
	}
}

func WithDialer(ds ...Dialer) Option {
	return func(gw *Server) {
		gw.dialers = ds
	}
}

func WithMiddleware(opts ...mux.MiddlewareFunc) Option {
	return func(gw *Server) {
		gw.middleware = opts
	}
}

func DefaultErrorHandler(
	ctx context.Context,
	_ *runtime.ServeMux,
	marshaller runtime.Marshaler,
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	resp, code := err2resp(ctx, err)
	w.Header().Del("Trailer")
	w.Header().Del("Transfer-Encoding")

	contentType := marshaller.ContentType(resp)
	w.Header().Set("Content-Type", contentType)

	if code == codes.Unauthenticated {
		w.Header().Set("WWW-Authenticate", resp.Error())
	}

	buf, mer := marshaller.Marshal(resp)
	if mer != nil {
		grpclog.Infof("Failed to marshal error message %v: %v", err, mer)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := io.WriteString(w, fallback); err != nil {
			grpclog.Infof("Failed to write response: %v", err)
		}
		return
	}

	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		grpclog.Infof("Failed to extract ServerMetadata from context")
	}

	for k, vs := range md.HeaderMD {
		h := runtime.MetadataHeaderPrefix + k
		for _, v := range vs {
			w.Header().Add(h, v)
		}
	}

	te := r.Header.Get("TE")
	doForwardTrailers := strings.Contains(strings.ToLower(te), "trailers")

	if doForwardTrailers {
		for k := range md.TrailerMD {
			tKey := textproto.CanonicalMIMEHeaderKey(runtime.MetadataTrailerPrefix + k)
			w.Header().Add("Trailer", tKey)
		}
		w.Header().Set("Transfer-Encoding", "chunked")
	}

	st := runtime.HTTPStatusFromCode(code)
	w.WriteHeader(st)
	if _, err := w.Write(buf); err != nil {
		grpclog.Infof("Failed to write response: %v", err)
	}

	if doForwardTrailers {
		for k := range md.TrailerMD {
			tKey := textproto.CanonicalMIMEHeaderKey(runtime.MetadataTrailerPrefix + k)
			w.Header().Add("Trailer", tKey)
		}
	}
}

// DefaultRoutingErrorHandler is our default handler for routing errors.
// By default, http error codes mapped on the following error codes:
//
//	NotFound -> grpc.NotFound
//	StatusBadRequest -> grpc.InvalidArgument
//	MethodNotAllowed -> grpc.Unimplemented
//	Other -> grpc.Internal, method is not expecting to be called for anything else
func DefaultRoutingErrorHandler(
	ctx context.Context,
	mux *runtime.ServeMux,
	marshaller runtime.Marshaler,
	w http.ResponseWriter,
	r *http.Request,
	httpStatus int,
) {
	var sterr error
	switch httpStatus {
	case http.StatusBadRequest:
		sterr = errors.NewArgumentf(ctx, http.StatusText(httpStatus))
	case http.StatusMethodNotAllowed:
		sterr = errors.New(ctx, codes.Unimplemented, http.StatusText(httpStatus))
	case http.StatusNotFound:
		sterr = errors.New(ctx, codes.NotFound, http.StatusText(httpStatus))
	default:
		sterr = errors.NewInternal(ctx, "Unexpected routing error")
	}
	runtime.HTTPError(ctx, mux, marshaller, w, r, sterr)
}
