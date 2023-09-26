package dto

import (
	"github.com/baowk/dilu-core/core/base"
)

type SysDeptGetPageReq struct {
	base.ReqPage `search:"-"`
	Status       int `json:"status" form:"status"` //
}

// SysDept
type SysDeptDto struct {
	DeptId   int    `json:"deptId"`   //主键
	ParentId int    `json:"parentId"` //
	DeptPath string `json:"deptPath"` //
	DeptName string `json:"deptName"` //
	Sort     int    `json:"sort"`     //
	Leader   string `json:"leader"`   //
	Phone    string `json:"phone"`    //
	Email    string `json:"email"`    //
	Status   int    `json:"status"`   //
}
