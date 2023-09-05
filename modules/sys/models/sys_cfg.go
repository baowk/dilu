package models

import (
    "time"
    
    "github.com/baowk/dilu-core/core/base"
)

//SysCfg
type SysCfg struct {
    base.Model
    
    Name string `json:"name" gorm:"type:varchar(128);comment:名字"` //名字 
    Key string `json:"key" gorm:"type:varchar(128);comment:key"` //key 
    Value string `json:"value" gorm:"type:json;comment:Value"` //Value 
    Type string `json:"type" gorm:"type:varchar(64);comment:Type"` //Type 
    Remark string `json:"remark" gorm:"type:varchar(128);comment:Remark"` //Remark 
    Status int `json:"status" gorm:"type:tinyint;comment:Status"` //Status 
    UpdateBy int `json:"updateBy" gorm:"type:int unsigned;comment:更新者"` //更新者 
    UpdatedAt time.Time `json:"updatedAt" gorm:"type:datetime(3);comment:最后更新时间"` //最后更新时间 
}

func (SysCfg) TableName() string {
    return "sys_cfg"
}