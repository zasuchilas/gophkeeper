run:
	@CONFIG_PATH=config/local.yml go run ./cmd/server

goose_create:
	@read -p "Enter migration name: " MIGRATION_NAME; \
	goose -dir db/server/migrations create "$$MIGRATION_NAME" go

goose_create_sql:
	@read -p "Enter migration name: " MIGRATION_NAME; \
	goose -dir db/server/migrations create "$$MIGRATION_NAME" sql

goose_up:
	@goose -dir db/server/migrations postgres "postgresql://gophkeeper:pass@127.0.0.1:9998/gophkeeper?sslmode=disable" up

goose_down:
	@goose -dir db/server/migrations postgres "postgresql://gophkeeper:pass@127.0.0.1:9998/gophkeeper?sslmode=disable" down

LOCAL_BIN:=$(CURDIR)/bin
install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

vendor-proto:
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi

generate:
	make generate-user-api
	make generate-secrets-api

generate-user-api:
	mkdir -p pkg/userv1
	protoc --proto_path api/userv1 \
	--go_out=pkg/userv1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/userv1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/userv1/user.proto

generate-secrets-api:
	mkdir -p pkg/secretsv1
	protoc --proto_path api/secretsv1 \
	--go_out=pkg/secretsv1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/secretsv1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/secretsv1/secrets.proto
