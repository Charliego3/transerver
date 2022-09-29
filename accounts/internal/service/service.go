package service

import (
	"github.com/google/wire"
	"github.com/transerver/commons"
)

var ProviderSet = wire.NewSet(
	MakeServices,
	NewAccountService,
	NewRsaService,
)

func MakeServices(
	s0 *AccountService,
	s1 *RsaService,
) []commons.Service {
	return []commons.Service{
		s0,
		s1,
	}
}
