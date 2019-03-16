# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOFMT=$(GOCMD)fmt
GOVET=$(GOCMD) vet
DOCKERBUILD=docker build

BINARY_NAME=port-domain-svc

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

DEP=dep

.PHONY: all
all: format vet test build

format:
	$(GOFMT) -l -w $(SRC)

build:
	$(GOBUILD) -o $(BINARY_NAME)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

deps:
	$(DEP) ensure

vet:
	$(GOVET) ./...

docker:
	$(DOCKERBUILD) -t port-domain-svc .

grpc:
	protoc --go_out=plugins=grpc:./src ./proto/port-domain-svc.proto
