package models

import "github.com/baowk/dilu-core/core/base"

type SysRole struct {
	RoleId    int       `json:"roleId" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:角色编码"` // 角色编码
	RoleName  string    `json:"roleName" gorm:"size:128;comment:角色名称"`                                 // 角色名称
	Status    int       `json:"status" gorm:"type:tinyint;comment:状态"`                                 //状态
	RoleKey   string    `json:"roleKey" gorm:"size:128;comment:角色代码"`                                  //角色代码
	RoleSort  int       `json:"roleSort" gorm:"type:int unsigned;comment:排序"`                          //角色排序
	Flag      string    `json:"flag" gorm:"size:128;comment:flag"`                                     //
	Remark    string    `json:"remark" gorm:"size:255;comment:备注"`                                     //备注
	Admin     bool      `json:"admin" gorm:"size:4;comment:管理员"`                                       //超管标识
	DataScope string    `json:"dataScope" gorm:"size:128;comment:数据权限"`
	SysMenu   []SysMenu `json:"sysMenu" gorm:"many2many:sys_role_menu;foreignKey:RoleId;joinForeignKey:role_id;references:Id;joinReferences:id;"`
	base.ControlBy
	base.ModelTime
}

func (SysRole) TableName() string {
	return "sys_role"
}
