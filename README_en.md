# Dilu - Go Web Rapid Development Framework

<p align="center">
  <img src="https://github.com/baowk/dilu/assets/142554979/ee341fb7-f98e-4f18-9658-f89b4f7d466f" alt="Dilu Logo" width="200">
</p>

<h3 align="center">Modern Web Development Framework Based on Gin + GORM</h3>

<p align="center">
  <a href="./README.md">🇨🇳 中文版本</a> • 🇺🇸 English
</p>

<p align="center">
  <a href="#project-overview">📖 Overview</a> •
  <a href="#core-features">✨ Features</a> •
  <a href="#quick-start">🚀 Quick Start</a> •
  <a href="#directory-structure">📁 Directory</a> •
  <a href="#ai-development-guide">🤖 AI Guide</a>
</p>

<p align="center">
  <a href="https://github.com/baowk/dilu/stargazers">
    <img src="https://img.shields.io/github/stars/baowk/dilu" alt="GitHub stars">
  </a>
  <a href="https://goreportcard.com/report/github.com/baowk/dilu">
    <img src="https://goreportcard.com/badge/github.com/baowk/dilu" alt="Go Report Card">
  </a>
</p>

## 📖 Project Overview

Dilu is a modern web rapid development framework built on top of Gin + GORM, designed specifically to enhance Go language web application development efficiency. Through long-term practice and iterative optimization, it has evolved into a stable and reliable development solution.

### 🔧 Version Information

- **Dilu Core**: Core simplified version focusing on basic web development functions
- **Dilu Ctl**: Scaffolding tool for rapid project creation

> 💡 **Recommendation**: Beginners are recommended to use the [Dilu Ctl](https://github.com/baowk/dilu-ctl) scaffolding tool to quickly create projects. In the AI era, [Dilu All](https://github.com/baowk/dilu-all) is not recommended.

## ✨ Core Features

### 🚀 Technical Architecture
- **Core Framework**: Gin + GORM combination with excellent performance and mature ecosystem
- **Code Generation**: Built-in command-line code generator supporting rapid CRUD development
- **SaaS Design**: Multi-tenant architecture support with flexible permission management system
- **Modern Approach**: Frontend-backend separation architecture based on Vue3 + Element Plus

### 🛠️ Key Features
- ⚡ High-performance HTTP service engine
- 🗃️ Comprehensive database ORM operations
- 🔐 JWT authentication and permission control
- 📊 Automatic Swagger API documentation generation
- 🐳 Docker containerized deployment support
- 📱 RESTful API design standards
- 🤖 **AI-Assisted Development** - Full support for Lingma, GitHub Copilot and other AI tools

## 🚀 Quick Start

### Install Dependencies
```bash
go mod tidy
```

### Configure Database
Edit `resources/config.dev.yaml` configuration file:
```yaml
dbcfg:
  driver: mysql
  dns: root:password@tcp(127.0.0.1:3306)/dilu-db?charset=utf8mb4&parseTime=True
```

### Start Service
```bash
go run main.go start -c resources/config.dev.yaml
```

Access Swagger documentation at: http://localhost:7888/swagger/index.html

## 📁 Directory Structure

```
dilu/
├── cmd/                    # Command line tool entry (start, version, etc.)
├── internal/               # Core business logic layer
│   ├── bootstrap/         # Application initialization
│   ├── common/            # Common utilities (tools, middleware, constants)
│   ├── tools/             # Code generator core implementation
│   │   ├── apis/          # API interface layer
│   │   ├── service/       # Business logic layer
│   │   ├── router/        # Router configuration layer
│   │   └── repository/    # Data access layer
│   │       ├── model/     # Model layer
│   │       └── query/     # Query layer
│   └── modules/           # Business modules
│       └── sys/           # System module (auto-generated + manual extension)
│           ├── repository/    # Data access layer (Model, Query)
│           │   ├── model/     # Model layer (*.gen.go)
│           │   └── query/     # Query layer (*.gen.go)
│           ├── apis/          # API interface layer
│           ├── service/       # Business logic layer
│           └── router/        # Router configuration layer
├── resources/              # Configuration files (different environments)
├── scripts/                # Scripts (database initialization, migration, etc.)
├── temp/                   # Temporary files (logs, SQLite database, etc.)
├── docs/                   # User documentation (API docs, Swagger auto-generated)
├── dev-docs/               # Development docs (design docs, AI development memory docs)
├── tests/                  # Test code (unit tests, integration tests)
├── main.go                 # Program entry point
└── go.mod                  # Dependency management
```

### 📂 Important Directories

- **`docs/`** - **User Documentation Directory**
  - API Documentation (Swagger auto-generated)
  - User manuals
  - Deployment guides
  - Public technical documentation

- **`dev-docs/`** - **Development Documentation Directory**
  - System design documents
  - Architecture design documents
  - Necessary memory documents during AI development process
  - Technical specification documents
  - Internal development references

- **`tests/`** - **Test Code Directory**
  - Unit tests (`*_test.go`)
  - Integration tests
  - Performance tests
  - Test data and Mock objects

## 🤖 AI Development Guide

Dilu framework provides comprehensive AI-assisted development support to help you efficiently build enterprise-grade Go applications.

### 📋 AI Code Specifications

- 🇨🇳 **[Chinese AI Code Spec](./AI_CODE_SPEC.md)** - Development specifications and conventions for AI code generation
- 🌐 **[English AI Code Spec](./dev-docs/ai/AI_CODE_SPEC_en.md)** - Human-readable note (English)

### Read AI Development Guides

- 🇨🇳 **[Chinese AI Development Guide](./README_AI_DEV.md)** - Detailed Chinese tutorial
- 🌐 **[English AI Dev Guide](./README_AI_DEV_en.md)** - Complete English guide

### Core AI Development Capabilities

- ✅ **Intelligent Code Generation** - Use `dilu-ctl gen` to rapidly generate Model, Service, Api, Router code
- ✅ **API Development Assistant** - AI-assisted RESTful API design with automatic Swagger documentation
- ✅ **Database Design** - AI-assisted table structure design, GORM model and migration script generation
- ✅ **Frontend Development** - AI generates Vue3 components and TypeScript API wrappers
- ✅ **Test Writing** - AI automatically generates unit tests and integration tests
- ✅ **Documentation** - AI-assisted API documentation and technical writing

### Quick Example

```
# 1. Prepare database (using SQLite)
go run scripts/init_sqlite.go

# 2. Call dilu-ctl to generate code
dilu-ctl gen -t sys_user -d "sqlite:./data/dev.db" -P . -f false

# 3. View generated files
ls internal/modules/sys/repository/model/sys_user.gen.go
ls internal/modules/sys/repository/query/sys_user.gen.go

# 4. Extend business logic in Service layer
vim internal/modules/sys/service/sys_user_extend.go

# 5. Start service
go run main.go start -c resources/config.sqlite.yaml
```

For detailed usage methods and best practices, please refer to the **AI Development Guide** documents above.

---

<p align="center">
  Made with ❤️ by <a href="https://github.com/baowk">baowk</a>
</p>
