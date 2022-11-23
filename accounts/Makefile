.PHONY: build
build:
	go build -ldflags="-X github.com/transerver/commons/app.Name=accounts" -o ../bin/accounts .

.PHONY: run
run:
	go run ./cmd/... -config ./internal/conf/config.yaml -i18n ./internal/i18n

.PHONY: sqlc
sqlc:
	cd ./internal/db && sqlc generate

.PHONY: ent
ent:
	cd internal/ent \
	&& go run entgo.io/ent/cmd/ent generate --idtype int64 ./schema
