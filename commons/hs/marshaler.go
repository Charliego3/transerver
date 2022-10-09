package hs

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/transerver/utils"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func solveResponse(v any) utils.ResponseEntity {
	var resp utils.ResponseEntity
	switch t := v.(type) {
	case utils.ResponseEntity:
		resp = t
	case interface{ GRPCStatus() *status.Status }:
		resp = utils.NewErrResponse(t.GRPCStatus().Code(), t.GRPCStatus().Message())
	case error:
		resp = utils.NewErrResponse(codes.Internal, t.Error())
	case *wrapperspb.BytesValue:
		resp = utils.NewResponse(t.Value)
	case *wrapperspb.StringValue:
		resp = utils.NewResponse(t.Value)
	case *wrapperspb.BoolValue:
		resp = utils.NewResponse(t.Value)
	case *wrapperspb.DoubleValue:
		resp = utils.NewResponse(t.Value)
	case *wrapperspb.Int32Value:
		resp = utils.NewResponse(t.Value)
	case *wrapperspb.Int64Value:
		resp = utils.NewResponse(t.Value)
	case *wrapperspb.UInt32Value:
		resp = utils.NewResponse(t.Value)
	case *wrapperspb.UInt64Value:
		resp = utils.NewResponse(t.Value)
	case *wrapperspb.FloatValue:
		resp = utils.NewResponse(t.Value)
	default:
		resp = utils.NewResponse(t)
	}
	return resp
}

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

	resp := solveResponse(v)
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
