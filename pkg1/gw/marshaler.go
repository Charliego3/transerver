package gw

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/transerver/mapp/utils"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func err2resp(_ context.Context, err error) (resp utils.ResponseEntity, code codes.Code) {
	switch e := err.(type) {
	case utils.ResponseEntity:
		code = codes.Internal
		resp = e
	case interface{ GRPCStatus() *status.Status }:
		var msg string
		code = e.GRPCStatus().Code()
		if code == codes.Unavailable {
			msg = "内部异常"
		} else {
			msg = e.GRPCStatus().Message()
		}
		resp = utils.NewErrResponse(convert(code), msg)
	default:
		code = codes.Internal
		resp = utils.NewErrResponse(http.StatusInternalServerError, err.Error())
	}
	return resp, code
}

func respWrap(v any) utils.ResponseEntity {
	switch t := v.(type) {
	case utils.ResponseEntity:
		return t
	case *wrapperspb.BytesValue:
		return utils.NewResponse(t.Value)
	case *wrapperspb.StringValue:
		return utils.NewResponse(t.Value)
	case *wrapperspb.BoolValue:
		return utils.NewResponse(t.Value)
	case *wrapperspb.DoubleValue:
		return utils.NewResponse(t.Value)
	case *wrapperspb.Int32Value:
		return utils.NewResponse(t.Value)
	case *wrapperspb.Int64Value:
		return utils.NewResponse(t.Value)
	case *wrapperspb.UInt32Value:
		return utils.NewResponse(t.Value)
	case *wrapperspb.UInt64Value:
		return utils.NewResponse(t.Value)
	case *wrapperspb.FloatValue:
		return utils.NewResponse(t.Value)
	default:
		return utils.NewResponse(t)
	}
}

type JSONMarshaller struct {
	*runtime.HTTPBodyMarshaler
	pretty bool
}

func (h *JSONMarshaller) ContentType(_ interface{}) string {
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

	resp := respWrap(v)
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
