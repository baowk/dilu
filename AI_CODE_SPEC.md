# Dilu AI 代码生成规范

<p align="center">
  🇨🇳 中文版本 • <a href="./AI_CODE_SPEC_en.md">🇺🇸 English</a>
</p>

> **⚠️ 重要说明**：本文档是**给 AI 助手遵循的代码生成规范**，AI 必须严格遵守。

---

## 🎯 核心原则（必须遵守）

1. **优先使用gen生成代码** - 保持代码结构整洁
2. **使用TDD开发模式** - 测试先行
3. **生成的代码禁止修改** - gen 命令生成的 Repository/Model/Query 层代码严禁手动修改
4. **业务逻辑在 Service层** - 所有业务逻辑必须在 Service 层实现
5. **职责边界清晰** - Repository 管数据、Service 管业务、Api 管接口
6. **安全迭代** - 修改数据库 → 重新生成 → Service 层适配
7. **使用结构化日志** - 使用 log/slog
8. **不重复造轮子** - 标准库优先，三方库使用需谨慎
9. **swagger文档** - 接口采用swagger做文档，注释清晰，使用 go generate 生成
10. **遵守 golang 代码规范** - 版本 go 1.26 +

---

## 📋 基础规范

### 项目结构
```
dilu/
├── cmd/                    # 命令行工具
├── common/                 # 公共组件
├── internal/               # 核心业务
│   ├── bootstrap/         # 初始化
│   ├── tools/             # 代码生成器
│   │   ├── apis/          # API 接口
│   │   ├── service/       # 业务逻辑
│   │   ├── router/        # 路由配置
│   │   └── repository/    # 数据访问 ⭐
│   │       ├── model/     # Model 层
│   │       └── query/     # Query 层
│   └── sys/               # 此模块由gen生成，默认sys
├── resources/              # 配置文件
├── docs/                   # 用户文档
├── dev-docs/               # 开发文档
│   └── releases/          # 发布文档
└── tests/                  # 测试代码
```

### 分层职责
| 层级 | 职责 | 文件示例 |
|------|------|---------|
| **Repository** | 数据持久化 | `*.gen.go` (禁止修改,字段改变从新生成，默认会覆盖) |
| **Service** | 业务逻辑 | `sys_user_service.go` |
| **Api** | HTTP 接口封装 | `sys_user_api.go` |
| **Router** | 路由配置 | `sys_user_router.go` |

### 命名规范
- **Model**: 大驼峰，如 `SysUser`
- **方法**: 小驼峰，如 `GetUserByID`
- **变量**: 小驼峰，如 `userName`
- **常量**: 全大写，如 `TableNameSysUser`
- **包名**: 全小写，如 `model`, `query`, `service`

### 错误处理
```go
// ✅ 正确：立即返回错误
if err != nil {
    return nil, err
}

// ✅ 正确：包装错误上下文
if err != nil {
    return fmt.Errorf("get user failed: %w", err)
}

// ❌ 错误：忽略错误
getUser(id) // 未处理 err
```

### 日志规范
```go
// ✅ 正确：结构化日志（使用 slog）
slog.Debug("get user success", "user", user)   // Debug 级别
slog.Info("get user success", "user", user)    // Info 级别
slog.Warn("get user timeout", "user", user)    // Warn 级别
slog.Error("get user failed", "error", err)    // Error 级别

// ❌ 错误：无意义日志
log.Println("error occurred")
```

---

## 💻 代码生成流程（多库架构）

Dilu 支持多数据库架构，可以通过 `-d` 参数指定不同的数据库。

### 步骤 1：准备数据库
```bash
# SQLite（推荐开发使用）
go run scripts/init_sqlite.go

# MySQL（生产环境）
mysql -u root -p -e "CREATE DATABASE dilu_db"
```

### 步骤 2：运行生成器

#### 示例 1：生成系统库的 sys_user 表（默认数据库）
```bash
go run main.go gen \
  -c resources/config.sqlite.yaml \
  -t sys_user
```

#### 示例 2：生成 notice 库的 message 表（指定数据库）
```bash
go run main.go gen \
  -c resources/config.sqlite.yaml \
  -d notice \
  -t message
```

**参数说明**：
- `-c`：配置文件路径
- `-d`：数据库名称（可选，默认为 `sys`）
- `-t`：表名
- `-f`：覆盖已存在的文件（可选）
- `-p`：包名（可选，默认 根 -d 参数一致）

### 步骤 3：验证生成
```
# 检查生成的文件
ls internal/sys/repository/model/sys_user.gen.go
ls internal/sys/repository/query/sys_user.gen.go
```

### 步骤 4：扩展实现
``go
// ✅ 创建扩展文件（不要修改 .gen.go）
// internal/sys/repository/model/sys_user_extend.go
package model

type SysUserWithRoles struct {
    SysUser
    Roles []SysRole `json:"roles"`
}
```

---

## 🌐 API 开发规范

### RESTful 路由设计
```go
// QueryPage - POST /api/v1/sys/user/page     - 分页查询列表
// Get      - GET  /api/v1/sys/user/{id}      - 获取单个记录
// Create   - POST /api/v1/sys/user/create    - 创建记录
// Update   - POST /api/v1/sys/user/update    - 更新记录
// Del      - POST /api/v1/sys/user/del       - 删除记录
```

### 统一响应格式
``go
// 成功响应（带数据）
e.Ok(c, data)
// 返回：{"code": 200, "message": "success", "data": {...}}

// 成功响应（无数据）
e.Ok(c)
// 返回：{"code": 200, "message": "success"}

// 分页响应
e.Page(c, list, total, page, pageSize)
// 返回：{"code": 200, "data": {"list": [...], "total": 100}}

// 错误响应
e.Error(c, err)
// 返回：{"code": 500, "message": "error message"}
```

---

## 🗄️ 数据库设计

### 表设计规范
```sql
-- 用户表
CREATE TABLE `sys_user` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '用户 ID',
  `username` varchar(100) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_username` (`username`)
) ENGINE=InnoDB COMMENT='用户表';
```

### GORM Model 规范 由 GORM 生成
```go
// 示例模型结构
type SysUser struct {
    ID       int    `gorm:"primarykey" json:"id"`
    Username string `gorm:"uniqueIndex" json:"username"`
    Password string `json:"password"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

---

### 简洁优雅原则 ⭐

#### ✅ 应该做的
- **简单直接**：代码逻辑清晰，一目了然
- **职责单一**：每个函数只做一件事
- **命名准确**：见名知意，无需注释
- **错误处理**：立即返回，不嵌套
- **代码复用**：提取公共逻辑，避免重复

#### ❌ 不应该做的
- **过度设计**：不要为了炫技而使用复杂模式
- **炫技编程**：避免晦涩的语法糖和技巧
- **过早优化**：先让代码工作，再考虑性能
- **滥用抽象**：不要为了抽象而抽象
- **魔法数字**：使用有意义的常量

---

## ✅ TDD 编程实践

### 测试先行流程
```
1. 编写失败的测试 → 2. 编写最小实现 → 3. 重构优化
```

### 测试覆盖要求
- **Service层**: 核心业务逻辑覆盖率 ≥ 80%
- **Api 层**: 主要接口测试覆盖率 ≥ 70%

---

## 📝 最佳实践清单

### ✅ 应该做的
- [x] 使用 gen 命令生成代码
- [x] 在 Service 层实现业务逻辑
- [x] 为 Service层编写单元测试
- [x] 使用相对路径引用文件
- [x] 遵循 RESTful API 规范
- [x] 及时更新 Swagger 文档（使用 go generate）
- [x] 使用结构化日志（log/slog）
- [x] 妥善处理所有错误

### ❌ 禁止做的
- [ ] 修改 `.gen.go` 文件
- [ ] 在 Api 层实现业务逻辑
- [ ] 忽略错误返回值
- [ ] 使用绝对路径
- [ ] 写死配置信息
- [ ] 缺少 Swagger 注解
- [ ] 不写单元测试
- [ ] 直接打印日志
