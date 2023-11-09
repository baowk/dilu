package dto

import (
	"dilu/modules/sys/models"

	"github.com/baowk/dilu-core/core/base"
)

type SysCfgGetPageReq struct {
	base.ReqPage `query:"-"`
	SortOrder    string `json:"-" query:"type:order;column:id"`
	SysCfgQuery
}

type SysCfgQuery struct {
	Status int `json:"status" query:""` //Status
}

func (SysCfgGetPageReq) TableName() string {
	return models.TBSysCfg
}

func (SysCfgQuery) TableName() string {
	return models.TBSysCfg
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
