GOPATH := $(shell go env GOPATH)
BIN_DIR := $(GOPATH)/bin
MEGACHECK := $(BIN_DIR)/megacheck

export GOFLAGS = -mod=vendor

all: lint test

.PHONY: test
test:
	go test -race ./...

.PHONY: lint
lint: $(MEGACHECK)
	$(MEGACHECK) ./...

.PHONY: vendor
vendor:
	go mod vendor
	go mod tidy

$(MEGACHECK):
	GO111MODULE=off go get -u honnef.co/go/tools/cmd/megacheck
