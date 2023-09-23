package models

import (
	"time"
)

// SummaryPlanDay
type SummaryPlanDay struct {
	Id        int       `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"` //主键
	Day       int       `json:"day" gorm:"type:bigint;comment:天"`                                //天
	TeamId    int       `json:"teamId" gorm:"type:bigint;comment:团队id"`                          //团队id
	UserId    int       `json:"userId" gorm:"type:bigint;comment:用户id"`                          //用户id
	Summary   string    `json:"summary" gorm:"type:text;comment:今日总结"`                           //今日总结
	Plan      string    `json:"plan" gorm:"type:text;comment:明日计划"`                              //明日计划
	CreatedAt time.Time `json:"createdAt" gorm:"type:datetime;comment:创建时间"`                     //创建时间
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:datetime;comment:更新时间"`                     //更新时间
}

func (SummaryPlanDay) TableName() string {
	return "summary_plan_day"
}

func NewSummaryPlanDay() *SummaryPlanDay {
	return &SummaryPlanDay{}
}

func (e *SummaryPlanDay) SetId(id int) *SummaryPlanDay {
	e.Id = id
	return e
}
func (e *SummaryPlanDay) SetDay(day int) *SummaryPlanDay {
	e.Day = day
	return e
}
func (e *SummaryPlanDay) SetTeamId(teamId int) *SummaryPlanDay {
	e.TeamId = teamId
	return e
}
func (e *SummaryPlanDay) SetUserId(userId int) *SummaryPlanDay {
	e.UserId = userId
	return e
}
func (e *SummaryPlanDay) SetSummary(summary string) *SummaryPlanDay {
	e.Summary = summary
	return e
}
func (e *SummaryPlanDay) SetPlan(plan string) *SummaryPlanDay {
	e.Plan = plan
	return e
}
func (e *SummaryPlanDay) SetCreatedAt(createdAt time.Time) *SummaryPlanDay {
	e.CreatedAt = createdAt
	return e
}
func (e *SummaryPlanDay) SetUpdatedAt(updatedAt time.Time) *SummaryPlanDay {
	e.UpdatedAt = updatedAt
	return e
}
