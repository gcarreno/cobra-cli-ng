# GO env
GO      := $(shell which go)
GOPATH  := $(shell $(GO) env | grep GOPATH | cut -d"'" -f2)
GOOS    := $(shell $(GO) env | grep GOOS | cut -d"'" -f2)
ifeq ($(GOOS), set GOOS=windows)
GOOS := windows
endif

# File containing "func main()"
MAIN := main.go

# Project module name
MODULE := github.com/gcarreno/cobra-cli-ng

# Bin folder
BIN := bin

# Version
VERSION := $(shell git describe --abbrev=0 --tags)

# Proper executable name
ifeq ($(GOOS), windows)
EXE := cobra-cli-ng.exe
BINARY64 := cobra-cli-ng-$(GOOS)_amd64.exe
else
EXE := cobra-cli-ng
BINARY64 := cobra-cli-ng-$(GOOS)_amd64
endif
RELEASE64 := cobra-cli-ng-$(VERSION)-$(GOOS)_amd64

# Silence  directory printing
MAKEFLAGS += --no-print-directory

default: all

all: test install

test:
	$(info ========== Testing)
	$(GO) test -v -run "^Test" "$(MODULE)/tests"

binary: target/$(BINARY64)

ifeq ($(GOOS), windows)
release: binary
	$(info ========== Building release for "$(GOOS)" "$(RELEASE64)")
	@cd target && md5sum $(BINARY64) > $(RELEASE64).md5
	@cd target && zip -q $(RELEASE64).zip $(BINARY64) $(RELEASE64).md5
	@cd target && rm -f $(BINARY64) $(RELEASE64).md5
else
release: binary
	$(info ========== Building release for "$(GOOS)" "$(RELEASE64)")
	@cd target && md5sum $(BINARY64) > $(RELEASE64).md5
	@cd target && tar -czf $(RELEASE64).tgz $(BINARY64) $(RELEASE64).md5
	@cd target && rm -f $(BINARY64) $(RELEASE64).md5
endif

release-all: clean
	$(info ========== Release all $(VERSION))
	@GOOS=darwin $(MAKE) release
	@GOOS=linux $(MAKE) release
	@GOOS=windows $(MAKE) release

target:
	@mkdir -p $@

target/$(BINARY64):
	$(info ========== Building "$(@)")
	@CGO_ENABLED=0 GOARCH=amd64 go build -o $@ $(MAIN)

build:
	$(info ========== Building for "$(GOOS)" into "$(BIN)/$(EXE)")
	@mkdir -p $(BIN)
	@$(GO) build -o $(BIN)/$(EXE) $(MAIN)

install: test build
	$(info ========== Installing "$(BIN)/$(EXE)" to "$(GOPATH)/$(BIN)/$(EXE)")
	@install $(BIN)/$(EXE) $(GOPATH)/$(BIN)/$(EXE)

clean:
	$(info ========== Cleaning)
	@rm -rf $(BIN)
	@rm -rf target
	@rm -f $(GOPATH)/bin/$(EXE)

.PHONY: all test target release-all clean