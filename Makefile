GOPATH := $(shell go env GOPATH)
BIN_DIR := $(GOPATH)/bin
MEGACHECK := $(BIN_DIR)/megacheck
BENCHCMP := $(BIN_DIR)/benchcmp

export GOFLAGS = -mod=vendor

all: lint test build

.PHONY: test
test:
	go test -race ./...

.PHONY: test-integration
test-integration:
	go test -tags=integration -race ./...

.PHONY: build
build:
	CGO_ENABLED=0 go build -o builds/todo cmd/todo/main.go

.PHONY: benchmark
benchmark:
	CGO_ENABLED=0 go test -run None -bench . ./...

.PHONY: lint
lint: $(MEGACHECK)
	$(MEGACHECK) ./...

.PHONY: vendor
vendor:
	go mod vendor
	go mod tidy

.PHONY: benchcmp
benchcmp: $(BENCHCMP)
	$(BENCHCMP) old.txt new.txt

$(BENCHCMP):
	GO111MODULE=off go get -u golang.org/x/tools/cmd/benchcmp

$(MEGACHECK):
	GO111MODULE=off go get -u honnef.co/go/tools/cmd/megacheck
