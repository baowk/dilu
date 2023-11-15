package models

import (
	"time"

	"gorm.io/gorm"
)

// SysRole
type SysRole struct {
	Id        int            `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"` //主键
	Name      string         `json:"name" gorm:"type:varchar(128);comment:角色名称"`                      //角色名称
	RoleKey   string         `json:"roleKey" gorm:"type:varchar(128);comment:角色代码"`                   //角色代码
	RoleSort  int            `json:"roleSort" gorm:"type:int unsigned;comment:排序"`                    //排序
	Status    int            `json:"status" gorm:"type:tinyint;comment:状态"`                           //状态
	TeamId    int            `json:"team_id" gorm:"type:tinyint(1);comment:团队"`                       //团队
	Remark    string         `json:"remark" gorm:"type:varchar(255);comment:备注"`                      //备注
	CreateBy  int            `json:"createBy" gorm:"type:int unsigned;comment:创建者"`                   //创建者
	UpdateBy  int            `json:"updateBy" gorm:"type:int unsigned;comment:更新者"`                   //更新者
	CreatedAt time.Time      `json:"createdAt" gorm:"type:datetime(3);comment:创建时间"`                  //创建时间
	UpdatedAt time.Time      `json:"updatedAt" gorm:"type:datetime(3);comment:最后更新时间"`                //最后更新时间
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`                                     //删除时间
	//SysMenu   *[]SysMenu     `json:"sysMenu" gorm:"many2many:sys_role_menu;foreignKey:Id;joinForeignKey:role_id;references:Id;joinReferences:menu_id;"`
}

func (SysRole) TableName() string {
	return "sys_role"
}

func NewSysRole() *SysRole {
	return &SysRole{}
}

func (e *SysRole) SetId(roleId int) *SysRole {
	e.Id = roleId
	return e
}
func (e *SysRole) SetName(roleName string) *SysRole {
	e.Name = roleName
	return e
}
func (e *SysRole) SetStatus(status int) *SysRole {
	e.Status = status
	return e
}
func (e *SysRole) SetRoleKey(roleKey string) *SysRole {
	e.RoleKey = roleKey
	return e
}
func (e *SysRole) SetRoleSort(roleSort int) *SysRole {
	e.RoleSort = roleSort
	return e
}

func (e *SysRole) SetRemark(remark string) *SysRole {
	e.Remark = remark
	return e
}

func (e *SysRole) SetCreateBy(createBy int) *SysRole {
	e.CreateBy = createBy
	return e
}
func (e *SysRole) SetUpdateBy(updateBy int) *SysRole {
	e.UpdateBy = updateBy
	return e
}
func (e *SysRole) SetCreatedAt(createdAt time.Time) *SysRole {
	e.CreatedAt = createdAt
	return e
}
func (e *SysRole) SetUpdatedAt(updatedAt time.Time) *SysRole {
	e.UpdatedAt = updatedAt
	return e
}
