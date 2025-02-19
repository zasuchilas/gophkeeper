run:
	@CONFIG_PATH=config/local.yml go run ./cmd/server

cli:
	@go run ./cmd/client

cli_port:
	@go run ./cmd/client -s localhost:9999



# ---------------------------------------------------------------------------------------------------------------------
# db migrations by goose

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

# ---------------------------------------------------------------------------------------------------------------------
# grpc develop

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


# test

test:
	@go test -v ./...

clean_test_cache:
	@go clean -testcache


# build

build_client_linux:
	#@cd ./cmd/client/ && GOOS=linux GOARCH=amd64 go build -ldflags "-X main.buildVersion=$(git describe --tags --abbrev=0) -X 'main.buildDate=$(date +'%Y.%m.%d %H:%M:%S')' -X main.buildCommit=$(git rev-parse HEAD)" -o gophkeeper_linux_amd64
	# GOOS=linux GOARCH=amd64 go build -ldflags "-X main.buildVersion=$(git describe --tags --abbrev=0) -X 'main.buildDate=$(date +'%Y.%m.%d %H:%M:%S')' -X main.buildCommit=$(git rev-parse HEAD)" -o gophkeeper_linux_amd64

build_client_windows:
	# GOOS=windows GOARCH=amd64 go build -ldflags "-X main.buildVersion=$(git describe --tags --abbrev=0) -X 'main.buildDate=$(date +'%Y.%m.%d %H:%M:%S')' -X main.buildCommit=$(git rev-parse HEAD)" -o gophkeeper_windows_amd64

build_client_darwin:
	# GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.buildVersion=$(git describe --tags --abbrev=0) -X 'main.buildDate=$(date +'%Y.%m.%d %H:%M:%S')' -X main.buildCommit=$(git rev-parse HEAD)" -o gophkeeper_darwin_amd64
	# GOOS=darwin GOARCH=arm64 go build -ldflags "-X main.buildVersion=$(git describe --tags --abbrev=0) -X 'main.buildDate=$(date +'%Y.%m.%d %H:%M:%S')' -X main.buildCommit=$(git rev-parse HEAD)" -o gophkeeper_darwin_arm64
