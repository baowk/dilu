# Dilu - Go Web Rapid Development Framework

<p align="center">
  <img src="https://github.com/baowk/dilu/assets/142554979/ee341fb7-f98e-4f18-9658-f89b4f7d466f" alt="Dilu Logo" width="200">
</p>

<h3 align="center">Modern Web Development Framework Based on Gin + GORM</h3>

<p align="center">
  <a href="./README.md">🇨🇳 中文版本</a> • 🇺🇸 English
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

## 📖 Project Overview

Dilu is a modern web rapid development framework built on top of Gin + GORM, designed specifically to enhance Go language web application development efficiency. Through long-term practice and iterative optimization, it has evolved into a stable and reliable development solution.

### 🔧 Version Information

- **Dilu Core**: Core simplified version focusing on basic web development functions
- **Dilu All**: Full version with all plugins and advanced features
- **Dilu Plugin**: Plugin library for extended functionality modules
- **Dilu Ctl**: Scaffolding tool for rapid project creation

> 💡 **Recommendation**: Beginners are recommended to use the [Dilu Ctl](https://github.com/baowk/dilu-ctl) scaffolding tool to quickly create projects, or start learning from [Dilu All](https://github.com/baowk/dilu-all).

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

## 🚀 Quick Start

### Method 1: Using Scaffolding Tool (Recommended)

Use the [Dilu Ctl](https://github.com/baowk/dilu-ctl) scaffolding tool to create projects with one command:

```bash
# Install scaffolding tool
go install github.com/baowk/dilu-ctl@latest

# Create basic Dilu project
dilu-ctl -n myproject

# Create complete project (including admin frontend)
dilu-ctl -n myproject -a

# Enter project directory and start
cd myproject
go run main.go start -c resources/config.dev.yaml
```

### Method 2: Manual Installation

1. **Get Source Code**
```bash
git clone https://github.com/baowk/dilu.git
# or
git clone https://gitee.com/walkbao/dilu.git
```

2. **Install Dependencies**
```bash
cd dilu
go mod tidy
```

3. **Database Configuration**
```sql
-- Create database
CREATE DATABASE `dilu-db` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

Modify configuration file `resources/config.dev.yaml`:
```yaml
dbcfg:
  driver: mysql
  dns: root:password@tcp(127.0.0.1:3306)/dilu-db?charset=utf8&parseTime=True&loc=Local
```

4. **Code Generation Example**
```bash
# Generate code for sys_user table
go run main.go gen -c resources/config.dev.yaml -d sys -t sys_user -f false
```

5. **Start Service**
```bash
go run main.go start -c resources/config.dev.yaml
```

Service runs by default on `http://localhost:8000`

## 🏗️ Project Structure

```
dilu/
├── cmd/           # Command line tools
├── common/        # Common components
├── docs/          # Documentation
├── modules/       # Business modules
├── resources/     # Configuration files and resources
├── temp/          # Temporary files
├── main.go        # Application entry point
└── go.mod         # Dependency management
```

## 📚 Usage Guide

### Command Line Tools
```bash
# View help
go run main.go --help

# Start service
go run main.go start -c config.yaml

# Generate code
go run main.go gen -c config.yaml -d database -t table -f true

# Build project
go build -o dilu main.go
```

### API Documentation
After starting the service, visit: `http://localhost:8000/swagger/index.html`

## 🛠️ Development Tools

### Dilu Ctl Scaffolding Tool
[Dilu Ctl](https://github.com/baowk/dilu-ctl) is a scaffolding tool specially designed for the Dilu framework, with main features including:

- ✅ Rapid project creation via command line
- ✅ Support for different template repositories (Core/All versions)
- ✅ Intelligent package name replacement and configuration updates
- ✅ Automatic generation of correct go.mod files
- ✅ Support for both SSH and HTTPS Git protocols
- ✅ Option to create accompanying admin frontend projects

**Installation and Usage:**
```bash
# Installation
go install github.com/baowk/dilu-ctl@latest

# Create project
dilu-ctl -n project-name [-a] [--https] [-u username]

# View help
dilu-ctl -h
```

## 🤝 Community

Welcome to join our technical discussion group!

<div align="center">
  <img src="https://github.com/baowk/dilu/assets/142554979/29a6863c-4bdc-4963-99c2-0c400e132f6f" width="300" alt="WeChat Group QR Code">
  <br>
  <sub>Please note "Dilu" when adding WeChat</sub>
</div>

### Contribution Guidelines
We welcome all forms of contributions:
- 🐛 Submit bug reports
- 💡 Propose feature suggestions
- 🔧 Submit code improvements
- 📝 Improve documentation

## 🔗 Related Projects

- [Dilu Admin](https://github.com/baowk/dilu-admin) - Frontend admin interface
- [Dilu All](https://github.com/baowk/dilu-all) - Full-featured version
- [Dilu Plugin](https://github.com/baowk/dilu-plugin) - Plugin extension library
- [Dilu Ctl](https://github.com/baowk/dilu-ctl) - Project scaffolding tool ⭐
- [Online Demo](http://dilu.youwan.art) - System demonstration environment

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

<p align="center">
  Made with ❤️ by <a href="https://github.com/baowk">baowk</a>
</p>