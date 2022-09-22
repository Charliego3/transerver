package hs

import (
	"context"
	"io"
	"net/http"
	"net/textproto"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/transerver/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
)

func DefaultErrorHandler(
	ctx context.Context,
	mux *runtime.ServeMux,
	marshaler runtime.Marshaler,
	w http.ResponseWriter,
	r *http.Request,
	err error,
) {
	// return Internal when Marshal failed
	const fallback = `{"code": 13, "message": "failed to marshal error message"}`

	var rtn utils.ResponseEntity
	if v, ok := err.(utils.ResponseEntity); !ok {
		rtn = utils.NewErrResponse(codes.Unknown, err.Error())
	} else {
		rtn = v
	}

	w.Header().Del("Trailer")
	w.Header().Del("Transfer-Encoding")

	contentType := marshaler.ContentType(rtn)
	w.Header().Set("Content-Type", contentType)

	if rtn.Code() == codes.Unauthenticated {
		w.Header().Set("WWW-Authenticate", rtn.Error())
	}

	buf, merr := marshaler.Marshal(rtn)
	if merr != nil {
		grpclog.Infof("Failed to marshal error message %v: %v", rtn, merr)
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

	st := runtime.HTTPStatusFromCode(rtn.Code())
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
