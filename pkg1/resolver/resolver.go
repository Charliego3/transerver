package resolver

import (
	"context"
	"github.com/transerver/pkg1/logger"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/serviceconfig"
	"sync"
)

type Resolver struct {
	cc  resolver.ClientConn
	wch endpoints.WatchChannel
	ctx context.Context
	cf  context.CancelFunc
	wg  sync.WaitGroup
	sc  *serviceconfig.ParseResult
}

func (r *Resolver) watch() {
	defer r.wg.Done()

	addresses := make(map[string]*endpoints.Update)
	for {
		select {
		case <-r.ctx.Done():
			return
		case ups, ok := <-r.wch:
			if !ok {
				return
			}

			for _, up := range ups {
				switch up.Op {
				case endpoints.Add:
					logger.Sugar().Debugf("Added endpoint: %q -> [%s]", up.Key, up.Endpoint.Addr)
					addresses[up.Key] = up
				case endpoints.Delete:
					logger.Sugar().Debugf("Deleted endpoint: %q -> [%s]", up.Key, up.Endpoint.Addr)
					delete(addresses, up.Key)
				}
			}

			var addrs []resolver.Address
			for _, up := range addresses {
				addr := resolver.Address{
					Addr: up.Endpoint.Addr,
				}
				addrs = append(addrs, addr)
			}

			err := r.cc.UpdateState(resolver.State{Addresses: addrs, ServiceConfig: r.sc})
			if err != nil {
				logger.Sugar().Errorf("Resolver update state error: %v", err)
			}
		}
	}
}

// ResolveNow is a no-op here.
// It's just a hint, Resolver can ignore this if it's not necessary.
func (r *Resolver) ResolveNow(resolver.ResolveNowOptions) {}

func (r *Resolver) Close() {
	r.cf()
	r.wg.Wait()
}
