package service

import (
	"github.com/google/wire"
	"github.com/transerver/commons"
)

var ProviderSet = wire.NewSet(
	MakeServices,
	NewGreeterSerivce,
)

func MakeServices(
	s0 *GreeterService,
) []commons.Service {
	return []commons.Service{
		s0,
	}
}
