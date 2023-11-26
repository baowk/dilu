package models

import (
	"time"

	"gorm.io/gorm"
)

// Task
type Task struct {
	Id             int            `json:"id" gorm:"type:bigint unsigned;primaryKey;autoIncrement;comment:主键"` //主键
	TeamId         int            `json:"teamId" gorm:"type:int;comment:团队id"`                                //团队id
	UserId         int            `json:"userId" gorm:"type:int;comment:用户id"`                                //用户id
	Title          string         `json:"title" gorm:"type:varchar(255);comment:任务标题"`                        //任务标题
	Content        string         `json:"content" gorm:"type:varchar(1024);comment:任务内容"`                     //任务内容
	TaskType       int            `json:"taskType" gorm:"type:tinyint;comment:任务类型"`                          //任务类型
	Op             int            `json:"op" gorm:"type:int;comment:操作类型"`                                    //操作类型
	OpId           int            `json:"opId" gorm:"type:int;comment:操作id"`                                  //操作id
	BeginAt        time.Time      `json:"beginAt" gorm:"type:datetime;comment:开始时间"`                          //开始时间
	EndAt          time.Time      `json:"endAt" gorm:"type:datetime;comment:结束时间"`                            //结束时间
	ReminderTime   time.Time      `json:"reminderTime" gorm:"type:datetime;comment:提醒时间"`                     //提醒时间
	Status         int            `json:"status" gorm:"type:tinyint;comment:状态1开启2关闭"`                        //状态1开启2关闭
	ReminderStatus int            `json:"reminderStatus" gorm:"type:tinyint;comment:提醒状态 1开启 2关闭"`            //提醒状态 1开启 2关闭
	CreatedAt      time.Time      `json:"createdAt" gorm:"type:datetime;comment:创建时间"`                        //创建时间
	UpdatedAt      time.Time      `json:"updatedAt" gorm:"type:datetime;comment:更新时间"`                        //更新时间
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`                                        //删除时间
}

const TBTask = "task"

func (Task) TableName() string {
	return TBTask
}

func NewTask() *Task {
	return &Task{}
}
