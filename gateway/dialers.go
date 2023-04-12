package main

import (
	"github.com/transerver/pkg1/gw"
	"github.com/transerver/protos/acctspb"
)

var dialers = map[string][]gw.DialerFunc{
	"accounts": {
		acctspb.RegisterAccountServiceHandler,
		acctspb.RegisterRsaServiceHandler,
		acctspb.RegisterRegionServiceHandler,
	},
}
