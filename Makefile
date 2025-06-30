# Copyright The Authors.

WHALE = "☁️"

PKG=github.com/cuilan/go-project
COMMANDS=your-go-project

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
VERSION ?= $(shell git describe --tags --abbrev=0 2>/dev/null | sed 's/^v//' || echo "0.0.0")
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

.PHONY: all build clean help lint test test-cover install-tools vendor dist

# ====================================================================================
# 项目变量
# ====================================================================================

# 应用名称，应与 cmd/ 目录下的子目录名一致
APP_NAME := your-go-project
# 从 git tag 获取语义版本号 (例如: 1.2.3)，如果无 tag 则默认为 0.0.0
GIT_VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null | sed 's/^v//' || echo "0.0.0")
# 获取 git commit hash (短格式)
GIT_COMMIT := $(shell git rev-parse --short HEAD)
# Go 相关命令
GO := go
GOFMT := gofmt
GOTEST := gotest
# 默认编译平台列表 (空格分隔)
PLATFORMS ?= "linux/amd64 windows/amd64 darwin/amd64 darwin/arm64"

# ====================================================================================
# 核心构建与分发
# ====================================================================================

all: build ## 构建所有二进制文件到 bin/ 目录

build: ## 使用 Go 脚本进行跨平台构建
	@echo "▶️  正在开始构建 (版本: $(GIT_VERSION)-$(GIT_COMMIT))..."
	@PLATFORMS="$(PLATFORMS)" \
	COMMANDS="$(APP_NAME)" \
	VERSION="$(GIT_VERSION)" \
	COMMIT="$(GIT_COMMIT)" \
	$(GO) run ./tools/build.go
	@echo "✅  构建完成。产物位于 bin/ 目录。"

dist: build ## 构建并打包成 ZIP 可分发文件
	@echo "▶️  正在创建分发包..."
	@mkdir -p dist
	@# 为每个平台创建对应的压缩包
	@for p in $(PLATFORMS); do \
		GOOS=`echo $$p | cut -d'/' -f1`; \
		GOARCH=`echo $$p | cut -d'/' -f2`; \
		VERSIONED_APP_NAME=$(APP_NAME)_$(GIT_VERSION)_$(GIT_COMMIT)_$${GOOS}_$${GOARCH}; \
		ZIP_NAME=$(APP_NAME)_$(GIT_VERSION)_$(GIT_COMMIT)_$${GOOS}_$${GOARCH}.zip; \
		BINARY_NAME=$(APP_NAME); \
		if [ "$${GOOS}" = "windows" ]; then \
			BINARY_NAME=$(APP_NAME).exe; \
			VERSIONED_APP_NAME=$(APP_NAME)_$(GIT_VERSION)_$(GIT_COMMIT)_$${GOOS}_$${GOARCH}.exe; \
		fi; \
		echo "  📦 正在打包 $$p..."; \
		TEMP_DIR=dist/staging; \
		rm -rf $$TEMP_DIR; \
		mkdir -p $$TEMP_DIR/release; \
		cp bin/$$VERSIONED_APP_NAME $$TEMP_DIR/release/$$BINARY_NAME; \
		cp -r configs $$TEMP_DIR/release/; \
		cp -r init $$TEMP_DIR/release/; \
		(cd $$TEMP_DIR && zip -r ../$$ZIP_NAME release > /dev/null); \
		rm -rf $$TEMP_DIR; \
	done
	@echo "✅  ZIP 分发包创建完成，位于 dist/ 目录。"

# ====================================================================================
# 代码质量与测试
# ====================================================================================

lint: ## 运行所有代码检查器 (fmt, vet, staticcheck)
	@echo "▶️  正在运行代码格式化检查..."
	@$(GOFMT) -l -w .
	@echo "▶️  正在运行 go vet..."
	@$(GO) vet ./...
	@echo "▶️  正在运行 staticcheck..."
	@staticcheck ./...
	@echo "✅  所有代码检查完成。"

test: ## 运行所有单元测试 (不包括代码覆盖率)
	@echo "▶️  正在运行单元测试..."
	@$(GO) test -v ./...
	@echo "✅  测试完成。"

test-cover: ## 运行测试并生成 HTML 覆盖率报告
	@echo "▶️  正在生成代码覆盖率报告..."
	@$(GO) test -v -cover -coverprofile=coverage.out ./...
	@$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "✅  覆盖率报告已生成: coverage.html"

# ====================================================================================
# 依赖与工具管理
# ====================================================================================

vendor: mod-tidy ## 更新 vendor 目录
	@echo "▶️  正在更新 vendor 目录..."
	@$(GO) mod vendor
	@echo "✅  Vendor 目录已更新。"

mod-tidy: ## 整理 go.mod 文件
	@echo "▶️  正在整理 go.mod..."
	@$(GO) mod tidy
	@echo "✅  go.mod 已整理。"

mod-download: ## 下载模块到本地缓存
	@echo "▶️  正在下载依赖..."
	@$(GO) mod download
	@echo "✅  依赖下载完成。"

# ====================================================================================
# 清理与帮助
# ====================================================================================

# 根据操作系统定义清理命令
ifeq ($(OS),Windows_NT)
    CLEAN_CMD = rmdir /s /q
else
    CLEAN_CMD = rm -rf
endif

clean: ## 清理构建产物和临时文件
	@echo "▶️  正在清理..."
	@$(CLEAN_CMD) bin
	@$(CLEAN_CMD) dist
	@$(CLEAN_CMD) coverage.html coverage.out
	@echo "✅  清理完成。"

install-tools: ## 安装代码检查等开发工具
	@echo "▶️  正在安装开发工具 (staticcheck)..."
	@$(GO) install honnef.co/go/tools/cmd/staticcheck@latest
	@echo "✅  工具安装完成。"

help: ## 显示此帮助信息
	@echo "用法: make [目标]"
	@echo ""
	@echo "可用目标:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help