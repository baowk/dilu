package models

import (
	"time"
)

// TargetTask
type TargetTask struct {
	Id             int       `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"` //主键
	DayType        int       `json:"dayType" gorm:"type:tinyint unsigned;comment:时间类型:月 30,周 7"`      //时间类型:月 30,周 7
	Day            int       `json:"day" gorm:"type:int unsigned;comment:时间:202310"`                  //时间:202310
	TeamId         int       `json:"teamId" gorm:"type:int unsigned;comment:团队id"`                    //团队id
	UserId         int       `json:"userId" gorm:"type:int unsigned;comment:用户id"`                    //用户id
	DeptPath       string    `json:"deptPath" gorm:"type:varchar(255);comment:部门路径"`                  //部门路径
	TaskType       int       `json:"taskType"  gorm:"type:int unsigned;comment:任务类型 1正式 算人员数量"`       //任务类型 1正式 算人员数量
	NewCustomerCnt int       `json:"newCustomerCnt" gorm:"type:int unsigned;comment:留存任务"`            //留存任务
	FirstDiagnosis int       `json:"firstDiagnosis" gorm:"type:int unsigned;comment:导诊任务"`            //导诊任务
	Deal           int       `json:"deal" gorm:"type:int unsigned;comment:成交任务"`                      //成交任务
	CreateBy       int       `json:"createBy" gorm:"type:int unsigned;comment:创建者"`                   //创建者
	UpdateBy       int       `json:"updateBy" gorm:"type:int unsigned;comment:更新者"`                   //更新者
	CreatedAt      time.Time `json:"createdAt" gorm:"type:datetime;comment:创建时间"`                     //创建时间
	UpdatedAt      time.Time `json:"updatedAt" gorm:"type:datetime;comment:更新时间"`                     //更新时间
}

func (TargetTask) TableName() string {
	return "target_task"
}

func NewTargetTask() *TargetTask {
	return &TargetTask{}
}

func (e *TargetTask) SetId(id int) *TargetTask {
	e.Id = id
	return e
}
func (e *TargetTask) SetDayType(dayType int) *TargetTask {
	e.DayType = dayType
	return e
}
func (e *TargetTask) SetDay(day int) *TargetTask {
	e.Day = day
	return e
}
func (e *TargetTask) SetTeamId(teamId int) *TargetTask {
	e.TeamId = teamId
	return e
}
func (e *TargetTask) SetUserId(userId int) *TargetTask {
	e.UserId = userId
	return e
}
func (e *TargetTask) SetDeptPath(deptPath string) *TargetTask {
	e.DeptPath = deptPath
	return e
}
func (e *TargetTask) SetNewCustomerCnt(newCustomerCnt int) *TargetTask {
	e.NewCustomerCnt = newCustomerCnt
	return e
}
func (e *TargetTask) SetFirstDiagnosis(firstDiagnosis int) *TargetTask {
	e.FirstDiagnosis = firstDiagnosis
	return e
}
func (e *TargetTask) SetDeal(deal int) *TargetTask {
	e.Deal = deal
	return e
}
func (e *TargetTask) SetCreateBy(createBy int) *TargetTask {
	e.CreateBy = createBy
	return e
}
func (e *TargetTask) SetUpdateBy(updateBy int) *TargetTask {
	e.UpdateBy = updateBy
	return e
}
func (e *TargetTask) SetCreatedAt(createdAt time.Time) *TargetTask {
	e.CreatedAt = createdAt
	return e
}
func (e *TargetTask) SetUpdatedAt(updatedAt time.Time) *TargetTask {
	e.UpdatedAt = updatedAt
	return e
}
