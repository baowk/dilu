package dto

import (
	"dilu/modules/notice/models"
	"time"

	"github.com/baowk/dilu-core/core/base"
)

type TaskGetPageReq struct {
	base.ReqPage `query:"-"`
	SortOrder    string `json:"-" query:"type:order;column:id"`
	Status       int    `json:"status" query:"column:status"` //状态1开启2关闭ReminderStatus int `json:"reminderStatus" query:"column:reminder_status"` //提醒状态 1开启 2关闭
	TeamId       int    `json:"teamId" query:""`
	UserId       int    `json:"userId" query:""`
}

func (TaskGetPageReq) TableName() string {
	return models.TBTask
}

// Task
type TaskDto struct {
	Id             int       `json:"id"`             //主键
	TeamId         int       `json:"teamId"`         //团队id
	UserId         int       `json:"userId"`         //用户id
	Title          string    `json:"title"`          //任务标题
	Content        string    `json:"content"`        //任务内容
	TaskType       int       `json:"taskType"`       //任务类型
	Op             int       `json:"op"`             //操作类型
	OpId           int       `json:"opId"`           //操作id
	BeginAt        time.Time `json:"beginAt"`        //开始时间
	EndAt          time.Time `json:"endAt"`          //结束时间
	ReminderTime   time.Time `json:"reminderTime"`   //提醒时间
	Status         int       `json:"status"`         //状态1开启2关闭
	ReminderStatus int       `json:"reminderStatus"` //提醒状态 1开启 2关闭
}
