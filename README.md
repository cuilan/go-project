# Go 项目模板

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)](https://github.com/your-repo/go-project)

## 📖 文档导航

- [项目概述](#-项目概述) - 本文档
- [架构设计](./documentation/architecture/README.md) - 系统架构和设计原则  
- [快速开始](./documentation/getting-started/README.md) - 快速上手指南
- [配置管理](./documentation/configuration/README.md) - 配置文件详解

## 🎯 项目概述

这是一个功能强大、工程化的 Go 项目模板，旨在为新的 Go 应用提供一个坚实的起点。它整合了社区最佳实践，包括项目布局、跨平台构建、自动化测试、打包分发以及服务部署。

### ✨ 核心特性

- **🏗️ 标准化的项目布局**: 遵循 Go 社区广泛接受的 [`golang-standards/project-layout`](https://github.com/golang-standards/project-layout) 规范
- **🔧 强大的构建系统**: 基于 `Makefile` 和 Go 脚本，提供统一、跨平台的开发体验
- **🚀 一键式跨平台编译**: 通过 `make build` 轻松编译出 `Linux`, `Windows`, `macOS` 的可执行文件
- **📦 自动化打包与分发**: 使用 `make dist` 命令可一键生成包含配置文件和安装脚本的 `.zip` 分发包
- **🔍 内置代码质量工具**: 集成了 `fmt`, `vet`, `staticcheck` 等工具，通过 `make lint` 运行
- **🧪 测试与覆盖率报告**: 支持 `make test` 单元测试和 `make test-cover` 覆盖率报告生成
- **🎛️ 模块化架构**: 支持可插拔的模块系统，支持面向接口编程
- **🗄️ 多ORM支持**: 支持GORM和database/sql，可配置切换
- **📝 统一日志系统**: 基于slog的现代化日志系统
- **🔐 配置驱动**: 通过配置文件控制应用行为
- **🛡️ 完善的服务部署方案**:
  - **Linux**: 提供 `systemd` 服务安装脚本，支持开机自启、守护进程
  - **Windows**: 实现原生 Windows 服务，并提供安装和卸载脚本
- **🏷️ 版本信息注入**: 编译时自动将 Git `tag` 和 `commit hash` 注入到文件名中，便于版本追溯

## 📂 项目结构

```
.
├── bin/                      # (本地开发) 编译后的二进制文件 (被 .gitignore 忽略)
├── build/                    # 构建相关脚本和配置
│   ├── build_in_docker.sh  # Docker 构建脚本
│   └── Dockerfile           # Docker 镜像定义
├── cmd/                      # 项目主程序的入口
│   ├── client/              # 客户端应用 (Gin框架)
│   ├── server/              # 服务端应用 (net/http)
│   └── tui/                 # 终端用户界面应用
├── configs/                  # 配置文件模板
│   ├── app.yaml             # 基础配置文件
│   ├── app-dev.yaml         # 开发环境配置
│   └── examples/            # 配置示例
├── dist/                     # (打包分发) 生成的 .zip 分发包 (被 .gitignore 忽略)
├── docs/                     # API 文档 (Swagger 生成)
│   └── swagger/             # Swagger 文档目录
│       ├── client/          # Client 应用 API 文档
│       └── server/          # Server 应用 API 文档
├── documentation/            # 项目说明文档
│   ├── architecture/        # 架构设计文档
│   ├── configuration/       # 配置管理文档
│   └── getting-started/     # 快速开始指南
├── init/                     # 部署和初始化脚本
│   ├── linux/               # systemd (Linux) 服务相关脚本
│   └── windows/             # Windows 服务相关脚本
├── internal/                 # 项目内部私有代码
│   ├── api/                 # API 层 (框架无关)
│   │   ├── gin/            # Gin 框架实现
│   │   ├── nethttp/        # net/http 实现
│   │   ├── result.go       # 统一响应结构
│   │   └── types.go        # 通用模型定义
│   ├── conf/               # 配置管理
│   ├── logger/             # 日志系统
│   ├── module/             # 模块化系统
│   ├── orm/                # 数据访问层
│   │   ├── gorm/          # GORM 实现
│   │   ├── gosql/         # database/sql 实现
│   │   ├── models/        # 数据模型
│   │   └── repository/    # 仓储接口
│   ├── rdb/               # Redis 客户端
│   ├── service/           # 业务服务层
│   └── utils/             # 工具函数
├── pkg/                     # 可供外部使用的公共库代码
├── scripts/                  # 其他辅助脚本
├── sql/                     # 数据库脚本
├── test/                     # 测试相关
├── tools/                    # 项目工具
├── version/                  # 版本信息
├── .gitignore
├── go.mod
├── go.sum
├── Makefile                  # 项目自动化任务入口
└── README.md
```

## 🚀 快速开始

### 1. 克隆项目

```bash
git clone https://github.com/your-repo/go-project.git
cd go-project
make help
```

### 2. 安装开发工具

```bash
make install-tools
```

### 3. 本地构建

```bash
make build
```

### 4. 运行应用

```bash
# 运行 Client 应用 (Gin框架)
make run client

# 运行 Server 应用 (net/http)
make run server

# 运行 TUI 应用
make run tui

# 或者直接运行编译后的二进制文件
./bin/client --config-dir ./configs
./bin/server --config-dir ./configs
```

### 5. 生成API文档

```bash
# 为 Client 应用生成 Swagger 文档
make swag-init client

# 为 Server 应用生成 Swagger 文档
make swag-init server

# 查看生成的文档
# Client: docs/swagger/client/swagger.json
# Server: docs/swagger/server/swagger.json
```

### 6. 运行测试

```bash
# 运行所有测试
make test

# 生成覆盖率报告
make test-cover

# 代码质量检查
make lint
```

## 🛠️ 核心功能

### 🏗️ 框架无关的API设计

项目采用框架无关的API设计原则：

- **统一模型定义**: 通用的请求/响应模型定义在 `internal/api/types.go`
- **多框架支持**: 同时支持 Gin 和 net/http 框架
- **Swagger 文档**: 自动生成独立的API文档，避免冲突
- **类型安全**: 所有API接口都有完整的类型定义和验证

### 📚 完整的API文档系统

- **自动生成**: 基于代码注释自动生成Swagger文档
- **独立文档**: 每个应用生成独立的API文档，避免冲突
- **丰富注释**: 包含完整的请求/响应示例和错误码说明
- **在线预览**: 支持Swagger UI在线预览和测试

### 🔧 模块化系统

项目采用模块化设计，支持可插拔的组件：

- **日志模块**: 基于slog的现代化日志系统
- **数据库模块**: 支持GORM和database/sql
- **Redis模块**: Redis客户端支持
- **HTTP模块**: 支持Gin和net/http框架

### 🎯 面向接口编程

实现了完整的面向接口编程架构：

- **接口定义**: 在`internal/orm/repository/`中定义数据访问接口
- **具体实现**: 支持多种ORM实现（GORM、database/sql）
- **依赖注入**: 通过容器管理所有依赖
- **配置驱动**: 通过配置文件切换不同实现

### ⚙️ 智能配置管理

支持通过配置文件自动配置不同组件：

```yaml
# 数据库配置
gorm:
  driver: "mysql"
  dsn: "user:pass@tcp(localhost:3306)/dbname"
  config:
    enable_logger: true
    log_level: "info"

# Redis配置
redis:
  addr: "localhost:6379"
  pool_size: 10

# HTTP服务配置
gin:
  mode: "debug"
  port: 8080
```

## 🛠️ 开发指南

### 环境要求

- **Go**: 1.21 或更高版本
- **数据库**: MySQL 5.7+, PostgreSQL 12+, 或 SQLite 3.x
- **Redis**: 6.0+ (可选)
- **工具**: make, git, swag (Swagger文档生成)

### 开发流程

```bash
# 1. 克隆项目
git clone https://github.com/your-repo/go-project.git
cd go-project

# 2. 安装依赖
make install-tools
go mod download

# 3. 配置环境
cp configs/examples/gin.yaml configs/app-dev.yaml
# 编辑 configs/app-dev.yaml 配置数据库连接

# 4. 运行应用
make run client  # 或 make run server

# 5. 生成API文档
make swag-init client
make swag-init server

# 6. 运行测试
make test
make test-cover
```

### 添加新的API接口

1. **在 `internal/api/types.go` 中定义请求/响应模型**
2. **在对应框架包中实现处理器**
3. **添加完整的Swagger注释**
4. **编写单元测试**
5. **更新API文档**

## 📖 API 文档

项目提供完整的API文档系统：

### Swagger 文档

```bash
# 生成 Client 应用 API 文档 (Gin框架)
make swag-init client

# 生成 Server 应用 API 文档 (net/http)
make swag-init server

# 文档位置
docs/swagger/client/    # Client 应用文档
docs/swagger/server/    # Server 应用文档
```

### API 接口概览

#### Client 应用 (Gin框架)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/` | 根路径接口 |
| GET | `/ping` | Ping测试接口 |
| GET | `/health` | 健康检查接口 |
| POST | `/user/login` | 用户登录接口 |

#### Server 应用 (net/http)

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/health` | 健康检查接口 |
| POST | `/user/register` | 用户注册接口 |
| POST | `/user/login` | 用户登录接口 |

### API 模型定义

所有API模型都定义在 `internal/api/types.go` 中：

- `UserRegisterRequest`: 用户注册请求
- `UserLoginRequest`: 用户登录请求
- `User`: 用户信息响应
- `HealthResponse`: 健康检查响应
- `SuccessResponse`: 成功响应包装
- `ErrorResponse`: 错误响应包装

## 🚀 部署指南

### 本地部署

```bash
# 构建应用
make build

# 直接运行
./bin/client --config-dir ./configs
./bin/server --config-dir ./configs
```

### Docker 部署

```bash
# 构建Docker镜像
docker build -t go-project .

# 运行容器
docker run -d -p 8080:8080 -v $(pwd)/configs:/app/configs go-project
```

### 生产部署

```bash
# 构建生产版本
make build-all

# 创建分发包
make dist

# 生成的分发包位于 dist/ 目录
# 包含二进制文件、配置文件和安装脚本
```

#### Linux 服务部署

```bash
# 使用提供的安装脚本
cd dist/client_v1.0.0_linux_amd64/
sudo ./install.sh

# 启动服务
sudo systemctl start your-go-project
sudo systemctl enable your-go-project
```

#### Windows 服务部署

```cmd
# 使用提供的批处理脚本
cd dist\client_v1.0.0_windows_amd64\
install-service.bat

# 启动服务
net start YourGoProjectService
```

## 🧪 测试指南

### 单元测试

```bash
# 运行所有测试
make test

# 运行特定包的测试
go test ./internal/service/...
go test ./internal/orm/...
```

### 覆盖率测试

```bash
# 生成覆盖率报告
make test-cover

# 查看覆盖率报告
open coverage.html
```

### 集成测试

```bash
# 运行集成测试
go test ./test/...

# 使用测试配置
go test -config ./test/test.yaml ./test/...
```

### API 测试

```bash
# 启动应用
make run client

# 使用curl测试API
curl -X GET http://localhost:8080/health
curl -X POST http://localhost:8080/user/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"123456"}'
```

### 性能测试

```bash
# 运行基准测试
go test -bench=. ./internal/...

# 生成性能报告
go test -bench=. -cpuprofile=cpu.prof ./internal/...
go tool pprof cpu.prof
```

## 📋 Makefile 命令参考

### 构建命令

```bash
make build          # 构建当前平台版本
make build-all      # 构建所有平台版本
make dist           # 创建分发包
make release        # 创建发布包
```

### 运行命令

```bash
make run client     # 运行 Client 应用
make run server     # 运行 Server 应用
make run tui        # 运行 TUI 应用
make run-bin client # 运行编译后的二进制文件
```

### 文档命令

```bash
make swag-init client    # 生成 Client API 文档
make swag-init server    # 生成 Server API 文档
make swag-fmt client     # 格式化 Client Swagger 注释
make swag-fmt server     # 格式化 Server Swagger 注释
make swag-clean          # 清理生成的文档
make swag-check          # 检查 swag 工具
```

### 测试命令

```bash
make test           # 运行单元测试
make test-cover     # 生成覆盖率报告
make lint           # 代码质量检查
```

### 代码质量命令

```bash
make lint           # 运行代码检查
make lint-fix       # 运行代码检查并自动修复
make format         # 格式化代码
```

### 工具命令

```bash
make install-tools  # 安装开发工具 (包含 golangci-lint)
make mod-tidy       # 整理 go.mod
make mod-download   # 下载依赖
make mod-vendor     # 更新 vendor 目录
make clean          # 清理构建产物
make version        # 显示版本信息
make help           # 显示帮助信息
```

## 🔧 常用配置示例

### MySQL 配置

```yaml
gorm:
  driver: "mysql"
  dsn: "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
  max_open_conns: 20
  max_idle_conns: 10
  config:
    auto_migrate: true
    enable_logger: true
    log_level: "info"
```

### PostgreSQL 配置

```yaml
gorm:
  driver: "postgres"
  dsn: "host=localhost user=user password=password dbname=dbname port=5432 sslmode=disable"
  config:
    auto_migrate: true
    enable_logger: true
```

### Redis 配置

```yaml
redis:
  addr: "localhost:6379"
  password: ""
  db: 0
  pool_size: 10
  min_idle_conns: 5
```

### 日志配置

```yaml
log:
  console:
    enable_console: true
    level: "debug"
    console_format: "text"
  file:
    enable_file: true
    level: "info"
    path: "logs"
    filename: "app.log"
    file_max_size: 100
    file_max_backups: 3
    file_max_age: 7
```

## 📚 详细文档

- **[架构设计](./documentation/architecture/README.md)**: 深入了解系统架构和设计原则
- **[快速开始](./documentation/getting-started/README.md)**: 详细的快速上手指南  
- **[配置管理](./documentation/configuration/README.md)**: 配置文件详解和配置项说明

## 🤝 贡献指南

我们欢迎所有形式的贡献！请查看详细的 [贡献指南](./CONTRIBUTING.md)。

### 贡献方式

1. **报告问题**: 在 [Issues](../../issues) 中报告bug或提出改进建议
2. **提交代码**: Fork 项目，创建特性分支，提交 Pull Request
3. **完善文档**: 改进文档，修正错误，添加示例
4. **分享经验**: 在社区分享使用经验和最佳实践

### 开发规范

- 遵循项目 [代码规范](./CODE_STYLE.md)
- 使用 `make lint` 检查代码质量
- 编写单元测试，确保测试覆盖率
- 添加完整的代码注释和API文档
- 提交前运行 `make lint` 和 `make test`

更多详细信息请参阅 [CONTRIBUTING.md](./CONTRIBUTING.md)。

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](./LICENSE) 文件了解详情。

## 🆘 支持与反馈

如果您遇到问题或有疑问：

1. **查看文档**: 先查看相关文档是否有解决方案
2. **搜索Issues**: 搜索是否已有相似问题
3. **创建Issue**: 详细描述问题，包含环境信息和重现步骤
4. **社区讨论**: 参与社区讨论，分享经验

## 🏆 致谢

感谢所有贡献者和社区成员的支持！

特别感谢以下开源项目：

- [Gin](https://github.com/gin-gonic/gin) - HTTP Web框架
- [GORM](https://github.com/go-gorm/gorm) - Go ORM库
- [Viper](https://github.com/spf13/viper) - 配置管理
- [Swag](https://github.com/swaggo/swag) - Swagger文档生成

---

<div align="center">

**Happy Coding! 🎉** 

**⭐ 如果这个项目对你有帮助，请给个Star支持一下！**

Made with ❤️ by Echo·Green Cuilan!

</div>
