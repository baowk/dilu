# 发布说明 v1.2.1 - AI 代码规范体系完善

**发布日期**: 2026-03-14  
**版本类型**: 优化更新  
**兼容性**: ✅ 向后兼容（无破坏性变更）

---

## 🎯 核心更新

本次更新聚焦于建立完整的 AI 代码规范体系和多库架构支持：

1. **AI 代码规范文档** - 新增 `AI_CODE_SPEC.md` 系列文档（266 行中文版 + 273 行英文版）
2. **多库架构支持** - 代码生成器支持 `-d` 参数指定不同数据库
3. **TDD 开发实践** - 完整集成测试驱动开发规范
4. **代码简洁优雅** - 新增代码优雅原则和最佳实践清单

---

## ✨ 主要特性

### 1. AI 代码规范文档体系 🤖

**新增文件**：
- [`AI_CODE_SPEC.md`](./AI_CODE_SPEC.md) - AI 代码生成规范（中文版，266 行）
- [`AI_CODE_SPEC_en.md`](./AI_CODE_SPEC_en.md) - English AI Code Specifications（273 行）

**核心原则（10 条）**：
```markdown
1. 优先使用 gen 生成代码 - 保持代码结构整洁
2. 使用 TDD 开发模式 - 测试先行
3. 生成的代码禁止修改 - Repository/Model/Query 层严禁手动修改
4. 业务逻辑在 Service层 - 所有业务逻辑必须在 Service 层实现
5. 职责边界清晰 - Repository 管数据、Service 管业务、Api 管接口
6. 安全迭代 - 修改数据库 → 重新生成 → Service 层适配
7. 使用结构化日志 - 使用 log/slog
8. 不重复造轮子 - 标准库优先
9. swagger 文档 - 接口采用 swagger 做文档，使用 go generate 生成
10. 遵守 golang 代码规范 - 版本 go 1.26+
```

**基础规范**：
- ✅ 项目结构图（包含生成的模块说明）
- ✅ 分层职责表（Repository/Service/Api/Router）
- ✅ 命名规范（Model/方法/变量/常量/包名）
- ✅ 错误处理示例（正反对比）
- ✅ 日志规范（slog 的 Debug/Info/Warn/Error 级别）

### 2. 多库架构支持 🔗

**代码生成器增强**：
```bash
# 示例 1：生成系统库的 sys_user 表（默认数据库）
go run main.go gen -c resources/config.sqlite.yaml -t sys_user

# 示例 2：生成 notice 库的 message 表（指定数据库）
go run main.go gen -c resources/config.sqlite.yaml -d notice -t message
```

**参数说明**：
- `-c`：配置文件路径
- `-d`：数据库名称（可选，默认为 `sys`）
- `-t`：表名
- `-f`：覆盖已存在的文件（可选）
- `-p`：包名（可选，默认与 `-d` 一致）

**配置文件关系**：
```yaml
# config.sqlite.yaml 中的 dbcfg 配置
dbcfg:
  driver: sqlite
  dns: "temp/dilu.db"
  # ... 其他配置
```

### 3. TDD 编程实践 ✅

**测试先行流程**：
```
1. 编写失败的测试 → 2. 编写最小实现 → 3. 重构优化
```

**测试覆盖要求**：
- **Service层**: 核心业务逻辑覆盖率 ≥ 80%
- **Api 层**: 主要接口测试覆盖率 ≥ 70%

**测试模板**：
```go
func TestUserService_Create(t *testing.T) {
    // Arrange - 准备数据
    svc := service.NewSysUserService()
    user := &model.SysUser{Username: "test"}

    // Act - 执行操作
    result, err := svc.Create(user)

    // Assert - 验证结果
    assert.NoError(t, err)
    assert.NotNil(t, result)
}
```

### 4. 代码简洁优雅原则 ⭐

**应该做的（5 项）**：
- ✅ 简单直接：代码逻辑清晰，一目了然
- ✅ 职责单一：每个函数只做一件事
- ✅ 命名准确：见名知意，无需注释
- ✅ 错误处理：立即返回，不嵌套
- ✅ 代码复用：提取公共逻辑，避免重复

**不应该做的（5 项）**：
- ❌ 过度设计：不要为了炫技而使用复杂模式
- ❌ 炫技编程：避免晦涩的语法糖和技巧
- ❌ 过早优化：先让代码工作，再考虑性能
- ❌ 滥用抽象：不要为了抽象而抽象
- ❌ 魔法数字：使用有意义的常量

**代码对比示例**：

❌ **炫技的代码**（不好）：
```go
result := users.Filter(func(u User) bool {
    return u.Age > 18
}).Map(func(u User) string {
    return u.Name
}).Reduce(...)
```

✅ **简洁的代码**（好）：
```go
var names []string
for _, user := range users {
    if user.Age > 18 {
        names = append(names, user.Name)
    }
}
result := strings.Join(names, ",")
```

### 5. API 开发规范完善 🌐

**RESTful 路由设计**：
```go
// QueryPage - POST /api/v1/sys/user/page     - 分页查询列表
// Get      - GET  /api/v1/sys/user/{id}      - 获取单个记录
// Create   - POST /api/v1/sys/user/create    - 创建记录
// Update   - POST /api/v1/sys/user/update    - 更新记录
// Del      - POST /api/v1/sys/user/del       - 删除记录
```

**统一响应格式**：
```go
e.Ok(c, data)           // 成功（带数据）
e.Ok(c)                 // 成功（无数据）
e.Page(c, list, ...)    // 分页响应
e.Error(c, err)         // 错误响应
```

### 6. 最佳实践清单 📝

**应该做的（8 项）**：
- [x] 使用 gen 命令生成代码
- [x] 在 Service 层实现业务逻辑
- [x] 为 Service 层编写单元测试
- [x] 使用相对路径引用文件
- [x] 遵循 RESTful API 规范
- [x] 及时更新 Swagger 文档（使用 go generate）
- [x] 使用结构化日志（log/slog）
- [x] 妥善处理所有错误

**禁止做的（8 项）**：
- [ ] 修改 `.gen.go` 文件
- [ ] 在 Api 层实现业务逻辑
- [ ] 忽略错误返回值
- [ ] 使用绝对路径
- [ ] 写死配置信息
- [ ] 缺少 Swagger 注解
- [ ] 不写单元测试
- [ ] 直接打印日志

---

## 📊 统计信息

| 类别 | 旧版本 | 新版本 | 变化 |
|------|--------|--------|------|
| **文档总行数** | ~2,500 | ~2,800 | +300 |
| **AI 规范文档** | 0 | 539 | +539 |
| **新增文件** | - | 2 | AI_CODE_SPEC* |
| **代码块数量** | ~20 | ~30 | +10 |
| **最佳实践** | 16 项 | 16 项 | 保持不变 |

**文档体积对比**：
- `AI_CODE_SPEC.md`: 266 行，7.4KB
- `AI_CODE_SPEC_en.md`: 273 行，8.1KB
- 总计：539 行，15.5KB

---

## 🐛 问题修复

- ✅ 日志规范格式问题（Python 风格注释 → Go 风格注释）
- ✅ 代码块语法修正（添加正确的语言标识符）
- ✅ 多库架构说明优化（标题更清晰）
- ✅ 参数说明完善（-c/-d/-t/-f/-p 完整说明）
- ✅ API 接口示例补充（完整实现代码）
- ✅ 最佳实践清单格式统一

---

## 📞 支持与反馈

遇到问题？请查看：
1. [README.md](./README.md) - 项目主文档
2. [AI_CODE_SPEC.md](./AI_CODE_SPEC.md) - AI 代码规范
3. [Issue Tracker](https://github.com/baowk/dilu/issues) - 提交问题

---

<p align="center">
  Made with ❤️ by <a href="https://github.com/baowk">baowk</a>
</p>