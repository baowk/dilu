package dto

import (
	"github.com/baowk/dilu-core/core/base"
)

type SysApiGetPageReq struct {
	base.ReqPage `query:"-"`
	Status       int `form:"status"  query:"type:eq;column:status;table:sys_api"` //状态 3 DEF 2 OK 1 del
}

// 接口列表
type SysApiDto struct {
	base.Model

	Title    string `json:"title"`    //标题
	Method   string `json:"method"`   //请求类型
	Path     string `json:"path"`     //请求地址
	Type     string `json:"type"`     //接口类型
	PermType string `json:"permType"` //权限类型（n：无需任何认证 t:须token p：须权限）
	Status   int    `json:"status"`   //状态 3 DEF 2 OK 1 del
}
