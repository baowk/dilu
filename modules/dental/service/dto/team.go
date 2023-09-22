package dto

import (
	"github.com/baowk/dilu-core/core/base"
)

type TeamGetPageReq struct {
	base.ReqPage `search:"-"`
	Status       int `json:"status" form:"status"` //状态
}

// 团队
type TeamDto struct {
	Id       int    `json:"id"`       //主键
	ParentId int    `json:"parentId"` //上级团队
	Path     string `json:"path"`     //团队路径
	Name     string `json:"name"`     //团队名
	Owner    int    `json:"owner"`    //团队拥有者
	Status   int    `json:"status"`   //状态
}
