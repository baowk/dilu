package models

import (
	"github.com/baowk/dilu-core/core/base"
)

type SysSms struct {
	base.Model
	Phone     string `json:"phone" gorm:"type:varchar(16);comment:手机号"`  //手机号
	Code      string `json:"code" gorm:"type:varchar(6);comment:验证码"`    //验证码
	Type      string `json:"type" gorm:"type:varchar(6);comment:类型"`     //类型，暂未定义
	Status    int    `json:"status" gorm:"type:tinyint;comment:状态"`      //发送状态 1成功
	UseStatus int    `json:"useStatus" gorm:"type:tinyint;comment:使用状态"` //使用状态 1已使用
	base.ModelIntTime
}

func (SysSms) TableName() string {
	return "sys_sms"
}
