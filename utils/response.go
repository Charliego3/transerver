package utils

import (
	json "github.com/json-iterator/go"
	"google.golang.org/grpc/codes"
)

type Error struct {
}

type ResponseEntity interface {
	Code() codes.Code
	Error() string
	Marshal() ([]byte, error)
	MarshalIdent() ([]byte, error)
}

type re struct {
	Codes   codes.Code `json:"code"`
	Msg     string     `json:"message"`
	Payload any        `json:"payload,omitempty"`
}

func (r re) Code() codes.Code {
	return r.Codes
}

func (r re) Error() string {
	return r.Msg
}

func (r re) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r re) MarshalIdent() ([]byte, error) {
	return json.MarshalIndent(r, "", "    ")
}

func NewErrResponse(code codes.Code, msg string) ResponseEntity {
	return &re{Codes: code, Msg: msg}
}

func NewResponse(payload any) ResponseEntity {
	return &re{Codes: codes.OK, Msg: "OK", Payload: payload}
}
