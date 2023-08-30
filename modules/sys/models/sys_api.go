package models

import (
	"time"

	"gorm.io/gorm"
)

type SysApi struct {
	Id        int            `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键编码"`
	Handle    string         `json:"handle" gorm:"size:128;comment:handle"`
	Title     string         `json:"title" gorm:"size:128;comment:标题"`
	Action    string         `json:"action" gorm:"size:16;comment:请求类型"`
	Path      string         `json:"path" gorm:"size:128;comment:请求地址"`
	Type      string         `json:"type" gorm:"size:16;comment:接口类型"`
	PermType  string         `json:"permType" gorm:"size:1;comment:权限类型（n：无需任何认证 t:须token p：须权限）"` // n：无需任何认证 t:须token p：须权限
	UpdateBy  int            `json:"updateBy" gorm:"type:int unsigned;comment:更新者"`                //更新者id
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:最后更新时间"`                              //更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`                                  //删除时间
}

func (SysApi) TableName() string {
	return "sys_api"
}
