VERSION=v0.0.1

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

version:
	@echo ${VERSION}

init:
	@mkdir -p bin

# Programs
hello: init
	$(GOBUILD) -o bin/hello -v ./cmd/hello
