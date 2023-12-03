package dto

import (
	"time"

	"github.com/baowk/dilu-core/core/base"
)

type SummaryPlanDayGetPageReq struct {
	base.ReqPage `search:"-"`
	TeamId       int    `json:"teamId"`
	UserId       int    `json:"userId"`
	DeptPath     string `json:"deptPath"`
}

// SummaryPlanDay
type SummaryPlanDayDto struct {
	Id      int       `json:"id"`      //主键
	Day     time.Time `json:"day"`     //天
	TeamId  int       `json:"teamId"`  //团队id
	UserId  int       `json:"userId"`  //用户id
	Summary string    `json:"summary"` //今日总结
	Plan    string    `json:"plan"`    //明日计划
}
