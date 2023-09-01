package models

import (
	"github.com/baowk/dilu-core/core/base"
)

type SysDept struct {
	DeptId    int       `json:"deptId" gorm:"type:int unsigned;primaryKey;autoIncrement;common:主键"` //部门编码
	ParentId  int       `json:"parentId" gorm:"type:int unsigned;common:父id"`                       //上级部门
	DeptPath  string    `json:"deptPath" gorm:"size:255;common:路径"`                                 //路径
	DeptName  string    `json:"deptName"  gorm:"size:128;common:部门名称"`                              //部门名称
	Sort      int       `json:"sort" gorm:"size:4;common:排序"`                                       //排序
	Leader    string    `json:"leader" gorm:"size:128;common:负责人"`                                  //负责人
	Phone     string    `json:"phone" gorm:"size:11;common:手机号"`                                    //手机
	Email     string    `json:"email" gorm:"size:64;common:邮箱"`                                     //邮箱
	Status    int       `json:"status" gorm:"size:1;common:状态"`                                     //状态
	DataScope string    `json:"dataScope" gorm:"-"`                                                 //数据域
	Params    string    `json:"params" gorm:"-"`                                                    //参数
	Children  []SysDept `json:"children" gorm:"-"`                                                  //子部门
	base.ControlBy
	base.ModelTime
}

func (SysDept) TableName() string {
	return "sys_dept"
}
