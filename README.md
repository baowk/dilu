<div align="center">
  <img src="https://github.com/baowk/dilu/assets/142554979/ee341fb7-f98e-4f18-9658-f89b4f7d466f" alt="Dilu Logo" width="200">
</div>

<h1 align="center">Dilu - Go Web快速开发框架</h1>

<p align="center">
  基于 Gin + GORM 的现代化Web开发框架
</p>

<p align="center">
  🇨🇳 中文版本 • <a href="./README_en.md">🇺🇸 English</a>
</p>

<p align="center">
  <a href="https://github.com/baowk/dilu/stargazers">
    <img src="https://img.shields.io/github/stars/baowk/dilu" alt="GitHub stars">
  </a>
  <a href="https://github.com/baowk/dilu/issues">
    <img src="https://img.shields.io/github/issues/baowk/dilu" alt="GitHub issues">
  </a>
  <a href="https://github.com/baowk/dilu/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/baowk/dilu" alt="GitHub">
  </a>
  <a href="https://goreportcard.com/report/github.com/baowk/dilu">
    <img src="https://goreportcard.com/badge/github.com/baowk/dilu" alt="Go Report Card">
  </a>
</p>

## 📖 项目简介

Dilu是一套基于Gin + GORM封装的现代化Web快速开发框架，专为提升Go语言Web应用开发效率而设计。经过长期实践和优化迭代，现已发展为稳定可靠的开发解决方案。

### 🔧 版本说明

- **Dilu Core**: 核心简化版本，专注于基础Web开发功能
- **Dilu All**: 完整版，包含所有插件和高级功能
- **Dilu Plugin**: 插件库，扩展功能模块
- **Dilu Ctl**: 脚手架工具，快速搭建项目

> 💡 **建议**: 新手用户推荐使用 [Dilu Ctl](https://github.com/baowk/dilu-ctl) 脚手架工具快速创建项目，或从 [Dilu All](https://github.com/baowk/dilu-all) 开始学习。

## ✨ 核心特性

### 🚀 技术架构
- **核心框架**: Gin + GORM 组合，性能优异且生态完善
- **代码生成**: 内置命令行代码生成器，支持快速CRUD开发
- **SaaS设计**: 支持多租户架构，灵活的权限管理体系
- **现代化**: 基于Vue3 + Element Plus的前后端分离架构

### 🛠️ 功能亮点
- ⚡ 高性能HTTP服务引擎
- 🗃️ 完善的数据库ORM操作
- 🔐 JWT认证和权限控制
- 📊 Swagger API文档自动生成
- 🐳 Docker容器化部署支持
- 📱 RESTful API设计规范

## 🚀 快速开始

### 方法一：使用脚手架工具（推荐）

使用 [Dilu Ctl](https://github.com/baowk/dilu-ctl) 脚手架工具可以一键创建项目：

```bash
# 安装脚手架工具
go install github.com/baowk/dilu-ctl@latest

# 创建基础Dilu项目
dilu-ctl -n myproject

# 创建完整项目（包含admin前端）
dilu-ctl -n myproject -a

# 进入项目目录并启动
cd myproject
go run main.go start -c resources/config.dev.yaml
```

### 方法二：手动安装

1. **获取源码**
```bash
git clone https://github.com/baowk/dilu.git
# 或者
git clone https://gitee.com/walkbao/dilu.git
```

2. **安装依赖**
```bash
cd dilu
go mod tidy
```

3. **数据库配置**
```sql
-- 创建数据库
CREATE DATABASE `dilu-db` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

修改配置文件 `resources/config.dev.yaml`:
```yaml
dbcfg:
  driver: mysql
  dns: root:password@tcp(127.0.0.1:3306)/dilu-db?charset=utf8&parseTime=True&loc=Local
```

4. **代码生成示例**
```bash
# 生成sys_user表的代码
go run main.go gen -c resources/config.dev.yaml -d sys -t sys_user -f false
```

5. **启动服务**
```bash
go run main.go start -c resources/config.dev.yaml
```

服务默认运行在 `http://localhost:8000`

## 🏗️ 项目结构

```
dilu/
├── cmd/           # 命令行工具
├── common/        # 公共组件
├── docs/          # 文档资料
├── modules/       # 业务模块
├── resources/     # 配置文件和资源
├── temp/          # 临时文件
├── main.go        # 程序入口
└── go.mod         # 依赖管理
```

## 📚 使用指南

### 命令行工具
```bash
# 查看帮助
go run main.go --help

# 启动服务
go run main.go start -c config.yaml

# 生成代码
go run main.go gen -c config.yaml -d database -t table -f true

# 构建项目
go build -o dilu main.go
```

### API文档
启动服务后访问: `http://localhost:8000/swagger/index.html`

## 🛠️ 开发工具

### Dilu Ctl 脚手架工具
[Dilu Ctl](https://github.com/baowk/dilu-ctl) 是专门为Dilu框架设计的脚手架工具，主要功能包括：

- ✅ 通过命令行快速创建 Dilu 项目
- ✅ 支持选择不同的模板仓库（Core/All版本）
- ✅ 智能包名替换和配置文件更新
- ✅ 自动生成正确的 go.mod 文件
- ✅ 支持SSH和HTTPS两种Git协议
- ✅ 可选择是否创建配套的 admin 前端项目

**安装和使用：**
```bash
# 安装
go install github.com/baowk/dilu-ctl@latest

# 创建项目
dilu-ctl -n 项目名称 [-a] [--https] [-u username]

# 查看帮助
dilu-ctl -h
```

## 🤝 社区交流

欢迎加入我们的技术交流群！

<div align="center">
  <img src="https://github.com/baowk/dilu/assets/142554979/29a6863c-4bdc-4963-99c2-0c400e132f6f" width="300" alt="微信群二维码">
  <br>
  <sub>添加微信时请备注：的卢</sub>
</div>

### 贡献指南
我们欢迎任何形式的贡献：
- 🐛 提交Bug报告
- 💡 提出功能建议
- 🔧 提交代码改进
- 📝 完善文档内容

## 🔗 相关项目

- [Dilu Admin](https://github.com/baowk/dilu-admin) - 前端管理界面
- [Dilu All](https://github.com/baowk/dilu-all) - 完整功能版本
- [Dilu Plugin](https://github.com/baowk/dilu-plugin) - 插件扩展库
- [Dilu Ctl](https://github.com/baowk/dilu-ctl) - 项目脚手架工具 ⭐
- [在线演示](http://dilu.youwan.art) - 系统演示环境

## 📄 许可证

本项目采用 MIT 许可证，详情请参见 [LICENSE](LICENSE) 文件。

---

<p align="center">
  Made with ❤️ by <a href="https://github.com/baowk">baowk</a>
</p>