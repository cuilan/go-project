# Windows 平台构建指南

本文档介绍如何在 Windows 平台上使用 `Makefile.windows` 进行项目构建和管理。

## 概述

`Makefile.windows` 是专门为 Windows 平台设计的构建脚本，提供了以下功能：

- **跨平台构建**：支持 Windows 386/AMD64/ARM/ARM64 架构
- **Windows 服务管理**：安装、启动、停止、重启 Windows 服务
- **Windows 特定优化**：使用 Windows 兼容的命令和路径
- **无乱码输出**：所有输出信息使用英文，避免字符编码问题

## 前置要求

1. **安装 Go**：确保已安装 Go 1.18 或更高版本
2. **安装 Make**：推荐使用 [GNU Make for Windows](http://gnuwin32.sourceforge.net/packages/make.htm)
3. **PowerShell**：用于压缩文件操作（Windows 10+ 默认支持）

## 使用方法

### 基本构建命令

```cmd
# 显示帮助信息
make -f Makefile.windows help

# 构建当前 Windows 平台的二进制文件
make -f Makefile.windows build

# 构建所有 Windows 平台的二进制文件
make -f Makefile.windows build-all

# 运行应用程序
make -f Makefile.windows run

# 运行已构建的可执行文件
make -f Makefile.windows run-bin
```

### Windows 服务管理

```cmd
# 安装 Windows 服务（需要管理员权限）
make -f Makefile.windows install-service

# 启动服务
make -f Makefile.windows start-service

# 停止服务
make -f Makefile.windows stop-service

# 重启服务
make -f Makefile.windows restart-service

# 卸载服务
make -f Makefile.windows uninstall-service
```

### 代码质量检查

```cmd
# 运行代码格式化检查
make -f Makefile.windows lint

# 运行单元测试
make -f Makefile.windows test

# 生成代码覆盖率报告
make -f Makefile.windows test-cover
```

### 依赖管理

```cmd
# 整理 go.mod 文件
make -f Makefile.windows mod-tidy

# 更新 vendor 目录
make -f Makefile.windows mod-vendor

# 下载依赖
make -f Makefile.windows mod-download
```

### 分发打包

```cmd
# 构建并创建 ZIP 分发包
make -f Makefile.windows dist
```

### 清理操作

```cmd
# 清理构建产物
make -f Makefile.windows clean

# 安装开发工具
make -f Makefile.windows install-tools
```

## 配置说明

### 应用名称配置

在 `Makefile.windows` 中修改 `APP_NAME` 变量：

```makefile
APP_NAME := your-app
```

### 服务配置

可以自定义 Windows 服务的名称和描述：

```makefile
SERVICE_NAME := your-app
SERVICE_DISPLAY_NAME := "Your Go Project Service"
SERVICE_DESCRIPTION := "Go Project Service for Windows"
```

### 安装路径配置

Windows 安装路径配置：

```makefile
PREFIX        ?= C:/Program Files/$(APP_NAME)
BINDIR        ?= $(PREFIX)/bin
```

## Windows 特定功能

### 1. 服务管理

`Makefile.windows` 提供了完整的 Windows 服务管理功能：

- **自动启动**：服务安装后会自动启动
- **错误处理**：包含服务不存在或未运行的错误处理
- **权限检查**：安装服务需要管理员权限

### 2. PowerShell 集成

使用 PowerShell 进行文件压缩操作：

```cmd
powershell -command "Compress-Archive -Path release -DestinationPath ..\filename.zip"
```

### 3. Windows 路径处理

- 使用反斜杠 `\` 作为路径分隔符
- 使用 `%~dp0` 获取脚本所在目录
- 使用 `if exist` 进行文件存在性检查

## 故障排除

### 常见问题

1. **权限不足**
   ```
   错误：sc create 失败
   解决：以管理员身份运行命令提示符
   ```

2. **Make 命令未找到**
   ```
   错误：'make' 不是内部或外部命令
   解决：安装 GNU Make for Windows
   ```

3. **Go 命令未找到**
   ```
   错误：'go' 不是内部或外部命令
   解决：安装 Go 并添加到 PATH 环境变量
   ```

4. **PowerShell 执行策略限制**
   ```
   错误：无法加载文件，因为在此系统上禁止运行脚本
   解决：以管理员身份运行 PowerShell 并执行：
        Set-ExecutionPolicy RemoteSigned
   ```

### 调试技巧

1. **查看详细输出**：移除 `@` 符号查看详细命令执行
2. **检查环境变量**：使用 `echo %PATH%` 检查环境变量
3. **验证文件存在**：使用 `dir` 命令检查文件是否存在

## 与主 Makefile 的区别

| 功能 | 主 Makefile | Makefile.windows |
|------|-------------|------------------|
| 跨平台支持 | 支持 Linux/macOS/Windows | 仅支持 Windows |
| 字符编码 | 可能乱码 | 无乱码 |
| 服务管理 | 无 | 完整的 Windows 服务管理 |
| 命令语法 | Unix 风格 | Windows 风格 |
| 路径分隔符 | 正斜杠 `/` | 反斜杠 `\` |

## 最佳实践

1. **开发环境**：使用主 Makefile 进行日常开发
2. **Windows 部署**：使用 Makefile.windows 进行 Windows 特定构建
3. **服务部署**：使用 Makefile.windows 的服务管理功能
4. **CI/CD**：在 Windows 构建环境中使用 Makefile.windows

## 贡献

如需修改 Windows 特定的构建逻辑，请编辑 `Makefile.windows` 文件。建议在修改前先在测试环境中验证功能。 