package dto

import (
	"github.com/baowk/dilu-core/core/base"
)


type SysCfgGetPageReq struct {
	base.ReqPage `search:"-"`
    Status int `form:"status"  search:"type:exact;column:status;table:sys_cfg" comment:"Status"` //Status
}




//SysCfg
type SysCfgDto struct {
    base.Model
    
    Name string `json:"name"` //名字 
    Key string `json:"key"` //key 
    Value string `json:"value"` //Value 
    Type string `json:"type"` //Type 
    Remark string `json:"remark"` //Remark 
    Status int `json:"status"` //Status 
}



