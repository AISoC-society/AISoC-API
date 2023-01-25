BINARY:=api
BUILD_FLAGS=-ldflags '-s -w'
FMT_REQUIRED:=$(shell gofmt -l $(shell find . -type f -iname *.go))

all: build
build:
	go mod tidy
	go build $(BUILD_FLAGS) ./cmd/$(BINARY)/

run: build
	./$(BINARY)

check:
	go test ./...
	@echo $(FMT_REQUIRED)
	@test -z $(FMT_REQUIRED)
	go vet ./...

allocs:
	$(MAKE) BUILD_FLAGS="-gcflags '-m' $(BUILD_FLAGS)" -s

release:
	go mod tidy
	mkdir -p release
	GOOS="linux" GOARCH="386" go build -ldflags "-s -w" -o release/$(BINARY)-linux-386 ./cmd/$(BINARY)/*.go
	GOOS="linux" GOARCH="amd64" go build -ldflags "-s -w" -o release/$(BINARY)-linux-amd64 ./cmd/$(BINARY)/*.go
	GOOS="linux" GOARCH="arm" go build -ldflags "-s -w" -o release/$(BINARY)-linux-arm ./cmd/$(BINARY)/*.go
	GOOS="linux" GOARCH="arm64" go build -ldflags "-s -w" -o release/$(BINARY)-linux-arm64 ./cmd/$(BINARY)/*.go

clean:
	go clean
	$(RM) -rf ./release
	find . -name "*.json" -type f|xargs rm -f

.PHONY: build release clean run check allocs
