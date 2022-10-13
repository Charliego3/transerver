package service

import (
	"github.com/google/wire"
	"github.com/transerver/commons/service"
)

var ProviderSet = wire.NewSet(
	MakeServices,
	NewAccountService,
	NewRsaService,
	NewRegionService,
)

func MakeServices(
	s0 *AccountService,
	s1 *PubService,
	s2 *RegionService,
) []service.Service {
	return []service.Service{
		s0,
		s1,
		s2,
	}
}
