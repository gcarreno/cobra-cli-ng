BIN := bin
EXE := cobra-cli-ng
MAIN := main.go
MODULE := github.com/gcarreno/cobra-cli-ng
GO := $(shell which go)

ifndef GOPATH
GOPATH := $(shell $(GO) env | grep GOPATH | cut -d"'" -f2)
endif

default: all

all: test build install

test:
	$(info ========== Testing)
	$(GO) test -r "^Test" "$(MODULE)/tests"

build:
	$(info ========== Building $(@))
	mkdir -p $(BIN)
	$(GO) build -o $(BIN)/$(EXE) $(MAIN)

install: test build
	$(info ========== Installing)
	install $(BIN)/$(EXE) $(GOPATH)/bin/$(EXE)

clean:
	$(info ========== Cleaning)
	rm -rfv $(BIN)

.PHONY: all test build install clean