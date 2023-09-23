package dto

import (
	"github.com/baowk/dilu-core/core/base"
)

type TargetTaskGetPageReq struct {
	base.ReqPage `search:"-"`
}

// TargetTask
type TargetTaskDto struct {
	Id             int `json:"id"`             //主键
	Month          int `json:"month"`          //月
	TeamId         int `json:"teamId"`         //团队id
	UserId         int `json:"userId"`         //用户id
	NewCustomerCnt int `json:"newCustomerCnt"` //留存任务
	FirstDiagnosis int `json:"firstDiagnosis"` //导诊任务
	Deal           int `json:"deal"`           //成交任务
}
