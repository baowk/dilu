package dto

import (
	"github.com/baowk/dilu-core/core/base"
)


type SysApiGetPageReq struct {
	base.ReqPage `search:"-"`
    Status int `form:"status"  search:"type:exact;column:status;table:sys_api" comment:"状态 3 DEF 2 OK 1 del"` //状态 3 DEF 2 OK 1 del
}




//SysApi
type SysApiDto struct {
    base.Model
    
    Title string `json:"title"` //标题 
    Action string `json:"action"` //请求类型 
    Path string `json:"path"` //请求地址 
    Type string `json:"type"` //接口类型 
    PermType string `json:"permType"` //权限类型（n：无需任何认证 t:须token p：须权限） 
    Status int `json:"status"` //状态 3 DEF 2 OK 1 del 
}



