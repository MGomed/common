LOCAL_BIN:=$(CURDIR)/bin

ACCESS_API_PROTO:=access_api_v1
ACCESS_API:=access_api

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

.PHONY: test
test:
	mkdir -p out/coverage
	go clean -testcache
	go test -coverprofile out/coverage/cover.out ./...
	go tool cover -html=cover.out -o out/coverage/coverage.html

install-deps:
	[ -f $(LOCAL_BIN)/golangci-lint ] || GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
	[ -f $(LOCAL_BIN)/protoc-gen-go ] || GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	[ -f $(LOCAL_BIN)/protoc-gen-go-grpc ] || GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

generate:
	make generate_mocks

generate_api:
	make generate_access_api

generate_access_api:
	mkdir -p $(ACCESS_API)
	protoc --proto_path api/$(ACCESS_API_PROTO) --proto_path vendor.protogen \
	--go_out=$(ACCESS_API) --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=$(ACCESS_API) --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/$(ACCESS_API_PROTO)/*

generate_mocks:
	./generate.sh
