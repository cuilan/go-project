# Go 项目模板 - 完整文档

## 📖 文档导航

- [项目概述](./README.md) - 本文档
- [架构设计](./docs/architecture/README.md) - 系统架构和设计原则
- [快速开始](./docs/getting-started/README.md) - 快速上手指南
- [开发指南](./docs/development/README.md) - 开发环境搭建和最佳实践
- [配置管理](./docs/configuration/README.md) - 配置文件详解
- [部署指南](./docs/deployment/README.md) - 生产环境部署
- [API 文档](./docs/api/README.md) - 接口文档
- [测试指南](./docs/testing/README.md) - 测试策略和工具
- [故障排除](./docs/troubleshooting/README.md) - 常见问题和解决方案

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
├── cmd/                      # 项目主程序的入口
│   ├── client/              # 客户端应用
│   ├── server/              # 服务端应用
│   └── tui/                 # 终端用户界面应用
├── configs/                  # 配置文件模板
│   └── examples/            # 配置示例
├── dist/                     # (打包分发) 生成的 .zip 分发包 (被 .gitignore 忽略)
├── docs/                     # 项目文档
│   ├── architecture/        # 架构设计文档
│   ├── getting-started/     # 快速开始指南
│   ├── development/         # 开发指南
│   ├── configuration/       # 配置管理文档
│   ├── deployment/          # 部署指南
│   ├── api/                 # API 文档
│   ├── testing/             # 测试指南
│   └── troubleshooting/     # 故障排除
├── init/                     # 部署和初始化脚本
│   ├── linux/               # systemd (Linux) 服务相关脚本
│   └── windows/             # Windows 服务相关脚本
├── internal/                 # 项目内部私有代码
│   ├── conf/               # 配置管理
│   ├── http/               # HTTP 服务
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
# 使用默认配置
./bin/client

# 指定配置文件目录
./bin/client --config-dir ./configs

# 指定环境配置
./bin/client --config-dir ./configs --profile dev
```

### 5. 运行测试

```bash
# 运行所有测试
make test

# 生成覆盖率报告
make test-cover

# 代码质量检查
make lint
```

## 🛠️ 核心功能

### 模块化系统

项目采用模块化设计，支持可插拔的组件：

- **日志模块**: 基于slog的现代化日志系统
- **数据库模块**: 支持GORM和database/sql
- **Redis模块**: Redis客户端支持
- **HTTP模块**: Gin Web框架支持

### 面向接口编程

实现了完整的面向接口编程架构：

- **接口定义**: 在`internal/orm/repository/`中定义数据访问接口
- **具体实现**: 支持多种ORM实现（GORM、database/sql）
- **依赖注入**: 通过容器管理所有依赖
- **配置驱动**: 通过配置文件切换不同实现

### 自动配置

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
```

## 📚 详细文档

- **[架构设计](./architecture/README.md)**: 深入了解系统架构和设计原则
- **[快速开始](./getting-started/README.md)**: 详细的快速上手指南
- **[开发指南](./development/README.md)**: 开发环境搭建和最佳实践
- **[配置管理](./configuration/README.md)**: 配置文件详解和配置项说明
- **[部署指南](./deployment/README.md)**: 生产环境部署指南
- **[API 文档](./api/README.md)**: 接口文档和使用示例
- **[测试指南](./testing/README.md)**: 测试策略和工具使用
- **[故障排除](./troubleshooting/README.md)**: 常见问题和解决方案

## 🤝 贡献指南

欢迎贡献代码！请查看 [贡献指南](./CONTRIBUTING.md) 了解如何参与项目开发。

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](../LICENSE) 文件了解详情。

## 🆘 支持

如果您遇到问题或有疑问，请：

1. 查看 [故障排除](./troubleshooting/README.md) 文档
2. 搜索 [Issues](../../issues)
3. 创建新的 [Issue](../../issues/new)

---

**Happy Coding! 🎉** 