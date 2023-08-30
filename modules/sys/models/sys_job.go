package models

import "github.com/baowk/dilu-core/core/base"

type SysJob struct {
	JobId          int    `json:"jobId" gorm:"type:int unsigned;primaryKey;autoIncrement;common:主键"` // 主键
	JobName        string `json:"jobName" gorm:"size:255;common:名称"`                                 // 名称
	JobGroup       string `json:"jobGroup" gorm:"size:255;common:任务分组"`                              // 任务分组
	JobType        int    `json:"jobType" gorm:"size:1;common:任务分类"`                                 // 任务类型
	CronExpression string `json:"cronExpression" gorm:"size:255;common:cron表达式"`                     // cron表达式
	InvokeTarget   string `json:"invokeTarget" gorm:"size:255;调用目标"`                                 // 调用目标
	Args           string `json:"args" gorm:"size:255;common:目标参数"`                                  // 目标参数
	MisfirePolicy  int    `json:"misfirePolicy" gorm:"size:255;common:执行策略"`                         // 执行策略
	Concurrent     int    `json:"concurrent" gorm:"size:1;common:是否并发"`                              // 是否并发
	Status         int    `json:"status" gorm:"size:1;common:状态"`                                    // 状态
	EntryId        int    `json:"entry_id" gorm:"size:11;common:job启动时返回的id"`                        // job启动时返回的id
	base.ModelTime
	base.ControlBy
}

func (SysJob) TableName() string {
	return "sys_job"
}
