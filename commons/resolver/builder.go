package resolver

import (
	"context"
	"github.com/transerver/commons/es"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/status"
)

type Builder struct{}

func (b Builder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &Resolver{cc: cc}
	r.ctx, r.cf = context.WithCancel(context.Background())

	if !opts.DisableServiceConfig {
		r.sc = cc.ParseServiceConfig(`{"loadBalancingPolicy": "round_robin"}`)
		if r.sc.Err != nil {
			return nil, r.sc.Err
		}
	}

	em, err := endpoints.NewManager(es.C().Client, target.URL.Path)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Resolver: failed to new endpoint manager: %s", err)
	}
	r.wch, err = em.NewWatchChannel(r.ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Resolver: failed to new watch channel: %s", err)
	}

	r.wg.Add(1)
	go r.watch()
	return r, nil
}

func (b Builder) Scheme() string {
	return "discovery"
}

// NewBuilder creates a Resolver Builder.
func NewBuilder() resolver.Builder {
	return Builder{}
}
