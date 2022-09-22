package http

import (
	"encoding/json"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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
		return httpBody.Data, nil
	}

	if v, ok := v.([]byte); ok {
		return v, nil
	}

	return json.Marshal(v)
}

func NewJSONMarshaler() *JSONMarshaler {
	return &JSONMarshaler{}
}
