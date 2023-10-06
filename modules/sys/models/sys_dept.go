package models

import (
	"time"

	"gorm.io/gorm"
)

// SysDept
type SysDept struct {
	Id        int            `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"` //主键
	ParentId  int            `json:"parentId" gorm:"type:int unsigned;comment:父id"`                   // 上级部门
	DeptPath  string         `json:"deptPath" gorm:"type:varchar(255);comment:部门路径"`                  //路径
	Name      string         `json:"name" gorm:"type:varchar(128);comment:部门名"`                       //部门名称
	Type      int            `json:"type" gorm:"type:tinyint;comment:类型"`                             //状态
	Principal string         `json:"principal" gorm:"type:varchar(128);comment:部门领导"`                 //部门领导
	Phone     string         `json:"phone" gorm:"size:11;comment:手机号"`                                //手机号
	Email     string         `json:"email" gorm:"size:128;comment:邮箱"`                                //邮箱
	Sort      int            `json:"sort" gorm:"type:tinyint;comment:排序"`                             //排序
	Status    int            `json:"status" gorm:"type:tinyint;comment:状态"`                           //状态
	Remark    string         `json:"remark" gorm:"type:varchar(255);comment:备注"`                      //备注
	TeamId    int            `json:"teamId" gorm:"type:int unsigned;comment:团队id"`                    //团队id
	CreateBy  int            `json:"createBy" gorm:"type:int unsigned;comment:创建者"`                   //创建者
	UpdateBy  int            `json:"updateBy" gorm:"type:int unsigned;comment:更新者"`                   //更新者
	CreatedAt time.Time      `json:"createdAt" gorm:"type:datetime(3);comment:创建时间"`                  //创建时间
	UpdatedAt time.Time      `json:"updatedAt" gorm:"type:datetime(3);comment:最后更新时间"`                //最后更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`                                     //删除时间
}

func (SysDept) TableName() string {
	return "sys_dept"
}

func NewSysDept() *SysDept {
	return &SysDept{}
}

func (e *SysDept) SetId(id int) *SysDept {
	e.Id = id
	return e
}
func (e *SysDept) SetParentId(parentId int) *SysDept {
	e.ParentId = parentId
	return e
}
func (e *SysDept) SetDeptPath(deptPath string) *SysDept {
	e.DeptPath = deptPath
	return e
}
func (e *SysDept) SetDeptName(deptName string) *SysDept {
	e.Name = deptName
	return e
}
func (e *SysDept) SetSort(sort int) *SysDept {
	e.Sort = sort
	return e
}

func (e *SysDept) SetStatus(status int) *SysDept {
	e.Status = status
	return e
}
func (e *SysDept) SetCreateBy(createBy int) *SysDept {
	e.CreateBy = createBy
	return e
}
func (e *SysDept) SetUpdateBy(updateBy int) *SysDept {
	e.UpdateBy = updateBy
	return e
}
func (e *SysDept) SetCreatedAt(createdAt time.Time) *SysDept {
	e.CreatedAt = createdAt
	return e
}
func (e *SysDept) SetUpdatedAt(updatedAt time.Time) *SysDept {
	e.UpdatedAt = updatedAt
	return e
}
