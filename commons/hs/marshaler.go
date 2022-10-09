package hs

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/transerver/utils"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

type JSONMarshaller struct {
	*runtime.HTTPBodyMarshaler
	pretty bool
}

func (h *JSONMarshaller) ContentType(v interface{}) string {
	if h.pretty {
		return "application/json+pretty"
	}
	return "application/json"
}

func (h *JSONMarshaller) Marshal(v interface{}) ([]byte, error) {
	if httpBody, ok := v.(*httpbody.HttpBody); ok {
		return h.HTTPBodyMarshaler.Marshal(httpBody)
	}

	if v, ok := v.([]byte); ok {
		return v, nil
	}

	var resp utils.ResponseEntity
	if r, ok := v.(utils.ResponseEntity); ok {
		resp = r
	} else if r, ok := v.(interface {
		GRPCStatus() *status.Status
	}); ok {
		resp = utils.NewErrResponse(r.GRPCStatus().Code(), r.GRPCStatus().Message())
	} else if e, ok := v.(error); ok {
		resp = utils.NewErrResponse(codes.Internal, e.Error())
	} else {
		resp = utils.NewResponse(v)
	}

	if h.pretty {
		return resp.MarshalIdent()
	}
	return resp.Marshal()
}

func NewJSONMarshaller(pretty ...bool) *JSONMarshaller {
	var p bool
	if len(pretty) > 0 {
		p = pretty[0]
	}
	return &JSONMarshaller{
		pretty: p,
		HTTPBodyMarshaler: &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					EmitUnpopulated: true,
				},
			},
		},
	}
}
