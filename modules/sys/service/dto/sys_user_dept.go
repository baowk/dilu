package dto

import (
	"github.com/baowk/dilu-core/core/base"
)

type SysUserDeptGetPageReq struct {
	base.ReqPage `search:"-"`
	Status       int `json:"status" form:"status"` //状态 1正常
}

// SysUserDept
type SysUserDeptDto struct {
	Id      int `json:"id"`      //主键
	DeptId  int `json:"deptId"`  //部门id
	UserId  int `json:"userId"`  //上级部门
	PostTag int `json:"postTag"` //职位标签 1主管 2副主管 3员工
	Status  int `json:"status"`  //状态 1正常
}
