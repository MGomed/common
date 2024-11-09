LOCAL_BIN:=$(CURDIR)/bin

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

.PHONY: test
test:
	mkdir -p out/coverage
	go clean -testcache
	go test -coverprofile out/coverage/cover.out ./...
	go tool cover -html=cover.out -o out/coverage/coverage.html

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

generate:
	make generate_mocks

generate_mocks:
	./generate.sh
