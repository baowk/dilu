package models

import (
	"time"

	"gorm.io/gorm"
)

// SysUserDept
type SysUserDept struct {
	Id        int            `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"` //主键
	DeptId    int            `json:"deptId" gorm:"type:int unsigned;comment:部门id"`                    //部门id
	UserId    int            `json:"userId" gorm:"type:int unsigned;comment:上级部门"`                    //上级部门
	PostTag   int            `json:"postTag" gorm:"type:tinyint unsigned;comment:职位标签 1主管 2副主管 3员工"`  //职位标签 1主管 2副主管 3员工
	Status    int            `json:"status" gorm:"type:tinyint;comment:状态 1正常 "`                      //状态 1正常
	CreateBy  int            `json:"createBy" gorm:"type:int unsigned;comment:创建者"`                   //创建者
	UpdateBy  int            `json:"updateBy" gorm:"type:int unsigned;comment:更新者"`                   //更新者
	CreatedAt time.Time      `json:"createdAt" gorm:"type:datetime(3);comment:创建时间"`                  //创建时间
	UpdatedAt time.Time      `json:"updatedAt" gorm:"type:datetime(3);comment:最后更新时间"`                //最后更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`                                     //删除时间
}

func (SysUserDept) TableName() string {
	return "sys_user_dept"
}

func NewSysUserDept() *SysUserDept {
	return &SysUserDept{}
}

func (e *SysUserDept) SetId(id int) *SysUserDept {
	e.Id = id
	return e
}
func (e *SysUserDept) SetDeptId(deptId int) *SysUserDept {
	e.DeptId = deptId
	return e
}
func (e *SysUserDept) SetUserId(userId int) *SysUserDept {
	e.UserId = userId
	return e
}
func (e *SysUserDept) SetPostTag(postTag int) *SysUserDept {
	e.PostTag = postTag
	return e
}
func (e *SysUserDept) SetStatus(status int) *SysUserDept {
	e.Status = status
	return e
}
func (e *SysUserDept) SetCreateBy(createBy int) *SysUserDept {
	e.CreateBy = createBy
	return e
}
func (e *SysUserDept) SetUpdateBy(updateBy int) *SysUserDept {
	e.UpdateBy = updateBy
	return e
}
func (e *SysUserDept) SetCreatedAt(createdAt time.Time) *SysUserDept {
	e.CreatedAt = createdAt
	return e
}
func (e *SysUserDept) SetUpdatedAt(updatedAt time.Time) *SysUserDept {
	e.UpdatedAt = updatedAt
	return e
}
