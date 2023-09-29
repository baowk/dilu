package dto

import (
    "github.com/baowk/dilu-core/core/base"
)

type SysTeamGetPageReq struct {
	base.ReqPage `search:"-"`
    Status int `json:"status" form:"status"` //状态
}

//团队
type SysTeamDto struct {
    
    Id int `json:"id"` //主键
    Name string `json:"name"` //团队名 
    Owner int `json:"owner"` //团队拥有者 
    Status int `json:"status"` //状态 
}



