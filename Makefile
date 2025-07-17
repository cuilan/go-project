# Copyright The Authors.

WHALE = "☁️"
RUN = "▶️"
OK = "✅"
INFO = "ℹ️"
WARNING = "⚠️"
ERROR = "❌"

# 包含公共定义
include Makefile.common

# 主包文件
MAIN_FILE := $(MAIN_PACKAGE)/main.go

# ====================================================================================
# 平台特定配置
# ====================================================================================

# 全平台
# PLATFORMS := $(ALL_PLATFORMS)
PLATFORMS := $(ALL_64)

# ====================================================================================
# 安装路径，不建议修改
# ====================================================================================

ifeq ($(OS),Linux)
	PREFIX        ?= /usr/local
	BINDIR        ?= $(PREFIX)/bin
	DATADIR       ?= $(PREFIX)/share
	DOCDIR        ?= $(DATADIR)/doc
	MANDIR        ?= $(DATADIR)/man
else ifeq ($(OS),Windows)
	PREFIX        ?= C:/Program Files/$(APP_NAME)
	BINDIR        ?= $(PREFIX)/bin
	DATADIR       ?= $(PREFIX)/share
	DOCDIR        ?= $(DATADIR)/doc
	MANDIR        ?= $(DATADIR)/man
endif

# ====================================================================================

.PHONY: all build build-all dist run-bin clean

# ====================================================================================
# 核心构建与分发
# ====================================================================================

all: build-all ## 构建所有平台（PLATFORMS）的二进制文件到 bin/ 目录

build: ## 构建当前平台的二进制文件
	@echo "$(WHALE) $@"
	@echo "$(RUN)  Building for current platform ($(CURRENT_PLATFORM)) (version: $(GIT_VERSION)-$(GIT_COMMIT))..."
	@PLATFORMS="$(CURRENT_PLATFORM)" \
	COMMANDS="$(APP_NAME)" \
	VERSION="$(GIT_VERSION)" \
	COMMIT="$(GIT_COMMIT)" \
	$(GO) run ./tools/build.go
	@echo "$(OK)  Build completed. Output located in bin/ directory."

build-all: ## 使用 Go 脚本进行跨平台构建
	@echo "$(WHALE) $@"
	@echo "$(RUN)  Starting build (version: $(GIT_VERSION)-$(GIT_COMMIT))..."
	@PLATFORMS="$(PLATFORMS)" \
	COMMANDS="$(APP_NAME)" \
	VERSION="$(GIT_VERSION)" \
	COMMIT="$(GIT_COMMIT)" \
	$(GO) run ./tools/build.go
	@echo "$(OK)  Build completed. Output located in bin/ directory."

dist: build-all ## 构建并打包成 ZIP 可分发文件
	@echo "$(WHALE) $@"
	@echo "$(RUN)  Creating distribution packages..."
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
		echo "  [PKG] Packaging $$p..."; \
		TEMP_DIR=dist/staging; \
		rm -rf $$TEMP_DIR; \
		mkdir -p $$TEMP_DIR/release; \
		cp bin/$$VERSIONED_APP_NAME $$TEMP_DIR/release/$$BINARY_NAME; \
		cp -r configs $$TEMP_DIR/release/; \
		cp -r init $$TEMP_DIR/release/; \
		(cd $$TEMP_DIR && zip -r ../$$ZIP_NAME release > /dev/null); \
		rm -rf $$TEMP_DIR; \
	done
	@echo "$(OK)  ZIP distribution packages created successfully in dist/ directory."

# ====================================================================================
# 运行
# ====================================================================================

run-bin: build ## 运行可执行文件
	@echo "$(WHALE) $@"
	@echo "$(RUN)  Running executable file..."
	./bin/$(APP_NAME)* --config-dir ./configs
	@echo "$(OK)  Run completed."

# ====================================================================================
# 清理
# ====================================================================================

CLEAN_CMD = rm -rf

clean: ## 清理构建产物和临时文件
	@echo "$(WHALE) $@"
	@echo "$(RUN)  Cleaning..."
	@$(CLEAN_CMD) bin
	@$(CLEAN_CMD) dist
	@$(CLEAN_CMD) coverage.html coverage.out
	@$(CLEAN_CMD) logs
	@echo "$(OK)  Cleanup completed."