package models

// SysRole
type SysRoleMenu struct {
	RoleId int `json:"roleId" gorm:"type:int unsigned;comment:角色id"` //角色id
	MenuId int `json:"menuId" gorm:"type:int unsigned;comment:菜单id"` //菜单id
}

func (SysRoleMenu) TableName() string {
	return "sys_role_menu"
}
