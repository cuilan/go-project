# Copyright The Authors.

WHALE = "☁️"

PKG=github.com/cuilan/your-go-project-name
COMMANDS=your-go-project-name

# all
#PLATFORMS=darwin/amd64 darwin/arm64 linux/386 linux/amd64 linux/arm linux/arm64 windows/amd64
#PLATFORMS=linux/amd64 windows/amd64 darwin/amd64
PLATFORMS=linux/amd64 windows/amd64

######################### 不建议修改的变量 #########################

# Go command to use for build
GO ?= go
INSTALL ?= install

# 获取当前工程绝对路径
# MAKEFILE_LIST makefile 预定义变量
ROOTDIR=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))

# Base path used to install.
# The files will be installed under `$(DESTDIR)/$(PREFIX)`.
# The convention of `DESTDIR` was changed in containerd v1.6.
PREFIX        ?= /usr/local
BINDIR        ?= $(PREFIX)/bin
DATADIR       ?= $(PREFIX)/share
DOCDIR        ?= $(DATADIR)/doc
MANDIR        ?= $(DATADIR)/man

TEST_IMAGE_LIST ?=

RELEASE=release

# Used to populate variables in version package.
VERSION ?= $(shell git describe --match 'v[0-9]*' --dirty='.m' --always)
REVISION ?= $(shell git rev-parse HEAD)$(shell if ! git diff --no-ext-diff --quiet --exit-code; then echo .m; fi)

GO_TAGS=$(if $(GO_BUILDTAGS),-tags "$(strip $(GO_BUILDTAGS))",)

# Project packages.
PACKAGES=$(shell $(GO) list ${GO_TAGS} ./internal/... | grep -v /vendor/ | grep -v /version)

#Replaces ":" (*nix), ";" (windows) with newline for easy parsing
GOPATHS=$(shell $(GO) env GOPATH | tr ":" "\n" | tr ";" "\n")

#include platform specific makefile
#-include Makefile.$(GOOS)

# Flags passed to `go test`
TESTFLAGS ?= $(TESTFLAGS_RACE) $(EXTRA_TESTFLAGS)
TESTFLAGS_PARALLEL ?= 8

# Use this to replace `go test` with, for instance, `gotestsum`
GOTEST ?= $(GO) test

OUTPUTDIR = $(join $(ROOTDIR), _output)

#------------------------------------------------------

.PHONY: clean all AUTHORS build test
.DEFAULT: default

# Forcibly set the default goal to all, in case an include above brought in a rule definition.
.DEFAULT_GOAL := all

all: clean

check: ## run all linters
	@echo "$(WHALE) $@"
	GOGC=75 gofmt -l -w .

AUTHORS: .mailmap .git/HEAD
	git log --format='%aN <%aE>' | sort -fu > $@

######################### build #########################

define CROSS_COMPILE
@for pl in ${PLATFORMS}; do \
	export GOOS=$$(echo $${pl} | cut -d'/' -f1); \
	export GOARCH=$$(echo $${pl} | cut -d'/' -f2); \
	export CGO_ENABLED=0; \
	export TARGET=${RELEASE}/${COMMANDS}; \
	if [ "$${GOOS}" == "windows" ]; then \
		export CGO_ENABLED=1; \
		export TARGET=$${TARGET}_$${GOOS}_$${GOARCH}.exe; \
	else \
		export TARGET=$${TARGET}_$${GOOS}_$${GOARCH}; \
	fi; \
	echo "Building for $${GOOS}/$${GOARCH} => $${TARGET}"; \
	$(GO) build -mod=vendor -o $${TARGET} || exit 1; \
done
endef

build: clean ## build the go packages
	@echo "$(WHALE) $@"
	@mkdir ${RELEASE}
	@echo "Starting build..."
	$(call CROSS_COMPILE)
	@echo "Build completed!"

######################### install #########################

######################### test #########################

test: ## run tests, except integration tests and tests that require root
	@echo "$(WHALE) $@"
	@$(GOTEST) ${TESTFLAGS} ${PACKAGES}

######################### clean #########################

clean: ## clean up binaries, releases and logs
	@echo "$(WHALE) $@"
	@rm -rf ${RELEASE}
	@echo "Cleaning build ${RELEASE}..."
	@rm -rf $(OUTPUTDIR)
	@rm -rf logs/*
	@echo "Cleaning logs..."

clean-test: ## clean up debris from previously failed tests
	@echo "$(WHALE) $@"

######################### vendor #########################

remove-replace:
	@echo "$(WHALE) $@"

vendor: ## ensure all the go.mod/go.sum files are up-to-date including vendor/ directory
	@echo "$(WHALE) $@"
	@$(GO) mod tidy
	@$(GO) mod vendor
	@$(GO) mod verify

verify-vendor: ## verify if all the go.mod/go.sum files are up-to-date
	@echo "$(WHALE) $@"
	@$(GO) mod tidy
	@$(GO) mod verify

clean-vendor: remove-replace vendor

######################### help #########################

help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort