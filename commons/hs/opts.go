package hs

import (
	"context"
	"google.golang.org/grpc/status"
	"io"
	"net/http"
	"net/textproto"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/transerver/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
)

// return Internal when Marshal failed
const fallback = `{"code": 13, "message": "failed to marshal error message"}`

type Option func(*Server)

func WithServeMuxOpts(opts ...runtime.ServeMuxOption) Option {
	return func(hs *Server) {
		hs.muxOpts = append(hs.muxOpts, opts...)
	}
}

func WithHandlers(handlers ...Handler) Option {
	return func(hs *Server) {
		hs.handlers = handlers
	}
}

func WithAuthFunc(fn func(*http.Request) error) Option {
	return func(hs *Server) {
		hs.authFunc = fn
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
	code := codes.Internal
	if r, ok := err.(interface{ GRPCStatus() *status.Status }); ok {
		code = r.GRPCStatus().Code()
	} else if r, ok := err.(interface{ Code() codes.Code }); ok {
		code = r.Code()
	}

	w.Header().Del("Trailer")
	w.Header().Del("Transfer-Encoding")

	contentType := marshaller.ContentType(err)
	w.Header().Set("Content-Type", contentType)

	if code == codes.Unauthenticated {
		w.Header().Set("WWW-Authenticate", err.Error())
	}

	buf, mer := marshaller.Marshal(err)
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
	sterr := utils.NewErrResponse(codes.Internal, "Unexpected routing error")
	switch httpStatus {
	case http.StatusBadRequest:
		sterr = utils.NewErrResponse(codes.InvalidArgument, http.StatusText(httpStatus))
	case http.StatusMethodNotAllowed:
		sterr = utils.NewErrResponse(codes.Unimplemented, http.StatusText(httpStatus))
	case http.StatusNotFound:
		sterr = utils.NewErrResponse(codes.NotFound, http.StatusText(httpStatus))
	}
	runtime.HTTPError(ctx, mux, marshaller, w, r, sterr)
}
