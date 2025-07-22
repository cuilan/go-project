# Copyright The Authors.

# ************************************************************************************
# 公共配置变量
# Common configuration variables
# ************************************************************************************

PKG=github.com/cuilan/go-project

# 应用名称列表，应与 cmd/ 目录下的子目录名一致
# Application names list, should match the subdirectory names under cmd/
APPS := client server

# ************************************************************************************

# Get current project absolute path
ROOTDIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
# Main package path function
# 定义一个函数来获取 main package 路径
define get_main_package
$(ROOTDIR)/cmd/$(1)
endef

# ====================================================================================
# Platform detection and configuration
# ====================================================================================

ifneq "$(strip $(shell command -v $(GO) 2>/dev/null))" ""
	GOOS ?= $(shell $(GO) env GOOS)
	GOARCH ?= $(shell $(GO) env GOARCH)
else
	ifeq ($(GOOS),)
		# approximate GOOS for the platform if we don't have Go and GOOS isn't
		# set. We leave GOARCH unset, so that may need to be fixed.
		ifeq ($(OS),Windows_NT)
			GOOS = windows
		else
			UNAME_S := $(shell uname -s)
			ifeq ($(UNAME_S),Linux)
				GOOS = linux
			endif
			ifeq ($(UNAME_S),Darwin)
				GOOS = darwin
			endif
			ifeq ($(UNAME_S),FreeBSD)
				GOOS = freebsd
			endif
		endif
	else
		GOOS ?= $$GOOS
		GOARCH ?= $$GOARCH
	endif
endif

#include platform specific makefile
include Makefile.$(GOOS)

# ====================================================================================
# Platform definitions
# ====================================================================================

# MacOS
DARWIN_AMD64 := darwin/amd64
DARWIN_ARM64 := darwin/arm64
DARWIN := $(DARWIN_AMD64) $(DARWIN_ARM64)

# Linux
LINUX_386 := linux/386
LINUX_AMD64 := linux/amd64
LINUX_ARM := linux/arm
LINUX_ARM64 := linux/arm64
LINUX := $(LINUX_386) $(LINUX_AMD64) $(LINUX_ARM) $(LINUX_ARM64)

# Windows
WINDOWS_386 := windows/386
WINDOWS_AMD64 := windows/amd64
WINDOWS_ARM := windows/arm
WINDOWS_ARM64 := windows/arm64
WINDOWS := $(WINDOWS_386) $(WINDOWS_AMD64) $(WINDOWS_ARM) $(WINDOWS_ARM64)

# Platform combinations
ALL_32 := $(LINUX_386) $(WINDOWS_386)
ALL_64 := $(DARWIN_AMD64) $(LINUX_AMD64) $(WINDOWS_AMD64)
ALL_ARM := $(LINUX_ARM) $(WINDOWS_ARM)
ALL_ARM64 := $(DARWIN_ARM64) $(LINUX_ARM64) $(WINDOWS_ARM64)
ALL_PLATFORMS := $(ALL_32) $(ALL_64) $(ALL_ARM) $(ALL_ARM64)

# All platforms
# PLATFORMS := $(ALL_PLATFORMS)
PLATFORMS := $(ALL_64)
# PLATFORMS := $(WINDOWS_AMD64)

# ====================================================================================
# Git related variables
# ====================================================================================

# Get semantic version from git tag
GIT_VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null | sed 's/^v//' || echo "0.0.0")
# Get git branch
GIT_REVISION := $(shell git rev-parse --abbrev-ref HEAD)
# Get git commit hash (short format)
# GIT_COMMIT := $(shell git rev-parse --short HEAD)
GIT_COMMIT :=

# ====================================================================================
# Go related commands
# ====================================================================================

GO := go
GOFMT := gofmt

# Current platform
CURRENT_PLATFORM := $(shell go env GOOS)/$(shell go env GOARCH)

# ====================================================================================
# Common targets
# ====================================================================================

version: ## Show version information
	@echo "$(WHALE) $@"
	@echo "ROOTDIR: $(ROOTDIR)"
	@echo "APP_NAME: $(APP_NAME)"
	@echo "Current platform: $(CURRENT_PLATFORM)"
	@echo "git tag: $(GIT_VERSION)"
	@echo "git branch: $(GIT_REVISION)"
	@echo "git commit: $(GIT_COMMIT)"
	@$(GO) run $(MAIN_PACKAGE) --version

mod-tidy: ## Tidy go.mod file
	@echo "$(WHALE) $@"
	@echo "$(RUN)  Tidying go.mod..."
	@$(GO) mod tidy
	@echo "$(OK)  go.mod tidied."

mod-download: ## Download modules to local cache
	@echo "$(WHALE) $@"
	@echo "$(RUN)  Downloading dependencies..."

install-tools: ## Install code checking development tools (staticcheck)
	@echo "$(WHALE) $@"
	@echo "$(RUN)  Installing development tools (staticcheck)..."
	@$(GO) install honnef.co/go/tools/cmd/staticcheck@latest
	@echo "$(OK)  Tools installation completed."

# ====================================================================================
# Test related targets
# ====================================================================================

test: ## Run all unit tests
	@echo "$(WHALE) $@"
	@echo "$(RUN)  Running unit tests..."
	@$(GO) test -v ./...
	@echo "$(OK)  Tests completed."

test-cover: ## Run tests and generate HTML coverage report
	@echo "$(WHALE) $@"
	@echo "$(RUN)  Generating code coverage report..."
	@$(GO) test -v -cover -coverprofile=coverage.out ./...
	@$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "$(OK)  Coverage report generated: coverage.html"

# ====================================================================================
# Code quality checks
# ====================================================================================

lint: ## Run all code checkers
	@echo "$(WHALE) $@"
	@echo "$(RUN)  Running code formatting check..."
	@$(GOFMT) -l -w .
	@echo "$(RUN)  Running go vet..."
	@$(GO) vet ./...
	@echo "$(RUN)  Running staticcheck..."
	@staticcheck ./...
	@echo "$(OK)  All code checks completed."

# ====================================================================================
# Dependency management
# ====================================================================================

mod-vendor: mod-tidy ## Update vendor directory
	@echo "$(WHALE) $@"
	@echo "$(RUN)  Updating vendor directory..."
	@$(GO) mod vendor
	@echo "$(OK)  Vendor directory updated."

# ====================================================================================
# Core build and distribution
# ====================================================================================

build: ## Build for current platform
	@echo "$(WHALE) $@"
	@echo "$(RUN)  Building for current platform ($(CURRENT_PLATFORM)) (version: $(GIT_VERSION)-$(GIT_COMMIT))..."
	@PLATFORMS="$(CURRENT_PLATFORM)" \
	COMMANDS="$(APPS)" \
	VERSION="$(GIT_VERSION)" \
	COMMIT="$(GIT_COMMIT)" \
	$(GO) run ./tools/build.go
	@echo "$(OK)  Build completed. Output located in bin/ directory."

build-all: ## Build for all platforms
	@echo "$(WHALE) $@"
	@echo "$(RUN)  Starting build (version: $(GIT_VERSION)-$(GIT_COMMIT))..."
	@PLATFORMS="$(PLATFORMS)" \
	COMMANDS="$(APPS)" \
	VERSION="$(GIT_VERSION)" \
	COMMIT="$(GIT_COMMIT)" \
	$(GO) run ./tools/build.go
	@echo "$(OK)  Build completed. Output located in bin/ directory."

dist: build-all ## Build and package distribution
	@echo "$(WHALE) $@"
	@mkdir -p dist
	@echo "$(RUN)  Creating distribution packages..."
	@for app in $(APPS); do $(call cross_build,$$app); done
	@echo "$(OK)  Distribution packages created successfully in dist/ directory."

# Define cross-build for all platforms
define cross_build
	APP_NAME=$(1); \
	echo "$(RUN)  binary: $$APP_NAME"; \
	for p in $(PLATFORMS); do \
		echo " $(DO)  $(DIST) Processing platform: $$p for app: $$APP_NAME"; \
		GOOS=`echo $$p | cut -d'/' -f1`; \
		GOARCH=`echo $$p | cut -d'/' -f2`; \
		VERSIONED_APP_NAME=$${APP_NAME}_$(GIT_VERSION)_$${GOOS}_$${GOARCH}; \
		BINARY_NAME=$$VERSIONED_APP_NAME; \
		EXT=""; \
		if [ "$${GOOS}" = "windows" ]; then \
			BINARY_NAME=$$VERSIONED_APP_NAME.exe; \
			EXT=".exe"; \
		fi; \
		echo "    Creating $$VERSIONED_APP_NAME directory..."; \
		mkdir -p dist/$$VERSIONED_APP_NAME; \
		mkdir -p dist/$$VERSIONED_APP_NAME/bin; \
		echo "    Copying binary file..."; \
		cp bin/$$VERSIONED_APP_NAME$$EXT dist/$$VERSIONED_APP_NAME/bin/$$BINARY_NAME; \
		echo "    Copying configs directory..."; \
		cp -r configs dist/$$VERSIONED_APP_NAME/configs; \
		echo "    Copying init directory..."; \
		if [ "$${GOOS}" = "linux" ]; then \
			cp -r init/linux/* dist/$$VERSIONED_APP_NAME; \
		fi; \
		if [ "$${GOOS}" = "windows" ]; then \
			cp -r init/windows/* dist/$$VERSIONED_APP_NAME; \
		fi; \
		echo "    $(OK) Platform $$p completed"; \
	done
endef

release: dist ## Create release zip packages
	@echo "$(WHALE) $@"
	@echo "$(RUN)  Creating release zip packages..."
	@for app in $(APPS); do $(call release_zip,$$app); done
	@echo "$(OK)  Release packages created successfully in dist/ directory."

# Define release_zip for all platforms
define release_zip
	APP_NAME=$(1); \
	for p in $(PLATFORMS); do \
		echo " $(DO)  $(DIST) Release platform: $$p for app: $$APP_NAME"; \
		GOOS=`echo $$p | cut -d'/' -f1`; \
		GOARCH=`echo $$p | cut -d'/' -f2`; \
		VERSIONED_APP_NAME=$${APP_NAME}_$(GIT_VERSION)_$${GOOS}_$${GOARCH}; \
		echo "    Creating zip file..."; \
		zip -r dist/$$VERSIONED_APP_NAME.zip dist/$$VERSIONED_APP_NAME > /dev/null; \
	done
endef

# ====================================================================================
# Run
# ====================================================================================

run: ## Run application (e.g., make run client)
	@APP=$(word 2,$(MAKECMDGOALS)); \
	if [ -z "$$APP" ]; then \
		echo "Error: Please specify the application (e.g., make run client)"; \
		exit 1; \
	fi; \
	$(GO) run $(call get_main_package,$$APP) --config-dir ./configs
%::
	@:

run-bin: build ## Run executable file (e.g., make run-bin client)
	@APP=$(word 2,$(MAKECMDGOALS)); \
	if [ -z "$$APP" ]; then \
		echo "Error: Please specify the application (e.g., make run-bin client)"; \
		exit 1; \
	fi; \
	echo "$(WHALE) $@"; \
	echo "$(RUN)  Running executable file..."; \
	./bin/$$APP* --config-dir ./configs; \
	echo "$(OK)  Run completed."

# ====================================================================================
# Clean
# ====================================================================================

clean: ## Clean build artifacts and temporary files
	@echo "$(WHALE) $@"
	@echo "$(RUN)  Cleaning..."
	@rm -rf bin
	@rm -rf dist
	@rm -f coverage.html
	@rm -f coverage.out
	@rm -rf logs
	@echo "$(OK)  Cleanup completed."

# ====================================================================================
# Help information
# ====================================================================================

help: ## Show this help information
	@echo "Usage: make [target]"
	@echo ""
	@echo "Available targets:"

	@echo ""
	@echo "Build targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^(build|build-all|dist|release):.*?## / {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

	@echo ""
	@echo "Test targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^(test|test-cover):.*?## / {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

	@echo ""
	@echo "Run targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^(run|run-bin):.*?## / {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

	@echo ""
	@echo "Common targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^(version|mod-tidy|mod-download|install-tools|lint|mod-vendor|clean|help):.*?## / {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

.PHONY:  version mod-tidy mod-download install-tools test test-cover lint mod-vendor run build build-all dist release clean help


.DEFAULT_GOAL := help
