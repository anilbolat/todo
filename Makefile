GOPATH := $(shell go env GOPATH)
BIN_DIR := $(GOPATH)/bin
MEGACHECK := $(BIN_DIR)/megacheck
BENCHCMP := $(BIN_DIR)/benchcmp
IMAGE_NAME := heppu/todo

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

.PHONY: docker-image
docker-image:
	docker build --rm -t $(IMAGE_NAME):latest .

.PHONY: docker-run
docker-run:
	docker run --name todo-go --rm --cpuset-cpus='0' -p 8000:8000 $(IMAGE_NAME):latest

.PHONY: docker-image-info
docker-image-info:
	docker images $(IMAGE_NAME):latest
	docker history $(IMAGE_NAME):latest

.PHONY: docker-times
docker-times:
	./docker-times.sh $(IMAGE_NAME)
