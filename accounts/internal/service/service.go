package service

import (
	"github.com/google/wire"
	"github.com/transerver/pkg1/gs"
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
) []gs.Service {
	return []gs.Service{
		s0,
		s1,
		s2,
	}
}
