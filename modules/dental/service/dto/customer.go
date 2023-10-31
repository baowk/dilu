package dto

import (
	"github.com/baowk/dilu-core/core/base"
)

type CustomerGetPageReq struct {
	base.ReqPage `search:"-"`
	UserId       int    `json:"userId"` //用户id
	TeamId       int    `json:"teamId"` //团队id
	DeptPath     string `json:"deptPath"`
}

// Customer
type CustomerDto struct {
	Id       int    `json:"id"`       //主键
	Name     string `json:"name"`     //姓名
	Birthday int    `json:"birthday"` //生日
	Phone    string `json:"phone"`    //手机号
	Wechat   string `json:"wechat"`   //微信号
	Gender   int    `json:"gender"`   //性别
	Address  string `json:"address"`  //地址
	Remark   string `json:"remark"`   //描述
	UserId   int    `json:"userId"`   //用户id
	TeamId   int    `json:"teamId"`   //团队id
}
