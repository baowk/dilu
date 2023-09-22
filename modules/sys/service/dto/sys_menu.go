package dto

import (
	"dilu/modules/sys/models"

	"github.com/baowk/dilu-core/core/base"
)

// SysMenuGetPageReq 列表或者搜索使用结构体
type SysMenuGetPageReq struct {
	base.ReqPage `search:"-"`
	Title        string `form:"title" search:"type:contains;column:title;table:sys_menu" comment:"菜单名称"` // 菜单名称
	Hidden       int    `form:"hidden" search:"type:exact;column:hidden;table:sys_menu" comment:"显示状态"`  // 显示状态
}

func (m *SysMenuGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SysMenuInsertReq struct {
	Id         int             `uri:"id" comment:"编码"`             // 编码
	MenuName   string          `form:"menuName" comment:"菜单name"`  //菜单name
	Title      string          `form:"title" comment:"显示名称"`       //显示名称
	Icon       string          `form:"icon" comment:"图标"`          //图标
	Path       string          `form:"path" comment:"路径"`          //路径
	Paths      string          `form:"paths" comment:"id路径"`       //id路径
	MenuType   string          `form:"menuType" comment:"菜单类型"`    //菜单类型
	Permission string          `form:"permission" comment:"权限编码"`  //权限编码
	ParentId   int             `form:"parentId" comment:"上级菜单"`    //上级菜单
	NoCache    bool            `form:"noCache" comment:"是否缓存"`     //是否缓存
	Breadcrumb string          `form:"breadcrumb" comment:"是否面包屑"` //是否面包屑
	Component  string          `form:"component" comment:"组件"`     //组件
	Sort       int             `form:"sort" comment:"排序"`          //排序
	Hidden     bool            `form:"hidden" comment:"是否隐藏"`      //是否隐藏
	SysApi     []models.SysApi `form:"sysApi"`
	base.ControlBy
}

func (s *SysMenuInsertReq) Generate(model *models.SysMenu) {
	if s.Id != 0 {
		model.MenuId = s.Id
	}
	model.MenuName = s.MenuName
	model.Title = s.Title
	model.Icon = s.Icon
	model.Path = s.Path
	model.Paths = s.Paths
	model.MenuType = s.MenuType
	model.SysApi = s.SysApi
	model.Permission = s.Permission
	model.ParentId = s.ParentId
	model.NoCache = s.NoCache
	model.Component = s.Component
	model.Sort = s.Sort
	model.Hidden = s.Hidden
	if s.CreateBy != 0 {
		model.CreateBy = s.CreateBy
	}
	if s.UpdateBy != 0 {
		model.UpdateBy = s.UpdateBy
	}
}

func (s *SysMenuInsertReq) GetId() interface{} {
	return s.Id
}

type SysMenuUpdateReq struct {
	Id         int             `uri:"id" comment:"编码"`            // 编码
	MenuName   string          `form:"menuName" comment:"菜单name"` //菜单name
	Title      string          `form:"title" comment:"显示名称"`      //显示名称
	Icon       string          `form:"icon" comment:"图标"`         //图标
	Path       string          `form:"path" comment:"路径"`         //路径
	Paths      string          `form:"paths" comment:"id路径"`      //id路径
	MenuType   string          `form:"menuType" comment:"菜单类型"`   //菜单类型
	SysApi     []models.SysApi `form:"sysApi"`
	Apis       []int           `form:"apis"`
	Permission string          `form:"permission" comment:"权限编码"` //权限编码
	ParentId   int             `form:"parentId" comment:"上级菜单"`   //上级菜单
	NoCache    bool            `form:"noCache" comment:"是否缓存"`    //是否缓存
	Component  string          `form:"component" comment:"组件"`    //组件
	Sort       int             `form:"sort" comment:"排序"`         //排序
	Hidden     bool            `form:"hidden" comment:"是否显示"`     //是否显示
	base.ControlBy
}

func (s *SysMenuUpdateReq) Generate(model *models.SysMenu) {
	if s.Id != 0 {
		model.MenuId = s.Id
	}
	model.MenuName = s.MenuName
	model.Title = s.Title
	model.Icon = s.Icon
	model.Path = s.Path
	model.Paths = s.Paths
	model.MenuType = s.MenuType
	model.Permission = s.Permission
	model.ParentId = s.ParentId
	model.NoCache = s.NoCache
	model.Component = s.Component
	model.Sort = s.Sort
	model.Hidden = s.Hidden
	if s.CreateBy != 0 {
		model.CreateBy = s.CreateBy
	}
	if s.UpdateBy != 0 {
		model.UpdateBy = s.UpdateBy
	}
}

func (s *SysMenuUpdateReq) GetId() interface{} {
	return s.Id
}

type MenuVo struct {
	Name      string    `json:"name"`
	Component string    `json:"component"`
	FullPath  string    `json:"fullPath"`
	Meta      RouteMeta `json:"meta"`
	Children  []MenuVo  `json:"children"`
}

type RouteMeta struct {
	OrderNo  int    `json:"orderNo"`
	Title    string `json:"title"`
	Icon     string `json:"icon"`
	HideMenu bool   `json:"hideMenu"`
	IsLink   bool   `json:"isLink"`
}

type SysMenuGetReq struct {
	Id int `uri:"id"`
}

func (s *SysMenuGetReq) GetId() interface{} {
	return s.Id
}

type SysMenuDeleteReq struct {
	Ids []int `json:"ids"`
	base.ControlBy
}

func (s *SysMenuDeleteReq) GetId() interface{} {
	return s.Ids
}

type MenuLabel struct {
	Id       int         `json:"id,omitempty" gorm:"-"`
	Label    string      `json:"label,omitempty" gorm:"-"`
	Children []MenuLabel `json:"children,omitempty" gorm:"-"`
}

type MenuRole struct {
	models.SysMenu
	IsSelect bool `json:"is_select" gorm:"-"`
}

type SelectRole struct {
	RoleId int `uri:"roleId"`
}
