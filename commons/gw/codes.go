package gw

import (
	"google.golang.org/grpc/codes"
	"net/http"
)

func convert(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return http.StatusNotAcceptable
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusAlreadyReported
	case codes.PermissionDenied:
	case codes.ResourceExhausted:
	case codes.FailedPrecondition:
	case codes.Aborted:
	case codes.OutOfRange:
	case codes.Unimplemented:
	case codes.Internal:
	case codes.Unavailable:
	case codes.DataLoss:
	case codes.Unauthenticated:
	}
	return http.StatusOK
}
