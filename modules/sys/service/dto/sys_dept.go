package dto

import (
	"github.com/baowk/dilu-core/core/base"
)

type SysDeptGetPageReq struct {
	base.ReqPage `search:"-"`
	Status       int `json:"status" form:"status"` //
	TeamId       int `json:"teamId"`               //团队id
}

// SysDept
type SysDeptDto struct {
	DeptId    int    `json:"deptId"`    //主键
	ParentId  int    `json:"parentId"`  //
	DeptPath  string `json:"deptPath"`  //
	DeptName  string `json:"deptName"`  //
	Sort      int    `json:"sort"`      //
	Leader    string `json:"leader"`    //
	Phone     string `json:"phone"`     //
	Email     string `json:"email"`     //
	Status    int    `json:"status"`    //
	TeamId    int    `json:"teamId"`    //团队id
	Id        int    `json:"id" `       //主键
	Name      string `json:"name"`      //部门名称
	Type      int    `json:"type"`      //状态
	Principal string `json:"principal"` //部门领导
	Remark    string `json:"remark"`    //备注
}
