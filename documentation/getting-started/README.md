# 快速开始指南

## 🚀 快速上手

本指南将帮助您在几分钟内启动并运行这个 Go 项目。

## 📋 前置要求

### 1. 系统要求

- **操作系统**: Linux, macOS, Windows
- **Go 版本**: 1.21 或更高版本
- **内存**: 至少 2GB 可用内存
- **磁盘空间**: 至少 1GB 可用空间

### 2. 必需软件

#### Go 环境

```bash
# 检查 Go 版本
go version

# 应该显示类似输出：
# go version go1.21.0 darwin/amd64
```

#### 数据库 (可选)

- **MySQL**: 5.7 或更高版本
- **PostgreSQL**: 12 或更高版本
- **SQLite**: 3.x

#### Redis (可选)

- **Redis**: 6.0 或更高版本

## 🔧 环境准备

### 1. 克隆项目

```bash
# 克隆项目到本地
git clone https://github.com/your-repo/go-project.git

# 进入项目目录
cd go-project
```

### 2. 安装依赖

```bash
# 下载 Go 模块依赖
go mod download

# 安装开发工具
make install-tools
```

### 3. 配置环境

#### 基础配置

项目提供了默认配置，可以直接运行：

```bash
# 使用默认配置运行
make build
./bin/client
```

#### 自定义配置 (可选)

如果需要自定义配置，可以修改配置文件：

```bash
# 复制配置文件模板
cp configs/examples/gorm.yaml configs/app-dev.yaml

# 编辑配置文件
vim configs/app-dev.yaml
```

## 🎯 快速体验

### 1. 最小化运行

```bash
# 构建项目
make build

# 运行客户端应用
./bin/client

# 应该看到类似输出：
# ================================================
# |            Your Go Project Service           |
# -------------------------------------------------
# > OS: darwin Arch: amd64
# > Go Version: go1.21.0
# > Project Version: v1.0.0
# > Config Path: ./configs
# ================================================
```

### 2. 带数据库运行

如果您有数据库环境，可以配置数据库连接：

```yaml
# configs/app-dev.yaml
gorm:
  driver: "mysql"
  dsn: "root:password@tcp(localhost:3306)/test_db"
  config:
    auto_migrate: true
    enable_logger: true
    log_level: "info"
```

然后运行：

```bash
./bin/client --config-dir ./configs
```

### 3. 带 Redis 运行

如果您有 Redis 环境，可以配置 Redis 连接：

```yaml
# configs/app-dev.yaml
redis:
  addr: "localhost:6379"
  pool_size: 10
```

## 📊 验证运行状态

### 1. 检查日志输出

正常运行时，您应该看到：

```
2024/01/01 12:00:00 INFO ================================================
2024/01/01 12:00:00 INFO |            Your Go Project Service           |
2024/01/01 12:00:00 INFO -------------------------------------------------
2024/01/01 12:00:00 INFO > OS: darwin Arch: amd64
2024/01/01 12:00:00 INFO > Go Version: go1.21.0
2024/01/01 12:00:00 INFO > Project Version: v1.0.0
2024/01/01 12:00:00 INFO > Config Path: ./configs
2024/01/01 12:00:00 INFO ================================================
2024/01/01 12:00:00 INFO application config completed name=demo profile=dev
2024/01/01 12:00:00 INFO initializing pluggable modules... count=2
2024/01/01 12:00:00 INFO initializing module=redis
2024/01/01 12:00:00 INFO redis connected successfully
2024/01/01 12:00:00 INFO initializing module=gorm
2024/01/01 12:00:00 INFO gorm database connected successfully driver=mysql max_open_conns=20 max_idle_conns=10
2024/01/01 12:00:00 INFO all pluggable modules initialized
```

### 2. 检查进程状态

```bash
# 检查进程是否运行
ps aux | grep client

# 检查端口占用 (如果有HTTP服务)
netstat -tlnp | grep 8080
```

### 3. 检查生成的文件

```bash
# 检查日志文件
ls -la logs/

# 检查数据库文件 (如果使用SQLite)
ls -la *.db
```

## 🔍 故障排除

### 1. 常见问题

#### 问题: `go: command not found`

**解决方案**:
```bash
# 安装 Go
# macOS
brew install go

# Ubuntu/Debian
sudo apt-get install golang-go

# Windows
# 下载并安装 Go: https://golang.org/dl/
```

#### 问题: `make: command not found`

**解决方案**:
```bash
# macOS
brew install make

# Ubuntu/Debian
sudo apt-get install make

# Windows
# 使用 Git Bash 或 WSL
```

#### 问题: 数据库连接失败

**解决方案**:
```bash
# 检查数据库服务状态
# MySQL
sudo systemctl status mysql

# PostgreSQL
sudo systemctl status postgresql

# 检查连接信息
mysql -u root -p -h localhost
```

#### 问题: Redis 连接失败

**解决方案**:
```bash
# 检查 Redis 服务状态
sudo systemctl status redis

# 测试 Redis 连接
redis-cli ping
```

### 2. 调试模式

启用详细日志输出：

```bash
# 设置日志级别为 debug
export LOG_LEVEL=debug

# 运行应用
./bin/client --config-dir ./configs
```

### 3. 获取帮助

```bash
# 查看帮助信息
./bin/client --help

# 查看版本信息
./bin/client --version
```

## 📚 下一步

### 1. 学习项目结构

- 查看 [架构设计](../architecture/README.md) 了解系统架构
- 查看 [配置管理](../configuration/README.md) 了解配置系统
- 查看 [开发指南](../development/README.md) 了解开发流程

### 2. 自定义开发

- 添加新的业务模块
- 实现新的数据访问层
- 扩展配置选项

### 3. 部署到生产环境

- 查看 [部署指南](../deployment/README.md)
- 配置生产环境参数
- 设置监控和日志

## 🎉 恭喜！

您已经成功运行了这个 Go 项目！现在您可以：

1. **探索代码**: 查看 `internal/` 目录了解项目结构
2. **修改配置**: 编辑 `configs/` 目录下的配置文件
3. **添加功能**: 在 `cmd/` 目录下添加新的应用入口
4. **扩展模块**: 在 `internal/` 目录下添加新的业务模块

如果您遇到任何问题，请查看 [故障排除](../troubleshooting/README.md) 文档或创建 Issue。

---

**Happy Coding! 🚀** 