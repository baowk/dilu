package models

import "github.com/baowk/dilu-core/core/base"

type SysMenu struct {
	MenuId     int       `json:"menuId" gorm:"type:int unsigned;primaryKey;autoIncrement; comment:编码"` // 编码
	MenuName   string    `json:"menuName" gorm:"size:128;comment:菜单name"`                              //菜单name
	Title      string    `json:"title" gorm:"size:128;comment:显示名称"`                                   //显示名称
	Icon       string    `json:"icon" gorm:"size:128;comment:图标"`                                      //图标
	Path       string    `json:"path" gorm:"size:128;comment:路径"`                                      //路径
	Paths      string    `json:"paths" gorm:"size:128;comment:路径ids/分割"`                               //id路径
	MenuType   string    `json:"menuType" gorm:"size:1;comment:菜单类型"`                                  //菜单类型
	Action     string    `json:"action" gorm:"size:16;comment:请求方式"`                                   //请求方式"`
	Permission string    `json:"permission" gorm:"size:255;comment:权限"`                                //权限
	ParentId   int       `json:"parentId" gorm:"type:int unsigned;comment:菜单父id"`                      //菜单父id
	NoCache    bool      `json:"noCache" gorm:"size:8;comment:是否缓存"`                                   //是否缓存
	Breadcrumb string    `json:"breadcrumb" gorm:"size:255;comment:是否面包屑"`                             //是否面包屑
	Component  string    `json:"component" gorm:"size:255;comment:前端组件"`                               //前端组件
	Sort       int       `json:"sort" gorm:"size:4;comment:排序倒叙"`                                      //排序
	Visible    string    `json:"visible" gorm:"size:1;comment:是否显示"`                                   //是否显示
	IsFrame    string    `json:"isFrame" gorm:"size:1;DEFAULT:0;comment:是否frame"`                      //是否frame
	SysApi     []SysApi  `json:"sysApi" gorm:"many2many:sys_menu_api_rule"`                            //api关联表
	Children   []SysMenu `json:"children,omitempty" gorm:"-"`                                          //子菜单
	IsSelect   bool      `json:"is_select" gorm:"-"`                                                   //是否选中
	base.ControlBy
	base.ModelTime
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
