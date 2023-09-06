package models

import (
	"time"

	"github.com/baowk/dilu-core/core/base"
)

// SysApi
type SysApi struct {
	base.Model

	//Handle    string    `json:"handle" gorm:"type:varchar(128);comment:handle"`                        //handle
	Title     string    `json:"title" gorm:"type:varchar(128);comment:标题"`                             //标题
	Action    string    `json:"action" gorm:"type:varchar(16);comment:请求类型"`                           //请求类型
	Path      string    `json:"path" gorm:"type:varchar(128);comment:请求地址"`                            //请求地址
	Type      string    `json:"type" gorm:"type:varchar(16);comment:接口类型"`                             //接口类型
	PermType  string    `json:"permType" gorm:"type:varchar(1);comment:权限类型（n：无需任何认证 t:须token p：须权限）"` //权限类型（n：无需任何认证 t:须token p：须权限）
	Status    int       `json:"status" gorm:"type:tinyint;comment:状态 3 DEF 2 OK 1 del"`                //状态 3 DEF 2 OK 1 del
	UpdateBy  int       `json:"updateBy" gorm:"type:int unsigned;comment:更新者"`                         //更新者
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:datetime(3);comment:最后更新时间"`                      //最后更新时间
}

func (SysApi) TableName() string {
	return "sys_api"
}
