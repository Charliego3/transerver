.PHONY: build
build:
	go build -ldflags="-X github.com/transerver/pkg1/app.Name=accounts" -o ../bin/accounts .

.PHONY: run
run:
	go run ./cmd/... --config ./internal/conf/config.yaml -i18n ./internal/i18n

.PHONY: postgres
postgres:
	docker run\
		--name postgres15.1\
		-p 5432:5432\
		-e POSTGRES_DB=accounts\
		-e POSTGRES_USER=root\
		-e POSTGRES_PASSWORD=root\
		-d postgres:15.1

.PHONY: sqlc
sqlc:
	cd ./internal/data && sqlc generate

.PHONY: migrate
migrate:
	migrate create -ext sql -dir internal/data/migration/ -seq init_schema

migrateup:
	migrate -path internal/data/migration -database "postgresql://root:root@localhost:5432/accounts?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/data/migration -database "postgresql://root:root@localhost:5432/accounts?sslmode=disable" -verbose down
