package models

import (
	"github.com/baowk/dilu-core/core/base"
)

type ThirdLogin struct {
	base.Model
	UserId    int    `json:"userId" gorm:"type:int unsigned;comment:用户id"`               //用户id
	Platform  int    `json:"platform" gorm:"type:tinyint unsigned;comment:平台 1 微信 2 钉钉"` //平台 1微信 2钉钉
	OpenId    string `json:"openId" gorm:"type:varchar(128);comment:第三方open_id"`         //第三方openid userid
	UnionId   string `json:"unionId" gorm:"type:varchar(128);comment:第三方union_id"`       //第三方union_id多平台使用
	ThirdData string `json:"thirdData" gorm:"type:json;comment:第三方返回数据"`                 //第三方返回数据
	base.ModelIntTime
}

func (ThirdLogin) TableName() string {
	return "third_login"
}
