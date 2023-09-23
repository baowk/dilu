package dto

import (
	"github.com/baowk/dilu-core/core/base"
)

type TeamMemberGetPageReq struct {
	base.ReqPage `search:"-"`
	Status       int `json:"status" form:"status"` //状态
}

// TeamMember
type TeamMemberDto struct {
	Id     int    `json:"id"`     //主键
	UserId int    `json:"userId"` //用户id
	TeamId int    `json:"teamId"` //团队id
	Name   string `json:"name"`   //姓名
	Phone  string `json:"phone"`  //电话
	Gender int    `json:"gender"` //性别
	Status int    `json:"status"` //状态
	Role   int    `json:"role"`   //角色 1主管 2副主管 4普通
}
