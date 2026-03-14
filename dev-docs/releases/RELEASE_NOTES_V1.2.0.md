# 发布说明 v1.2.0 - AI 代码规范与目录结构优化

**发布日期**: 2026-03-14  
**版本类型**: 重大更新  
**兼容性**: ⚠️ 破坏性变更（目录结构调整）

---

## 🎯 核心更新

本次更新聚焦于建立 AI 可读文档体系和标准化目录结构：

1. **AI 代码规范体系** - 新增 `AI_CODE_SPEC.md` 系列文档
2. **目录结构重构** - 统一 tools 和 sys 模块的分层架构
3. **文档国际化** - 全面优化中英文版本文档一致性
4. **路径引用规范化** - 所有绝对路径改为相对路径

---

## ✨ 主要特性

### 1. AI 代码规范文档 🤖

**新增文件**：
- [`AI_CODE_SPEC.md`](./AI_CODE_SPEC.md) - AI 代码生成规范（中文版）
- [`AI_CODE_SPEC_en.md`](./AI_CODE_SPEC_en.md) - English AI Code Specifications

**核心原则**：
```markdown
> ⚠️ 重要说明：本文档是给 AI 助手阅读和遵循的代码生成规范，
> 不是给人阅读的教程。AI 在生成代码时必须严格遵守以下规范和约定。
```

**关键规范**：
- ✅ 生成的代码禁止修改
- ✅ 业务逻辑集中于Service 层
- ✅ 清晰的职责边界（Repository → Service → Api → Router）
- ✅ 安全迭代流程：修改数据库 → 重新生成代码 → Service 层适配

### 2. 目录结构标准化 📂

**tools 目录重构**（与 sys 保持一致）：
```
internal/tools/
├── apis/              # API 接口层
├── service/           # 业务逻辑层
├── router/            # 路由配置层
└── repository/        # 数据访问层 ⭐ 新增
    ├── model/         # Model 模型层 ⭐ 新增
    └── query/         # Query 查询层 ⭐ 新增
```

**优势**：
- ✅ 统一的分层架构
- ✅ 清晰的职责划分
- ✅ 易于理解和维护

### 3. 发布文档管理 📋

**新增规范**：
- 所有发布文档存放在 `dev-docs/releases/` 目录
- 命名格式：`RELEASE_NOTES_V{version}.md`
- 示例：[`RELEASE_NOTES_V1.2.0.md`](./dev-docs/releases/RELEASE_NOTES_V1.2.0.md)

---

## 🔧 破坏性变更

### ⚠️ 重要提醒

本次更新包含**破坏性变更**，需要手动迁移现有项目。

#### 目录结构变更

**影响范围**：
- ❌ 删除：`internal/tools/models/` 目录
- ✅ 新增：`internal/tools/repository/model/` 目录
- ✅ 新增：`internal/tools/repository/query/` 目录

**受影响的文件**：
```bash
# 移动的文件
internal/tools/models/*.go → internal/tools/repository/model/*.go

# 修改的 package 声明
package models → package model
```

---

## ⚠️ 快速迁移指南

### 步骤 1：备份代码
```bash
cp -r dilu dilu_backup_$(date +%Y%m%d)
```

### 步骤 2：更新目录结构
```bash
cd dilu
mkdir -p internal/tools/repository/model
mv internal/tools/models/* internal/tools/repository/model/
rm -rf internal/tools/models
```

### 步骤 3：更新 package 声明
```bash
find internal/tools/repository/model -name "*.go" -exec \
  sed -i '' 's/^package models$/package model/' {} \;
```

### 步骤 4：验证编译
```bash
go clean && go build -o dilu main.go
```

---

## 📊 统计信息

| 类别 | 旧版本 | 新版本 | 变化 |
|------|--------|--------|------|
| 文档总行数 | ~2,500 | ~3,300 | +800 |
| AI 规范文档 | 0 | 1,163 | +1,163 |
| 新增文件 | - | 2 | AI_CODE_SPEC* |
| 删除文件 | - | 2 | ~~README_AI_DEV*~~ |

---

## 🐛 问题修复

- ✅ 文档路径引用问题（绝对路径 → 相对路径）
- ✅ tools 目录结构混乱问题
- ✅ AI 文档定位不清问题

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