.PHONY: install build

default: build

LIBDIR = /usr/local/lib/geodelib

# tell go to install to the current directory
export GOBIN=$(shell pwd)

install: libinstall bininstall

build:
	@go build -v -o bin/geode ./pkg/cmd/geode

libinstall:
	@rm -rf $(LIBDIR)
	@mkdir -p $(LIBDIR)
	@cp -a lib/* $(LIBDIR)
	
bininstall:
	@install ./bin/geode /usr/local/bin

gen:
	go generate -v ./...


all: build install
