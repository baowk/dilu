package dto

import (
    "github.com/baowk/dilu-core/core/base"
)

type SysUserDeptGetPageReq struct {
	base.ReqPage `search:"-"`
    Status int `json:"status" form:"status"` //Status
}

//SysUserDept
type SysUserDeptDto struct {
    
    Id int `json:"id"` //主键
    DeptId int `json:"deptId"` //部门id 
    UserId int `json:"userId"` //ParentId 
    PostTag int `json:"postTag"` //DeptPath 
    Status int `json:"status"` //Status 
}



