package models

import (
	"time"
)

type SysApi struct {
	Id        int       `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键编码"` //主键
	Handle    string    `json:"handle" gorm:"size:128;comment:handle"`                             //handler
	Title     string    `json:"title" gorm:"size:128;comment:标题"`                                  //标题
	Action    string    `json:"action" gorm:"size:16;comment:请求类型"`                                //请求方法
	Path      string    `json:"path" gorm:"size:128;comment:请求地址"`                                 //路径
	Type      string    `json:"type" gorm:"size:16;comment:接口类型"`                                  //接口类型
	PermType  string    `json:"permType" gorm:"size:1;comment:权限类型（n：无需任何认证 t:须token p：须权限）"`      // n：无需任何认证 t:须token p：须权限
	Status    int       `json:"status" gorm:"type:tinyint;comment:状态 3 DEF 2 OK 1 del"`            //状态 3 DEF 2 OK 1 del
	UpdateBy  int       `json:"updateBy" gorm:"type:int unsigned;comment:更新者"`                     //更新者id
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`                                   //更新时间
}

func (SysApi) TableName() string {
	return "sys_api"
}
