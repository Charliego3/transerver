package resolver

import (
	"context"
	"github.com/gookit/goutil/strutil"
	json "github.com/json-iterator/go"
	"github.com/transerver/pkg1/es"
	"github.com/transerver/utils"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/status"
)

type Builder struct {
	config ServiceConfig
}

type ServiceConfig struct {
	LoadBalancingPolicy string `json:"loadBalancingPolicy" yaml:"loadBalancingPolicy"`
}

func (b Builder) Build(target resolver.Target, cc resolver.ClientConn, opt resolver.BuildOptions) (resolver.Resolver, error) {
	r := &Resolver{cc: cc}
	r.ctx, r.cf = context.WithCancel(context.Background())

	if !opt.DisableServiceConfig {
		const defaultConfig = `{"loadBalancingPolicy": "round_robin"}`
		var config = defaultConfig
		if strutil.IsNotBlank(b.config.LoadBalancingPolicy) {
			buf, err := json.Marshal(b.config)
			if err != nil {
				return nil, err
			}
			config = utils.String(buf)
		}
		r.sc = cc.ParseServiceConfig(config)
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
func NewBuilder(config ServiceConfig) resolver.Builder {
	return Builder{config}
}
