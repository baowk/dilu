package models

import (
	"time"

	"github.com/baowk/dilu-core/core/base"
)

// 接口列表
type SysApi struct {
	base.Model

	Title     string    `json:"title" gorm:"type:varchar(128);comment:标题"`                               //标题
	Method    string    `json:"method" gorm:"type:varchar(16);uniqueIndex:idx_method_path;comment:请求类型"` //请求类型
	Path      string    `json:"path" gorm:"type:varchar(128);uniqueIndex:idx_method_path;comment:请求地址"`  //请求地址
	PermType  int       `json:"permType" gorm:"tinyint unsigned;comment:权限类型（1：无需认证 2:须token 3：须鉴权）"`    //权限类型（1：无需认证 2:须token 3：须鉴权）
	Status    int       `json:"status" gorm:"type:tinyint unsigned;comment:状态 3 DEF 2 OK 1 del"`         //状态 3 DEF 2 OK 1 del
	UpdateBy  int       `json:"updateBy" gorm:"type:int unsigned;comment:更新者"`                           //更新者
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:datetime(3);comment:最后更新时间"`                        //最后更新时间
}

func (SysApi) TableName() string {
	return "sys_api"
}

func NewSysApi() *SysApi {
	return &SysApi{}
}

func (e *SysApi) SetId(id int) *SysApi {
	e.Id = id
	return e
}

func (e *SysApi) SetTitle(title string) *SysApi {
	e.Title = title
	return e
}

func (e *SysApi) SetMethod(method string) *SysApi {
	e.Method = method
	return e
}

func (e *SysApi) SetPath(path string) *SysApi {
	e.Path = path
	return e
}

func (e *SysApi) SetPermType(permType int) *SysApi {
	e.PermType = permType
	return e
}

func (e *SysApi) SetStatus(status int) *SysApi {
	e.Status = status
	return e
}

func (e *SysApi) SetUpdateBy(updateBy int) *SysApi {
	e.UpdateBy = updateBy
	return e
}

func (e *SysApi) SetUpdatedAt(updatedAt time.Time) *SysApi {
	e.UpdatedAt = updatedAt
	return e
}
