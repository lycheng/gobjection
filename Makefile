.PHONY: version init hello

_GIT_LAST_COMMIT_TIME=$(shell TZ=UTC git log --pretty=format:'%cd' -1 --date=format-local:'%Y%m%d-%H%M%S')
_GIT_LAST_COMMIT_HASH=$(shell git rev-parse --short HEAD)

_VERSION=$(_GIT_LAST_COMMIT_TIME).$(_GIT_LAST_COMMIT_HASH)

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

GOLDFLAGS += -X main.Version=$(_VERSION)
GOFLAGS = -ldflags "$(GOLDFLAGS)"

version:
	@echo ${_VERSION}

init:
	@mkdir -p bin

# Programs
hello: init
	$(GOBUILD) $(GOFLAGS) -o bin/hello -v ./cmd/hello

tcpserver: init
	$(GOBUILD) $(GOFLAGS) -o bin/tcpserver -v ./cmd/tcpserver

start-tcpserver: tcpserver
	./bin/tcpserver
