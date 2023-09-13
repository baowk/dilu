package models

import (
	"time"

	"gorm.io/gorm"
)

// SysDept
type SysUserDept struct {
	Id        int            `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"` //主键
	DeptId    int            `json:"deptId" gorm:"type:int unsigned;comment:部门id"`                    //部门id
	UserId    int            `json:"userId" gorm:"type:int unsigned;comment:ParentId"`                //用户id
	PostTag   int            `json:"postTag" gorm:"type:tinyint unsigned;comment:DeptPath"`           //职位标记
	Status    int            `json:"status" gorm:"type:tinyint;comment:Status"`                       //状态
	CreateBy  int            `json:"createBy" gorm:"type:int unsigned;comment:创建者"`                   //创建者
	UpdateBy  int            `json:"updateBy" gorm:"type:int unsigned;comment:更新者"`                   //更新者
	CreatedAt time.Time      `json:"createdAt" gorm:"type:datetime(3);comment:创建时间"`                  //创建时间
	UpdatedAt time.Time      `json:"updatedAt" gorm:"type:datetime(3);comment:最后更新时间"`                //最后更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`                                     //删除时间
}

func (SysUserDept) TableName() string {
	return "sys_user_dept"
}
