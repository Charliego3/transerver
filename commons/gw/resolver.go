package gw

import (
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
	"google.golang.org/grpc/serviceconfig"
)

const (
	Schema = "discovery"
)

// ManualResolver is a ManualResolver (and resolver.Builder) that can be updated
// using SetEndpoints.
type ManualResolver struct {
	*manual.Resolver
	endpoints     []string
	serviceConfig *serviceconfig.ParseResult
}

func NewResolver(endpoints ...string) *ManualResolver {
	r := manual.NewBuilderWithScheme(Schema)
	return &ManualResolver{Resolver: r, endpoints: endpoints, serviceConfig: nil}
}

// Build returns itself for ManualResolver, because it's both a builder and a resolver.
func (r *ManualResolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r.serviceConfig = cc.ParseServiceConfig(`{"loadBalancingPolicy": "round_robin"}`)
	if r.serviceConfig.Err != nil {
		return nil, r.serviceConfig.Err
	}
	res, err := r.Resolver.Build(target, cc, opts)
	if err != nil {
		return nil, err
	}
	// Populates endpoints stored in r into ClientConn (cc).
	r.updateState()
	return res, nil
}

func (r *ManualResolver) SetEndpoints(endpoints []string) {
	r.endpoints = endpoints
	r.updateState()
}

func (r *ManualResolver) updateState() {
	if r.CC != nil {
		addresses := make([]resolver.Address, len(r.endpoints))
		for i, ep := range r.endpoints {
			addr, serverName := interpret(ep)
			addresses[i] = resolver.Address{Addr: addr, ServerName: serverName}
		}
		state := resolver.State{
			Addresses:     addresses,
			ServiceConfig: r.serviceConfig,
		}
		r.UpdateState(state)
	}
}
