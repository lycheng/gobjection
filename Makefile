.PHONY: version init hello tcpserver run-tcpserver

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

test: test-pkg
	@echo "ALL TESTS PASSED"

# Pkg

## Pkg test
test-pkg: test-pkg-ds
	@echo "PKG TESTS PASSED"

### Pkg ds test
test-pkg-ds:
	$(GOTEST) -v ./pkg/ds/lists/...
	$(GOTEST) -v ./pkg/ds/sets/...


# Programs
hello: init
	$(GOBUILD) $(GOFLAGS) -o bin/hello -v ./cmd/hello

echoserver: init
	$(GOBUILD) $(GOFLAGS) -o bin/echoserver -v ./cmd/echoserver

run-echoserver: echoserver
	./bin/echoserver
