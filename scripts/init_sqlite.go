package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键 ID"`
	CreatedAt time.Time `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"comment:更新时间"`
	Username  string    `json:"username" gorm:"size:100;not null;uniqueIndex;comment:用户名"`
	Password  string    `json:"-" gorm:"size:255;not null;comment:密码"`
	Nickname  string    `json:"nickname" gorm:"size:100;comment:昵称"`
	Email     string    `json:"email" gorm:"size:100;comment:邮箱"`
	Phone     string    `json:"phone" gorm:"size:20;comment:手机号"`
	Avatar    string    `json:"avatar" gorm:"size:500;comment:头像 URL"`
	Status    int       `json:"status" gorm:"default:1;comment:状态 1 正常 2 禁用"`
	DeptId    int       `json:"dept_id" gorm:"comment:部门 ID"`
	Remark    string    `json:"remark" gorm:"size:500;comment:备注"`
}

func (User) TableName() string {
	return "sys_user"
}

// Dept 部门模型
type Dept struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键 ID"`
	CreatedAt  time.Time `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"comment:更新时间"`
	DeptName   string    `json:"dept_name" gorm:"size:100;not null;comment:部门名称"`
	ParentId   int       `json:"parent_id" gorm:"default:0;comment:父部门 ID"`
	Sort       int       `json:"sort" gorm:"default:0;comment:排序"`
	Leader     string    `json:"leader" gorm:"size:50;comment:负责人"`
	Phone      string    `json:"phone" gorm:"size:20;comment:联系电话"`
	Email      string    `json:"email" gorm:"size:100;comment:邮箱"`
	Status     int       `json:"status" gorm:"default:1;comment:状态 1 正常 2 停用"`
	CreatedBy  string    `json:"created_by" gorm:"size:100;comment:创建者"`
	UpdateTime time.Time `json:"update_time" gorm:"comment:更新时间"`
}

func (Dept) TableName() string {
	return "sys_dept"
}

// Role 角色模型
type Role struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键 ID"`
	CreatedAt time.Time `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"comment:更新时间"`
	RoleName  string    `json:"role_name" gorm:"size:100;not null;comment:角色名称"`
	RoleKey   string    `json:"role_key" gorm:"size:100;not null;uniqueIndex;comment:角色权限字符串"`
	Sort      int       `json:"sort" gorm:"default:0;comment:排序"`
	Status    int       `json:"status" gorm:"default:1;comment:状态 1 正常 2 停用"`
	Admin     bool      `json:"admin" gorm:"default:false;comment:是否超级管理员"`
	DataScope string    `json:"data_scope" gorm:"size:20;comment:数据范围"`
	Remark    string    `json:"remark" gorm:"size:500;comment:备注"`
}

func (Role) TableName() string {
	return "sys_role"
}

func main() {
	// 确保 temp 目录存在
	tempDir := filepath.Join(".", "temp")
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		fmt.Printf("创建目录失败：%v\n", err)
		os.Exit(1)
	}

	// 数据库文件路径
	dbPath := filepath.Join(tempDir, "dilu.db")

	// 删除已存在的数据库文件（可选）
	os.Remove(dbPath)
	fmt.Printf("SQLite 数据库路径：%s\n", dbPath)

	// 连接数据库
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		fmt.Printf("连接数据库失败：%v\n", err)
		os.Exit(1)
	}

	fmt.Println("✓ 数据库连接成功")

	// 自动迁移表结构
	err = db.AutoMigrate(&User{}, &Dept{}, &Role{})
	if err != nil {
		fmt.Printf("自动迁移失败：%v\n", err)
		os.Exit(1)
	}

	fmt.Println("✓ 数据库表结构创建成功")

	// 插入示例数据
	seedData(db)

	fmt.Println("✓ 示例数据插入成功")
	fmt.Println("\n📊 数据库初始化完成！")
	fmt.Println("===========================================")
	fmt.Println("数据库文件：" + dbPath)
	fmt.Println("数据表:")
	fmt.Println("  - sys_user (用户表)")
	fmt.Println("  - sys_dept (部门表)")
	fmt.Println("  - sys_role (角色表)")
	fmt.Println("===========================================")
	fmt.Println("\n💡 使用提示:")
	fmt.Println("1. 启动服务：go run main.go start -c resources/config.sqlite.yaml")
	fmt.Println("2. 生成代码：go run main.go gen -c resources/config.sqlite.yaml -d main -t sys_user -f false")
	fmt.Println("3. 查看 Swagger: http://localhost:7888/swagger/index.html")
}

func seedData(db *gorm.DB) {
	// 创建默认部门
	dept := Dept{
		DeptName: "技术部",
		ParentId: 0,
		Sort:     1,
		Leader:   "张三",
		Phone:    "13800138000",
		Email:    "tech@example.com",
		Status:   1,
		CreatedBy: "admin",
	}
	db.Create(&dept)

	// 创建默认角色
	roles := []Role{
		{
			RoleName:  "超级管理员",
			RoleKey:   "admin",
			Sort:      1,
			Status:    1,
			Admin:     true,
			DataScope: "1",
			Remark:    "拥有所有权限",
		},
		{
			RoleName:  "普通用户",
			RoleKey:   "common",
			Sort:      2,
			Status:    1,
			Admin:     false,
			DataScope: "2",
			Remark:    "普通用户权限",
		},
	}
	for _, role := range roles {
		db.Create(&role)
	}

	// 创建默认用户
	users := []User{
		{
			Username: "admin",
			Password: "$2a$10$7JB9QhCZv8M1KqVxN7uLp.F5R6X8Y9W0E1D2C3B4A5F6G7H8I9J0K", // admin123
			Nickname: "系统管理员",
			Email:    "admin@example.com",
			Phone:    "13800138000",
			Status:   1,
			DeptId:   int(dept.ID),
			Remark:   "系统超级管理员",
		},
		{
			Username: "test",
			Password: "$2a$10$7JB9QhCZv8M1KqVxN7uLp.F5R6X8Y9W0E1D2C3B4A5F6G7H8I9J0K", // test123
			Nickname: "测试用户",
			Email:    "test@example.com",
			Phone:    "13900139000",
			Status:   1,
			DeptId:   int(dept.ID),
			Remark:   "测试账号",
		},
	}
	for _, user := range users {
		db.Create(&user)
	}

	fmt.Printf("  ✓ 创建部门：%s\n", dept.DeptName)
	fmt.Printf("  ✓ 创建角色：%d 个\n", len(roles))
	fmt.Printf("  ✓ 创建用户：%d 个\n", len(users))
}
