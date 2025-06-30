# Go 项目模板

这是一个功能强大、工程化的 Go 项目模板，旨在为新的 Go 应用提供一个坚实的起点。它整合了社区最佳实践，包括项目布局、跨平台构建、自动化测试、打包分发以及服务部署。

## ✨ 核心特性

- **标准化的项目布局**：遵循 Go 社区广泛接受的 [`golang-standards/project-layout`](https://github.com/golang-standards/project-layout) 规范。
- **强大的构建系统**：基于 `Makefile` 和 Go 脚本，提供统一、跨平台的开发体验。
- **一键式跨平台编译**：通过 `make build` 轻松编译出 `Linux`, `Windows`, `macOS` 的可执行文件。
- **自动化打包与分发**：使用 `make dist` 命令可一键生成包含配置文件和安装脚本的 `.zip` 分发包。
- **内置代码质量工具**：集成了 `fmt`, `vet`, `staticcheck` 等工具，通过 `make lint` 运行。
- **测试与覆盖率报告**：支持 `make test` 单元测试和 `make test-cover` 覆盖率报告生成。
- **完善的服务部署方案**：
    - **Linux**: 提供 `systemd` 服务安装脚本，支持开机自启、守护进程。
    - **Windows**: 实现原生 Windows 服务，并提供安装和卸载脚本。
- **版本信息注入**：编译时自动将 Git `tag` 和 `commit hash` 注入到文件名中，便于版本追溯。

## 📂 项目结构

```
.
├── bin/                      # (本地开发) 编译后的二进制文件 (被 .gitignore 忽略)
├── cmd/                      # 项目主程序的入口
│   └── your-go-project/
├── configs/                  # 配置文件模板
├── dist/                     # (打包分发) 生成的 .zip 分发包 (被 .gitignore 忽略)
├── init/                     # 部署和初始化脚本
│   ├── linux/                # systemd (Linux) 服务相关脚本
│   └── windows/              # Windows 服务相关脚本
├── internal/                 # 项目内部私有代码
├── pkg/                      # 可供外部使用的公共库代码
├── scripts/                  # 其他辅助脚本 (如：旧的交叉编译脚本)
├── tools/                    # 项目工具 (如：自定义的 Go 构建脚本)
├── .gitignore
├── go.mod
├── go.sum
├── Makefile                  # 项目自动化任务入口
└── README.md
```

## 🚀 快速开始

### 1. 克隆模板

```bash
git clone https://github.com/your-repo/go-project.git
cd go-project
```

### 2. (可选) 安装开发工具

为了使用 `make lint` 命令，需要安装 `staticcheck`。
```bash
make install-tools
```

### 3. 本地构建

在本地开发环境中编译二进制文件。产物将位于 `bin/` 目录下。
```bash
make build
```

### 4. 运行测试与检查

在提交代码前，确保所有检查都通过。
```bash
# 运行所有代码检查器
make lint

# 运行单元测试
make test

# 生成代码覆盖率报告 (生成 coverage.html)
make test-cover
```

### 5. 打包用于分发

当您准备发布一个新版本时，使用此命令。它会生成包含所有必要文件的 `.zip` 包，并存放在 `dist/` 目录下。
```bash
make dist
```
打包前，建议您先在本地打一个 `git tag`，例如 `git tag v1.0.0`，这样版本号就能正确地嵌入到包名中。

## 🛠️ Makefile 命令详解

- `make all`: 默认目标，等同于 `make build`。
- `make build`: 为所有在 `PLATFORMS` 变量中定义的目标平台进行交叉编译。
- `make dist`: 构建并为每个平台打包成一个可分发的 `.zip` 文件。
- `make lint`: 运行代码格式化、`go vet` 和 `staticcheck`。
- `make test`: 运行项目中的所有单元测试。
- `make test-cover`: 运行测试并生成 HTML 格式的代码覆盖率报告。
- `make clean`: 清理所有构建产物 (`bin/`, `dist/`) 和临时文件。
- `make vendor`: 整理 `go.mod` 并更新 `vendor` 目录。
- `make mod-tidy`: 仅整理 `go.mod` 文件。
- `make mod-download`: 下载依赖到本地模块缓存。
- `make install-tools`: 安装项目依赖的开发工具。
- `make help`: 显示所有可用的 `make` 命令及其说明。

## 部署指南

### Linux (Systemd)

1.  将对应平台的 `.zip` 包上传到服务器并解压。
2.  进入解压后的 `release/init/linux` 目录。
3.  **重要**: 修改 `install.sh` 脚本顶部的 `USER` 和 `GROUP` 变量，为您希望用来运行服务的用户。
4.  以 `root` 权限运行安装脚本：
    ```bash
    sudo ./install.sh
    ```
5.  脚本会自动完成用户创建、文件复制、服务注册和启动的所有工作。
6.  使用 `systemctl status your-go-project.service` 检查服务状态。

### Windows (服务)

1.  将对应 Windows 平台的 `.zip` 包解压。
2.  进入解压后的 `release/init/windows` 目录。
3.  **以管理员身份**运行 `install-service.bat`。
4.  脚本会自动注册并启动服务。
5.  您可以在 Windows 的"服务"管理单元中找到并管理此服务。卸载服务时，同样以管理员身份运行 `uninstall-service.bat`。
