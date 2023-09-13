package dto

import (
    "github.com/baowk/dilu-core/core/base"
)

type SysRoleGetPageReq struct {
	base.ReqPage `search:"-"`
    Status int `json:"status" form:"status"` //状态
}

//SysRole
type SysRoleDto struct {
    
    RoleId int `json:"roleId"` //主键
    RoleName string `json:"roleName"` //角色名称 
    Status int `json:"status"` //状态 
    RoleKey string `json:"roleKey"` //角色代码 
    RoleSort int `json:"roleSort"` //排序 
    Flag string `json:"flag"` //flag 
    Remark string `json:"remark"` //备注 
    Admin int `json:"admin"` //管理员 
    DataScope string `json:"dataScope"` //数据权限 
}



