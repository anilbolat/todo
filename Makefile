GOPATH := $(shell go env GOPATH)
BIN_DIR := $(GOPATH)/bin
MEGACHECK := $(BIN_DIR)/megacheck
BENCHCMP := $(BIN_DIR)/benchcmp

export GOFLAGS = -mod=vendor

all: lint test

.PHONY: test
test:
	go test -race ./...

.PHONY: benchmark
benchmark:
	go test -run None -bench . ./...

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
