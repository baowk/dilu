package dto

import (
	"github.com/baowk/dilu-core/core/base"
)

type SysJobGetPageReq struct {
	base.ReqPage `search:"-"`
	Status       int `json:"status" form:"status"` //
}

// SysJob
type SysJobDto struct {
	JobId          int    `json:"jobId"`          //主键
	JobName        string `json:"jobName"`        //
	JobGroup       string `json:"jobGroup"`       //
	JobType        int    `json:"jobType"`        //
	CronExpression string `json:"cronExpression"` //
	InvokeTarget   string `json:"invokeTarget"`   //
	Args           string `json:"args"`           //
	MisfirePolicy  int    `json:"misfirePolicy"`  //
	Concurrent     int    `json:"concurrent"`     //
	Status         int    `json:"status"`         //
	EntryId        int    `json:"entryId"`        //
}
