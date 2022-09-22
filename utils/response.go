package utils

import "google.golang.org/grpc/codes"

type Error struct {
}

type ResponseEntity interface {
	Code() codes.Code
	Error() string
}

type re struct {
	Codes  codes.Code `json:"code"`
	Msg    string     `json:"message"`
	Paylod any        `json:"paylod"`
}

func (r re) Code() codes.Code {
	return r.Codes
}

func (r re) Error() string {
	return r.Msg
}

func NewErrResponse(code codes.Code, msg string) ResponseEntity {
	return &re{Codes: code, Msg: msg}
}

func NewResponse(payload any) ResponseEntity {
	return &re{Codes: codes.OK, Msg: "OK", Paylod: payload}
}
