# Go 项目代码规范

## 📖 概述

本文档定义了项目的Go代码规范，旨在保证代码质量、可读性和团队协作效率。所有贡献者都应该遵循这些规范。

---

## 🛠️ 代码检查工具

### GolangCI-Lint

项目使用 [GolangCI-Lint](https://golangci-lint.run/) 作为主要的代码检查工具。

#### 安装

```bash
# macOS
brew install golangci-lint

# Linux
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2

# Windows
# 下载二进制文件或使用 Scoop
scoop install golangci-lint
```

#### 使用

```bash
# 检查所有文件
make lint

# 或直接使用 golangci-lint
golangci-lint run

# 检查指定目录
golangci-lint run ./internal/...

# 自动修复可修复的问题
golangci-lint run --fix
```

---

## 📝 代码风格规范

### 1. 命名规范

#### 包名 (Package Names)
- 使用小写字母
- 简短且有意义
- 避免下划线和大小写混合

```go
// ✅ 好的包名
package user
package config
package httputil

// ❌ 不好的包名
package userManager
package user_service
package HTTPUtil
```

#### 变量名 (Variable Names)
- 使用驼峰命名 (camelCase)
- 局部变量可以使用短名称
- 全局变量使用完整描述性名称

```go
// ✅ 好的变量名
var userCount int
var maxRetryAttempts int
func getUserByID(id int64) {}

// 局部变量可以简短
for i, v := range items {
    // i, v 是可接受的
}

// ❌ 不好的变量名
var user_count int
var MaxRetryAttempts int  // 不应该导出的变量
```

#### 函数名 (Function Names)
- 使用驼峰命名
- 导出函数使用大写字母开头
- 私有函数使用小写字母开头

```go
// ✅ 导出函数
func GetUserByID(id int64) (*User, error) {}
func CreateUser(user *User) error {}

// ✅ 私有函数
func validateUserInput(user *User) error {}
func hashPassword(password string) string {}
```

#### 常量名 (Constants)
- 使用驼峰命名或全大写+下划线
- 导出常量使用大写字母开头

```go
// ✅ 常量命名
const (
    DefaultTimeout = 30 * time.Second
    MaxRetryCount  = 3
)

const (
    STATUS_ACTIVE   = "active"
    STATUS_INACTIVE = "inactive"
)
```

### 2. 代码格式

#### 缩进和空格
- 使用 `gofmt` 自动格式化
- 使用 tab 缩进
- 运算符前后加空格

```go
// ✅ 正确格式
if user != nil && user.ID > 0 {
    return user.Name
}

// ❌ 错误格式
if user!=nil&&user.ID>0{
    return user.Name
}
```

#### 行长度
- 建议每行不超过 120 字符
- 长行应该合理换行

```go
// ✅ 合理换行
user, err := userService.CreateUser(
    ctx,
    &User{
        Name:  "John Doe",
        Email: "john@example.com",
        Age:   30,
    },
)

// ❌ 过长的行
user, err := userService.CreateUser(ctx, &User{Name: "John Doe", Email: "john@example.com", Age: 30, CreatedAt: time.Now()})
```

### 3. 函数设计

#### 函数长度
- 函数应该保持简短，建议不超过 50 行
- 复杂函数应该拆分为多个小函数

#### 参数数量
- 函数参数不应超过 5 个
- 参数过多时考虑使用结构体

```go
// ✅ 参数适量
func CreateUser(ctx context.Context, name, email string, age int) error {}

// ✅ 使用结构体
type CreateUserRequest struct {
    Name  string
    Email string
    Age   int
    Phone string
    Address string
}

func CreateUser(ctx context.Context, req *CreateUserRequest) error {}
```

#### 返回值
- 优先返回错误作为最后一个返回值
- 避免使用裸返回 (naked return)

```go
// ✅ 标准返回
func GetUser(id int64) (*User, error) {
    // ...
    return user, nil
}

// ❌ 裸返回 (避免使用)
func GetUser(id int64) (user *User, err error) {
    // ...
    return  // 裸返回
}
```

---

## 🏗️ 架构规范

### 1. 包结构

```
internal/
├── api/           # API层，框架无关
├── service/       # 业务逻辑层
├── orm/          # 数据访问层
├── conf/         # 配置管理
├── logger/       # 日志系统
└── utils/        # 工具函数
```

### 2. 依赖规则

- 上层可以依赖下层，下层不能依赖上层
- 同层之间通过接口交互
- 避免循环依赖

```go
// ✅ 正确的依赖方向
// API层 -> Service层 -> Repository层

// ❌ 错误的依赖
// Repository层 -> Service层 (违反分层原则)
```

### 3. 接口设计

- 接口应该小而专注
- 优先使用组合而非继承
- 接口命名使用 -er 后缀或 I 前缀

```go
// ✅ 好的接口设计
type UserRepository interface {
    Create(ctx context.Context, user *User) error
    GetByID(ctx context.Context, id int64) (*User, error)
    Update(ctx context.Context, user *User) error
    Delete(ctx context.Context, id int64) error
}

type UserReader interface {
    GetByID(ctx context.Context, id int64) (*User, error)
    GetByEmail(ctx context.Context, email string) (*User, error)
}

type UserWriter interface {
    Create(ctx context.Context, user *User) error
    Update(ctx context.Context, user *User) error
    Delete(ctx context.Context, id int64) error
}
```

---

## 🛡️ 错误处理

### 1. 错误创建

```go
// ✅ 使用 errors.New 或 fmt.Errorf
import "errors"
import "fmt"

var ErrUserNotFound = errors.New("user not found")

func GetUser(id int64) (*User, error) {
    if id <= 0 {
        return nil, fmt.Errorf("invalid user id: %d", id)
    }
    // ...
}
```

### 2. 错误包装

```go
// ✅ 使用 fmt.Errorf 和 %w 包装错误
func GetUserProfile(id int64) (*UserProfile, error) {
    user, err := userRepo.GetByID(ctx, id)
    if err != nil {
        return nil, fmt.Errorf("failed to get user profile: %w", err)
    }
    // ...
}
```

### 3. 错误检查

```go
// ✅ 使用 errors.Is 和 errors.As
if errors.Is(err, ErrUserNotFound) {
    // 处理用户未找到
}

var validationErr *ValidationError
if errors.As(err, &validationErr) {
    // 处理验证错误
}
```

---

## 📋 注释规范

### 1. 包注释

```go
// Package user 提供用户管理相关的功能。
//
// 这个包实现了用户的增删改查操作，包括用户认证、权限管理等功能。
// 支持多种数据库后端和缓存策略。
package user
```

### 2. 函数注释

```go
// GetUserByID 根据用户ID获取用户信息。
//
// 参数:
//   - ctx: 请求上下文，用于超时控制和取消操作
//   - id: 用户ID，必须大于0
//
// 返回值:
//   - *User: 用户信息，如果用户不存在则返回nil
//   - error: 错误信息，如果操作成功则返回nil
//
// 示例:
//   user, err := GetUserByID(ctx, 123)
//   if err != nil {
//       log.Printf("获取用户失败: %v", err)
//       return
//   }
func GetUserByID(ctx context.Context, id int64) (*User, error) {
    // 实现...
}
```

### 3. 结构体注释

```go
// User 表示系统中的用户实体。
//
// 包含用户的基本信息和状态信息。
type User struct {
    ID        int64     `json:"id" db:"id"`                    // 用户唯一标识
    Name      string    `json:"name" db:"name"`                // 用户姓名
    Email     string    `json:"email" db:"email"`              // 邮箱地址，用于登录
    CreatedAt time.Time `json:"created_at" db:"created_at"`    // 创建时间
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`    // 更新时间
}
```

---

## 🧪 测试规范

### 1. 测试文件命名

```
user.go       -> user_test.go
service.go    -> service_test.go
```

### 2. 测试函数命名

```go
// 单元测试
func TestGetUserByID(t *testing.T) {}
func TestCreateUser_Success(t *testing.T) {}
func TestCreateUser_InvalidInput(t *testing.T) {}

// 基准测试
func BenchmarkGetUserByID(b *testing.B) {}

// 示例测试
func ExampleGetUserByID() {}
```

### 3. 测试结构

```go
func TestCreateUser(t *testing.T) {
    tests := []struct {
        name    string
        input   *User
        want    error
        wantErr bool
    }{
        {
            name: "valid user",
            input: &User{
                Name:  "John Doe",
                Email: "john@example.com",
            },
            want:    nil,
            wantErr: false,
        },
        {
            name: "invalid email",
            input: &User{
                Name:  "John Doe",
                Email: "invalid-email",
            },
            want:    ErrInvalidEmail,
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := CreateUser(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !errors.Is(err, tt.want) {
                t.Errorf("CreateUser() error = %v, want %v", err, tt.want)
            }
        })
    }
}
```

---

## 🔒 安全规范

### 1. 输入验证

```go
// ✅ 验证输入参数
func CreateUser(user *User) error {
    if user == nil {
        return errors.New("user cannot be nil")
    }
    if user.Email == "" {
        return errors.New("email is required")
    }
    if !isValidEmail(user.Email) {
        return errors.New("invalid email format")
    }
    // ...
}
```

### 2. 敏感信息处理

```go
// ✅ 不在日志中输出敏感信息
func LoginUser(email, password string) error {
    log.Printf("用户登录尝试: email=%s", email) // ✅ 不记录密码
    // 不要: log.Printf("登录: email=%s, password=%s", email, password)
    
    // 使用哈希比较密码
    if !comparePassword(hashedPassword, password) {
        return errors.New("invalid credentials")
    }
    return nil
}
```

---

## 📊 性能规范

### 1. 避免不必要的内存分配

```go
// ✅ 预分配切片容量
users := make([]*User, 0, expectedCount)

// ✅ 重用字符串构建器
var builder strings.Builder
builder.Grow(expectedSize) // 预分配容量
```

### 2. 使用上下文控制

```go
// ✅ 传递上下文用于超时控制
func GetUsers(ctx context.Context) ([]*User, error) {
    select {
    case <-ctx.Done():
        return nil, ctx.Err()
    default:
        // 执行操作
    }
}
```

---

## 🔧 工具配置

### 1. VS Code 设置

```json
{
    "go.formatTool": "goimports",
    "go.lintTool": "golangci-lint",
    "go.lintOnSave": "package",
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
        "source.organizeImports": true
    }
}
```

### 2. Git Hooks

在 `.git/hooks/pre-commit` 中添加：

```bash
#!/bin/sh
# 运行代码检查
make lint
if [ $? -ne 0 ]; then
    echo "代码检查失败，请修复后再提交"
    exit 1
fi

# 运行测试
make test
if [ $? -ne 0 ]; then
    echo "测试失败，请修复后再提交"
    exit 1
fi
```

---

## ✅ 检查清单

提交代码前请确保：

- [ ] 代码通过 `make lint` 检查
- [ ] 代码通过 `make test` 测试
- [ ] 添加了必要的注释和文档
- [ ] 遵循了命名规范
- [ ] 错误处理完整
- [ ] 没有硬编码的敏感信息
- [ ] 测试覆盖了主要功能

---

## 📚 参考资源

- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [GolangCI-Lint Documentation](https://golangci-lint.run/)
- [Go语言规范](https://golang.org/ref/spec)

---

**注意**: 这些规范是指导性的，在特殊情况下可以适当调整，但需要在代码审查中说明原因。
