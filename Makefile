VERSION := $(shell cat VERSION)
LD_FLAGS := -X github.com/saeedafzal/sandbox-service/config.version=$(VERSION)

build:
	go build -ldflags="$(LD_FLAGS) -s -w" -o bin/sandbox-service

run: build
	./bin/sandbox-service -debug

test:
	go test ./...
