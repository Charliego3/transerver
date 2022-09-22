package hs

import (
	"encoding/json"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/transerver/utils"
	"google.golang.org/genproto/googleapis/api/httpbody"
)

type JSONMarshaler struct {
	*runtime.HTTPBodyMarshaler
}

func (h *JSONMarshaler) ContentType(v interface{}) string {
	return "application/json"
}

// Marshal marshals "v" by returning the body bytes if v is a
// google.api.HttpBody message, otherwise it falls back to the default Marshaler.
func (h *JSONMarshaler) Marshal(v interface{}) ([]byte, error) {
	if httpBody, ok := v.(*httpbody.HttpBody); ok {
		return h.HTTPBodyMarshaler.Marshal(httpBody)
	}

	if v, ok := v.([]byte); ok {
		return v, nil
	}

	if _, ok := v.(utils.ResponseEntity); !ok {
		v = utils.NewResponse(v)
	}

	return json.Marshal(v)
}

func NewJSONMarshaler() *JSONMarshaler {
	return &JSONMarshaler{}
}
