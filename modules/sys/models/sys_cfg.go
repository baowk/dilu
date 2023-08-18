package models

import "encoding/json"

type SysCfg struct {
	Id   int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Name string `json:"name" gorm:"type:varchar(128);comment:名字"`
}

func (m SysCfg) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (SysCfg) TableName() string {
	return "sys_cfg"
}
