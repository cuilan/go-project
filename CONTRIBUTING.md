# 贡献指南

感谢您对本项目的关注和贡献！我们欢迎所有形式的贡献，包括但不限于代码、文档、问题报告和功能建议。

## 🤝 如何贡献

### 1. 报告问题

如果您发现了 bug 或有功能建议：

1. **搜索现有问题**: 首先检查 [Issues](../../issues) 中是否已有相似问题
2. **创建新问题**: 如果没有找到相关问题，请创建新的 Issue
3. **详细描述**: 包含以下信息：
   - 问题的详细描述
   - 重现步骤
   - 期望的行为
   - 实际的行为
   - 环境信息（操作系统、Go版本等）
   - 相关的日志或截图

### 2. 提交代码

#### 开发流程

1. **Fork 项目**: 点击页面右上角的 "Fork" 按钮
2. **克隆仓库**: 
   ```bash
   git clone https://github.com/cuilan/go-project.git
   cd go-project
   ```
3. **创建分支**: 
   ```bash
   git checkout -b feature/your-feature-name
   # 或
   git checkout -b fix/your-fix-name
   ```
4. **安装依赖**: 
   ```bash
   make mod-tidy
   make install-tools
   ```

#### 开发规范

在开始编码前，请阅读我们的 [代码规范](./CODE_STYLE.md)。

**关键要求**:
- 遵循 Go 语言编码规范
- 编写清晰的中文注释
- 添加必要的单元测试
- 确保代码通过所有检查

#### 代码检查

提交前请运行以下命令确保代码质量：

```bash
# 格式化代码
make format

# 代码质量检查
make lint

# 运行测试
make test

# 生成测试覆盖率报告
make test-cover
```

#### 提交规范

使用清晰的提交信息：

```bash
# 功能添加
feat: 添加用户注册功能

# 问题修复
fix: 修复用户登录时的密码验证问题

# 文档更新
docs: 更新API文档

# 代码重构
refactor: 重构用户服务层代码

# 测试相关
test: 添加用户服务单元测试

# 构建相关
build: 更新Makefile构建脚本
```

#### 创建 Pull Request

1. **推送分支**: 
   ```bash
   git push origin feature/your-feature-name
   ```
2. **创建 PR**: 在 GitHub 上创建 Pull Request
3. **PR 描述**: 包含以下内容：
   - 变更的详细描述
   - 相关的 Issue 链接
   - 测试结果截图（如适用）
   - 破坏性变更说明（如有）

### 3. 文档贡献

我们同样欢迎文档方面的贡献：

- 修正错误或不清楚的表述
- 添加使用示例
- 翻译文档
- 改进文档结构

## 🎯 开发环境设置

### 必需工具

- Go 1.21 或更高版本
- Git
- Make
- golangci-lint（通过 `make install-tools` 安装）

### 推荐工具

- VS Code 或 GoLand
- Docker（用于数据库测试）
- Postman 或类似的 API 测试工具

### 环境配置

1. **克隆项目**: 
   ```bash
   git clone https://github.com/cuilan/go-project.git
   cd go-project
   ```

2. **安装依赖**: 
   ```bash
   make mod-tidy
   make install-tools
   ```

3. **运行测试**: 
   ```bash
   make test
   ```

4. **启动开发服务器**: 
   ```bash
   make run server  # 启动 NetHTTP 服务器
   make run client  # 启动 Gin 服务器
   ```

## 🔍 代码审查流程

所有的 Pull Request 都会经过代码审查：

1. **自动检查**: CI/CD 会自动运行测试和代码检查
2. **人工审查**: 维护者会审查代码质量、设计和功能
3. **反馈处理**: 根据审查意见修改代码
4. **合并**: 通过审查后合并到主分支

### 审查标准

- **功能正确性**: 代码是否实现了预期功能
- **代码质量**: 是否遵循代码规范和最佳实践
- **测试覆盖**: 是否有足够的测试覆盖
- **文档完整性**: 是否更新了相关文档
- **向后兼容性**: 是否破坏了现有API

## 🏷️ 版本发布

我们使用语义版本控制 (Semantic Versioning)：

- **MAJOR**: 不兼容的API变更
- **MINOR**: 向后兼容的功能添加
- **PATCH**: 向后兼容的问题修复

## 📝 许可证

通过贡献代码，您同意您的贡献将在 [MIT License](./LICENSE) 下发布。

## 🙋‍♀️ 获得帮助

如果您在贡献过程中遇到问题：

1. **查看文档**: 检查 [README.md](./README.md) 和相关文档
2. **搜索 Issues**: 查看是否有相似问题
3. **创建 Discussion**: 在 GitHub Discussions 中提问
4. **联系维护者**: 通过 Issue 或 Email 联系项目维护者

## 🎉 致谢

感谢所有贡献者的努力！您的贡献让这个项目变得更好。

---

再次感谢您的贡献！🚀
