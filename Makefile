run:
	@CONFIG_PATH=config/local.yml go run ./cmd/server

goose_create:
	@read -p "Enter migration name: " MIGRATION_NAME; \
	goose -dir db/migrations create "$$MIGRATION_NAME" go

goose_up:
	@goose -dir db//migrations postgres "postgresql://gophkeeper:pass@127.0.0.1:9998/gophkeeper?sslmode=disable" up

goose_down:
	@goose -dir db//migrations postgres "postgresql://gophkeeper:pass@127.0.0.1:9998/gophkeeper?sslmode=disable" down

