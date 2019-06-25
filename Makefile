export GO111MODULE=on
# overwrite host system gopath with current dir
# uncomment to enable
# export GOPATH := $(shell pwd)

# Go parameters
GOCMD=go
GOINSTALL=$(GOCMD) install
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

PACKAGE_NAME="github.com/zerjioang/etherniti"
BINARY_NAME="etherniti"
BINARY_UNIX=$(BINARY_NAME)_unix

all: deps test build
install:
	$(GOINSTALL) $(PACKAGE_NAME)
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
deps:
	dep ensure -v
upgrade:
	dep ensure -update -v
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v