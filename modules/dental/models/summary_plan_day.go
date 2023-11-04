package models

import (
	"time"
)

// SummaryPlanDay
type SummaryPlanDay struct {
	Id        int       `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"` //主键
	Day       time.Time `json:"day" gorm:"type:date;comment:天"`                                  //天
	TeamId    int       `json:"teamId" gorm:"type:int unsigned;comment:团队id"`                    //团队id
	UserId    int       `json:"userId" gorm:"type:int unsigned;comment:用户id"`                    //用户id
	DeptPath  string    `json:"deptPath" gorm:"type:varchar(255);comment:部门路径"`                  //路径
	Summary   string    `json:"summary" gorm:"type:text;comment:今日总结"`                           //今日总结
	Plan      string    `json:"plan" gorm:"type:text;comment:明日计划"`                              //明日计划
	CreatedAt time.Time `json:"createdAt" gorm:"type:datetime;comment:创建时间"`                     //创建时间
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:datetime;comment:更新时间"`                     //更新时间
	CreateBy  int       `json:"createBy" gorm:"type:int unsigned;index;comment:创建者"`             //创建者id
	UpdateBy  int       `json:"updateBy" gorm:"type:int unsigned;index;comment:更新者"`             //更新者id
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
func (e *SummaryPlanDay) SetDay(day time.Time) *SummaryPlanDay {
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
