package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type SysCfg struct {
	Id        int       `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键编码"` //主键编码
	Name      string    `json:"name" gorm:"type:varchar(128);comment:名字"`                          //名称
	Key       string    `json:"key" gorm:"size:128;comment:key"`                                   //key
	Value     []CfgData `json:"value" gorm:"type:json;default:(-);comment:Value"`                  //配置项
	Type      string    `json:"type" gorm:"size:64;comment:Type"`                                  //类型
	Remark    string    `json:"remark" gorm:"size:128;comment:Remark"`                             //备注
	Status    int       `json:"status" gorm:"size:4;comment:Status"`                               //状态
	UpdateBy  int       `json:"updateBy" gorm:"type:int unsigned;comment:更新者"`                     //更新者id
	UpdatedAt time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`                                   //更新时间
}

type CfgData struct {
	Label  string `json:"label" gorm:"size:128;comment:label"`   //配置项label
	Val    string `json:"val" gorm:"size:255;comment:Value"`     //配置项值
	Type   string `json:"type" gorm:"size:64;comment:Type"`      //类型
	IsDef  string `json:"isDef" gorm:"size:8;comment:IsDefault"` //是否默认值
	Status int    `json:"status" gorm:"size:1;comment:Status"`   //状态
	Remark string `json:"remark" gorm:"size:255;comment:Remark"` //备注
}

func (c CfgData) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *CfgData) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

func (m SysCfg) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (SysCfg) TableName() string {
	return "sys_cfg"
}
