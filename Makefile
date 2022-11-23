.PHONY: migrate-init
migrate-init:
	migrate create -ext sql -dir accounts/internal/db/migration/ -seq init_schema
