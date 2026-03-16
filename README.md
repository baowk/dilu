<div align="center">
  <img src="https://github.com/baowk/dilu/assets/142554979/ee341fb7-f98e-4f18-9658-f89b4f7d466f" alt="Dilu Logo" width="200">
</div>

<h1 align="center">Dilu - Go Web快速开发框架</h1>

<p align="center">
  基于 Gin + GORM 的现代化 Web 开发框架
</p>

<p align="center">
  🇨🇳 中文版本 • <a href="./README_en.md">🇺🇸 English</a>
</p>

<p align="center">
  <a href="#项目简介">📖 项目简介</a> •
  <a href="#核心特性">✨ 核心特性</a> •
  <a href="#快速开始">🚀 快速开始</a> •
  <a href="#目录结构">📁 目录结构</a> •
  <a href="#AI-开发指南">🤖 AI 开发指南</a>
</p>

<p align="center">
  <a href="https://github.com/baowk/dilu/stargazers">
    <img src="https://img.shields.io/github/stars/baowk/dilu" alt="GitHub stars">
  </a>
  <a href="https://goreportcard.com/report/github.com/baowk/dilu">
    <img src="https://goreportcard.com/badge/github.com/baowk/dilu" alt="Go Report Card">
  </a>
</p>

## 📖 项目简介

Dilu是一套基于Gin + GORM封装的现代化Web快速开发框架，专为提升Go语言Web应用开发效率而设计。经过长期实践和优化迭代，现已发展为稳定可靠的开发解决方案。

### 🔧 版本说明

- **Dilu Core**: 核心简化版本，专注于基础 Web 开发功能
- **Dilu Ctl**: 脚手架工具，快速搭建项目

> 💡 **建议**: 新手用户推荐使用 [Dilu Ctl](https://github.com/baowk/dilu-ctl) 脚手架工具快速创建项目，Ai时代不建议用 [Dilu All](https://github.com/baowk/dilu-all) 。

## ✨ 核心特性

### 🚀 技术架构
- **核心框架**: Gin + GORM 组合，性能优异且生态完善
- **代码生成**: 内置命令行代码生成器，支持快速 CRUD 开发
- **SaaS 设计**: 支持多租户架构，灵活的权限管理体系
- **现代化**: 基于 Vue3 + Element Plus 的前后端分离架构

### 🛠️ 功能亮点
- ⚡ 高性能 HTTP 服务引擎
- 🗃️ 完善的数据库 ORM 操作
- 🔐 JWT 认证和权限控制
- 📊 Swagger API 文档自动生成
- 🐳 Docker 容器化部署支持
- 📱 RESTful API 设计规范
- 🤖 **AI 辅助开发** - 全面支持通义灵码、GitHub Copilot 等 AI 工具

## 🚀 快速开始

### 安装依赖
```bash
go mod tidy
```

### 配置数据库
编辑 `resources/config.dev.yaml` 配置文件：
```yaml
dbcfg:
  driver: mysql
  dns: root:password@tcp(127.0.0.1:3306)/dilu-db?charset=utf8mb4&parseTime=True
```

### 启动服务
```bash
go run main.go start -c resources/config.dev.yaml
```

访问 Swagger 文档：http://localhost:7888/swagger/index.html

## 📁 目录结构

```
dilu/
├── cmd/                    # 命令行工具入口（start、version 等）
├── internal/               # 核心业务逻辑层
│   ├── bootstrap/         # 应用初始化引导
│   ├── common/            # 公共组件库（工具函数、中间件、常量定义）
│   ├── tools/             # 代码生成器核心实现
│   │   ├── apis/          # API 接口层
│   │   ├── service/       # 业务逻辑层
│   │   ├── router/        # 路由配置层
│   │   └── repository/    # 数据访问层
│   │       ├── model/     # Model 模型层
│   │       └── query/     # Query 查询层
│   └── modules/           # 业务模块
│       └── sys/           # 系统管理模块（自动生成 + 手动扩展）
│           ├── repository/    # 数据访问层（Model、Query）
│           │   ├── model/     # Model 模型层 (*.gen.go)
│           │   └── query/     # Query 查询层 (*.gen.go)
│           ├── apis/          # API 接口层
│           ├── service/       # 业务逻辑层
│           └── router/        # 路由配置层
├── resources/              # 配置文件区（不同环境配置）
├── scripts/                # 脚本工具区（数据库初始化、迁移等）
├── temp/                   # 临时文件区（日志、SQLite 数据库等）
├── docs/                   # 用户文档（API 文档、Swagger 自动生成）
├── dev-docs/               # 开发文档（设计文档、AI 开发过程记忆文档）
├── tests/                  # 测试代码（单元测试、集成测试）
├── main.go                 # 程序入口
└── go.mod                  # 依赖管理
```

### 📂 重要目录说明

- **`docs/`** - **用户文档目录**
  - API 文档（Swagger 自动生成）
  - 用户使用手册
  - 部署指南
  - 对外公开的技术文档

- **`dev-docs/`** - **开发文档目录**
  - 系统设计文档
  - 架构设计文档
  - AI 开发过程中的必要记忆文档
  - 技术规范文档
  - 内部开发参考资料

- **`tests/`** - **测试代码目录**
  - 单元测试（`*_test.go`）
  - 集成测试
  - 性能测试
  - 测试数据和 Mock 对象

## 🤖 AI 开发指南

Dilu框架提供完整的 AI 辅助开发支持，帮助您高效构建企业级 Go 应用。

### 📋 AI 代码规范

- 🇨🇳 **[中文版 AI 代码规范](./AI_CODE_SPEC.md)** - AI 生成代码的开发规范和约定
- 🌐 **[English AI Code Spec](./AI_CODE_SPEC_en.md)** - AI code generation specifications and conventions


### AI 开发核心能力

- ✅ **智能代码生成** - 使用 `dilu-ctl gen` 快速生成 Model、Service、Api、Router 代码
- ✅ **API 开发辅助** - AI 辅助设计 RESTful API，自动生成 Swagger 文档
- ✅ **数据库设计** - AI 辅助设计表结构，生成 GORM 模型和迁移脚本
- ✅ **前端开发** - AI 生成 Vue3 组件和 TypeScript API 封装
- ✅ **测试编写** - AI 自动生成单元测试和集成测试
- ✅ **文档生成** - AI 辅助编写 API 文档和技术文档

### 快速示例

```
# 1. 准备数据库（使用 SQLite）
go run scripts/init_sqlite.go

# 2. 调用 dilu-ctl 生成代码
dilu-ctl gen -t sys_user -d "sqlite:./data/dev.db" -P . -f false

# 3. 查看生成的文件
ls internal/modules/sys/repository/model/sys_user.gen.go
ls internal/modules/sys/repository/query/sys_user.gen.go

# 4. 在 Service 层扩展业务逻辑
vim internal/modules/sys/service/sys_user_extend.go

# 5. 启动服务
go run main.go start -c resources/config.sqlite.yaml
```

详细的使用方法和最佳实践，请查阅上方的 **AI 开发指南** 文档。

---

<p align="center">
  Made with ❤️ by <a href="https://github.com/baowk">baowk</a>
</p>
