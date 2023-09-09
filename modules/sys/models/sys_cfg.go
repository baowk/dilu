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

func NewSysCfg() *SysCfg{
    return &SysCfg{}
}





func (e *SysCfg) SetId(id int) *SysCfg {
	e.Id = id
	return e
}



func (e *SysCfg) SetName(name string) *SysCfg {
	e.Name = name
	return e
}



func (e *SysCfg) SetKey(key string) *SysCfg {
	e.Key = key
	return e
}



func (e *SysCfg) SetValue(value string) *SysCfg {
	e.Value = value
	return e
}



        

func (e *SysCfg) SetType(aType string) *SysCfg {
	e.Type = aType
	return e
}



func (e *SysCfg) SetRemark(remark string) *SysCfg {
	e.Remark = remark
	return e
}



func (e *SysCfg) SetStatus(status int) *SysCfg {
	e.Status = status
	return e
}



func (e *SysCfg) SetUpdateBy(updateBy int) *SysCfg {
	e.UpdateBy = updateBy
	return e
}



func (e *SysCfg) SetUpdatedAt(updatedAt time.Time) *SysCfg {
	e.UpdatedAt = updatedAt
	return e
}


