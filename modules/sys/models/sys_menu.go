package models

import "github.com/baowk/dilu-core/core/base"

type SysMenu struct {
	Id           int      `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement; comment:主键"` // 主键
	MenuName     string   `json:"menuName" gorm:"size:128;comment:菜单名"`                             //菜单名
	Title        string   `json:"title" gorm:"size:128;comment:显示名称"`                               //显示名称
	Icon         string   `json:"icon" gorm:"size:128;comment:图标"`                                  //图标
	Path         string   `json:"path" gorm:"size:128;comment:路径"`                                  //路径
	PlatformType int      `json:"platformType" gorm:"tinyint unsigned;comment:平台类型 1 平台管理 2团队管理"`   //平台类型 1 平台管理 2团队管理
	MenuType     int      `json:"menuType" gorm:"size:1;comment:菜单类型 1 分类 2菜单 3方法按钮"`               //菜单类型
	Permission   string   `json:"permission" gorm:"size:255;comment:权限"`                            //权限
	ParentId     int      `json:"parentId" gorm:"type:int unsigned;comment:菜单父id"`                  //菜单父id
	NoCache      bool     `json:"noCache" gorm:"size:4;comment:是否缓存"`                               //是否缓存
	Component    string   `json:"component" gorm:"size:255;comment:前端组件路径"`                         //前端组件路径
	Sort         int      `json:"sort" gorm:"size:4;comment:排序倒叙"`                                  //排序
	Hidden       bool     `json:"hidden" gorm:"size:1;comment:是否隐藏"`                                //隐藏
	SysApi       []SysApi `json:"sysApi" gorm:"many2many:sys_menu_api_rule"`                        //api关联表
	//Children     []SysMenu `json:"children" gorm:"-"`
	base.ControlBy
	base.ModelTime
}

func (SysMenu) TableName() string {
	return "sys_menu"
}
