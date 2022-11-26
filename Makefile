.PHONY: redis
redis:
	docker run --name redis7.0 -p 6379:6379 -d redis:7.0

.PHONY: etcd
etcd:
	docker network create app-tier --driver bridge\
	&& docker run \
		--name etcd\
		--network app-tier\
		--publish 2379:2379\
		--publish 2380:2380\
		--env ALLOW_NONE_AUTHENTICATION=yes\
		--env ETCD_ADVERTISE_CLIENT_URLS=http://etcd-server:2379\
		-d bitnami/etcd:latest