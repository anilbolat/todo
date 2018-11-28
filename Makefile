GOPATH := $(shell go env GOPATH)
BIN_DIR := $(GOPATH)/bin
MEGACHECK := $(BIN_DIR)/megacheck
BENCHCMP := $(BIN_DIR)/benchcmp
VEGETA := $(BIN_DIR)/vegeta
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

.PHONY: docker-stats
docker-stats:
	docker stats --format "table {{.Name}} \t{{.CPUPerc}}\t{{.MemUsage}}"

.PHONY: wrk
wrk:
	wrk -t2 -c100 -d30s http://127.0.0.1:8000/api/tasks/1

.PHONY: vegeta
vegeta: $(VEGETA)
	echo "GET http://127.0.0.1:8000/api/tasks/1" | $(VEGETA) attack -name go -duration 30s -rate 2000 | tee result.bin | $(VEGETA) report; $(VEGETA) plot result.bin > plot.html

$(VEGETA):
	GO111MODULE=off go get -u github.com/tsenart/vegeta
