# 系统架构设计

## 🏗️ 架构概述

本项目采用分层架构设计，遵循依赖倒置原则和面向接口编程理念，实现了高内聚、低耦合的系统架构。

## 📐 架构层次

```
┌─────────────────────────────────────────────────────────────┐
│                    应用层 (Application Layer)                │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │   Client    │  │   Server    │  │     TUI     │        │
│  │  (cmd/)     │  │  (cmd/)     │  │   (cmd/)    │        │
│  └─────────────┘  └─────────────┘  └─────────────┘        │
└─────────────────────────────────────────────────────────────┘
┌─────────────────────────────────────────────────────────────┐
│                    服务层 (Service Layer)                   │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │ UserService │  │ BookService │  │ OtherService│        │
│  │ (service/)  │  │ (service/)  │  │ (service/)  │        │
│  └─────────────┘  └─────────────┘  └─────────────┘        │
└─────────────────────────────────────────────────────────────┘
┌─────────────────────────────────────────────────────────────┐
│                  仓储层 (Repository Layer)                  │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │ IUserRepo   │  │ IBookRepo   │  │ IOtherRepo  │        │
│  │ (repository/│  │ (repository/│  │ (repository/│        │
│  └─────────────┘  └─────────────┘  └─────────────┘        │
└─────────────────────────────────────────────────────────────┘
┌─────────────────────────────────────────────────────────────┐
│                  数据访问层 (Data Access Layer)             │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │   GORM      │  │  database/  │  │   Other     │        │
│  │ (gorm/)     │  │    sql      │  │   ORM       │        │
│  │             │  │ (gosql/)    │  │             │        │
│  └─────────────┘  └─────────────┘  └─────────────┘        │
└─────────────────────────────────────────────────────────────┘
┌─────────────────────────────────────────────────────────────┐
│                  基础设施层 (Infrastructure Layer)           │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │   MySQL     │  │    Redis    │  │   Config    │        │
│  │ PostgreSQL  │  │   Cache     │  │   Logger    │        │
│  │   SQLite    │  │   Queue     │  │   Module    │        │
│  └─────────────┘  └─────────────┘  └─────────────┘        │
└─────────────────────────────────────────────────────────────┘
```

## 🎯 设计原则

### 1. 依赖倒置原则 (DIP)

- **高层模块不依赖低层模块**：应用层和服务层不直接依赖具体的数据访问实现
- **抽象不依赖细节**：通过接口定义数据访问契约
- **细节依赖抽象**：具体实现依赖接口定义

```go
// 接口定义（抽象）
type IUserRepository interface {
    Create(ctx context.Context, user *models.User) error
    GetByID(ctx context.Context, id int64) (*models.User, error)
    GetByUsername(ctx context.Context, username string) (*models.User, error)
    Delete(ctx context.Context, id int64) error
}

// 服务层依赖接口（高层依赖抽象）
type UserService struct {
    userRepo repository.IUserRepository
}

// 具体实现依赖接口（细节依赖抽象）
type gormUserRepository struct {
    db *gorm.DB
}

func (r *gormUserRepository) Create(ctx context.Context, user *models.User) error {
    return r.db.WithContext(ctx).Create(user).Error
}
```

### 2. 单一职责原则 (SRP)

每个模块和类都有明确的单一职责：

- **Repository**: 负责数据访问逻辑
- **Service**: 负责业务逻辑
- **Module**: 负责模块初始化和生命周期管理
- **Config**: 负责配置管理

### 3. 开闭原则 (OCP)

系统对扩展开放，对修改关闭：

```go
// 添加新的ORM实现无需修改现有代码
type sqlxUserRepository struct {
    db *sqlx.DB
}

func (r *sqlxUserRepository) Create(ctx context.Context, user *models.User) error {
    // 新的实现
}
```

## 🔧 核心组件

### 1. 模块化系统

```go
// Module 接口定义
type Module interface {
    Name() string
    Init(v *viper.Viper) error
    Close() error
}

// 模块注册和管理
var modules []Module

func Register(m ...Module) {
    for _, module := range m {
        modules = append(modules, module)
    }
}

func InitModules(v *viper.Viper) {
    for _, m := range modules {
        if v.IsSet(m.Name()) {
            m.Init(v)
        }
    }
}
```

### 2. 依赖注入容器

```go
// 容器管理所有依赖
type repoContainer struct {
    repoMap map[string]interface{}
    mu      sync.RWMutex
}

// 注册和获取依赖
func RegisterRepository(name string, repo interface{}) {
    getRepoContainer().register(name, repo)
}

func GetRepository(name string) (interface{}, bool) {
    return getRepoContainer().get(name)
}
```

### 3. 配置驱动

```go
// 配置结构体
type GormConfig struct {
    Driver string `mapstructure:"driver"`
    Dsn    string `mapstructure:"dsn"`
    Config struct {
        EnableLogger bool   `mapstructure:"enable_logger"`
        LogLevel     string `mapstructure:"log_level"`
    } `mapstructure:"config"`
}

// 条件初始化
func (m *gormModule) Init(v *viper.Viper) error {
    if !v.IsSet("gorm") {
        return nil // 跳过初始化
    }
    // 解析配置并初始化
}
```

## 🏛️ 架构模式

### 1. 分层架构 (Layered Architecture)

- **表现层**: cmd/ 目录下的应用入口
- **业务层**: internal/service/ 业务逻辑
- **数据访问层**: internal/orm/ 数据访问
- **基础设施层**: 数据库、缓存、配置等

### 2. 模块化架构 (Modular Architecture)

每个功能模块都是独立的，可以独立开发、测试和部署：

```
internal/
├── module/          # 模块化系统
├── orm/            # 数据访问模块
│   ├── gorm/      # GORM 实现
│   ├── gosql/     # database/sql 实现
│   └── repository/ # 仓储接口
├── rdb/            # Redis 模块
├── logger/         # 日志模块
└── http/           # HTTP 模块
```

### 3. 依赖注入模式 (Dependency Injection)

通过容器管理所有依赖，实现松耦合：

```go
// 自动注入
func autowired(db *gorm.DB) {
    repository.RegisterRepository(repository.UserRepositoryName, NewGormUserRepository(db))
    repository.RegisterRepository(repository.BookRepositoryName, NewGormBookRepository(db))
}

// 使用依赖
func NewUserService() *UserService {
    return &UserService{
        userRepo: repository.GetUserRepository(),
    }
}
```

## 🔄 数据流

### 1. 请求处理流程

```
HTTP Request → Router → Middleware → Controller → Service → Repository → Database
```

### 2. 配置加载流程

```
启动 → 加载配置文件 → 初始化日志 → 注册模块 → 条件初始化 → 启动应用
```

### 3. 依赖注入流程

```
模块注册 → 配置解析 → 条件初始化 → 依赖注入 → 容器管理 → 服务使用
```

## 🎨 设计模式

### 1. 工厂模式 (Factory Pattern)

```go
// 仓储工厂
func NewGormUserRepository(db *gorm.DB) repository.IUserRepository {
    return &gormUserRepository{db: db}
}

func NewUserRepository() repository.IUserRepository {
    return &gosqlUserRepository{db: GetDB()}
}
```

### 2. 单例模式 (Singleton Pattern)

```go
// 全局容器单例
var globalContainer *repoContainer
var once sync.Once

func getRepoContainer() *repoContainer {
    once.Do(func() {
        globalContainer = &repoContainer{
            repoMap: make(map[string]interface{}),
        }
    })
    return globalContainer
}
```

### 3. 策略模式 (Strategy Pattern)

```go
// 不同的ORM实现作为策略
type IUserRepository interface {
    Create(ctx context.Context, user *models.User) error
    GetByID(ctx context.Context, id int64) (*models.User, error)
}

// GORM 策略
type gormUserRepository struct{}

// database/sql 策略
type gosqlUserRepository struct{}
```

## 🔒 安全设计

### 1. 配置安全

- 敏感信息脱敏显示
- 环境变量支持
- 配置文件权限控制

### 2. 数据安全

- 参数化查询防止SQL注入
- 连接池管理
- 事务处理

### 3. 日志安全

- 敏感信息脱敏
- 日志级别控制
- 日志轮转

## 📈 扩展性设计

### 1. 水平扩展

- 无状态设计
- 配置外部化
- 服务发现支持

### 2. 垂直扩展

- 模块化设计
- 接口标准化
- 插件化架构

### 3. 功能扩展

- 新ORM实现
- 新业务模块
- 新存储引擎

## 🧪 可测试性

### 1. 单元测试

- 接口隔离
- Mock支持
- 依赖注入

### 2. 集成测试

- 数据库测试
- API测试
- 端到端测试

### 3. 性能测试

- 基准测试
- 压力测试
- 监控指标

## 📊 监控和可观测性

### 1. 日志系统

- 结构化日志
- 日志级别
- 日志聚合

### 2. 指标监控

- 性能指标
- 业务指标
- 系统指标

### 3. 链路追踪

- 请求追踪
- 依赖追踪
- 性能分析

---

这个架构设计确保了系统的可维护性、可扩展性和可测试性，为项目的长期发展奠定了坚实的基础。 