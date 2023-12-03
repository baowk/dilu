package dto

import (
	"github.com/baowk/dilu-core/core/base"
)

type TargetTaskGetPageReq struct {
	base.ReqPage `search:"-"`
	TeamId       int    `json:"teamId"`
	UserId       int    `json:"userId"`
	DeptPath     string `json:"deptPath"`
}

// TargetTask
type TargetTaskDto struct {
	Id             int    `json:"id"`             //主键
	DayType        int    `json:"dayType"`        //时间类型:月 30,周 7
	Day            int    `json:"day"`            //时间:202310
	TeamId         int    `json:"teamId"`         //团队id
	UserId         int    `json:"userId"`         //用户id
	DeptPath       string `json:"deptPath"`       //部门路径
	NewCustomerCnt int    `json:"newCustomerCnt"` //留存任务
	FirstDiagnosis int    `json:"firstDiagnosis"` //导诊任务
	Deal           int    `json:"deal"`           //成交任务
}
