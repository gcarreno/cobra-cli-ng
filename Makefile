# GO env
GO     := $(shell which go)
GOPATH := $(shell $(GO) env | grep GOPATH | cut -d"'" -f2)
GOOS   := $(shell $(GO) env | grep GOOS | cut -d"'" -f2)
ifeq ($(GOOS), set GOOS=windows)
GOOS := windows
endif

# Bin folder
BIN    := bin

# Proper executable name
ifeq ($(GOOS), windows)
EXE    := cobra-cli-ng.exe
else
EXE    := cobra-cli-ng
endif

# File containing "func main()"
MAIN   := main.go

# Project module name
MODULE := github.com/gcarreno/cobra-cli-ng

default: all

all: test build install

test:
	$(info ========== Testing)
	$(GO) test -r "^Test" "$(MODULE)/tests"

build:
	$(info ========== Building for $(GOOS) into $(BIN)/$(EXE))
	mkdir -p $(BIN)
	$(GO) build -o $(BIN)/$(EXE) $(MAIN)

install: test build
	$(info ========== Installing)
	install $(BIN)/$(EXE) $(GOPATH)/bin/$(EXE)

clean:
	$(info ========== Cleaning)
	rm -rfv $(BIN)

.PHONY: all test build install clean