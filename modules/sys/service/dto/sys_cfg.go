package dto

import (
	"github.com/baowk/dilu-core/core/base"
)

type SysCfgGetPageReq struct {
	base.ReqPage `search:"-"`
	Status       int `json:"status" form:"status"` //Status
}

// 配置
type SysCfgDto struct {
	Id     int    `json:"id"`     //主键
	Name   string `json:"name"`   //名字
	Key    string `json:"key"`    //key
	Value  string `json:"value"`  //Value
	Type   string `json:"type"`   //Type
	Remark string `json:"remark"` //Remark
	Status int    `json:"status"` //Status
}
