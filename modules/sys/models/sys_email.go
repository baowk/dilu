package models

import (
	"github.com/baowk/dilu-core/core/base"
)

type SysEmail struct {
	base.Model

	Email     string `json:"email" gorm:"type:varchar(255);comment:邮箱地址"` //邮箱地址
	Code      string `json:"code" gorm:"type:varchar(6);comment:验证码"`     //验证码
	Type      string `json:"type" gorm:"type:varchar(6);comment:类型"`      //类型 暂未定义
	Status    int    `json:"status" gorm:"type:tinyint;comment:状态"`       //状态 1 发送成功
	UseStatus int    `json:"useStatus" gorm:"type:tinyint;comment:使用状态"`  //使用状态 1已使用
	base.ModelIntTime
}

func (SysEmail) TableName() string {
	return "sys_email"
}
