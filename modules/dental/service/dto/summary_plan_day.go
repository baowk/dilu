package dto

import (
    "github.com/baowk/dilu-core/core/base"
)

type SummaryPlanDayGetPageReq struct {
	base.ReqPage `search:"-"`
}

//每日总结计划
type SummaryPlanDayDto struct {
    
    Id int `json:"id"` //主键
    Day int `json:"day"` //天 
    TeamId int `json:"teamId"` //团队id 
    UserId int `json:"userId"` //用户id 
    Summary string `json:"summary"` //今日总结 
    Plan string `json:"plan"` //明日计划 
}



