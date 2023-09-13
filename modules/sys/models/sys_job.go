package models

import (
            "gorm.io/gorm"
    "time"
    
)

//SysJob
type SysJob struct {
    
    JobId int `json:"jobId" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"` //主键
    JobName string `json:"jobName" gorm:"type:varchar(255);comment:JobName"` // 
    JobGroup string `json:"jobGroup" gorm:"type:varchar(255);comment:JobGroup"` // 
    JobType int `json:"jobType" gorm:"type:tinyint;comment:JobType"` // 
    CronExpression string `json:"cronExpression" gorm:"type:varchar(255);comment:CronExpression"` // 
    InvokeTarget string `json:"invokeTarget" gorm:"type:varchar(255);comment:InvokeTarget"` // 
    Args string `json:"args" gorm:"type:varchar(255);comment:Args"` // 
    MisfirePolicy int `json:"misfirePolicy" gorm:"type:bigint;comment:MisfirePolicy"` // 
    Concurrent int `json:"concurrent" gorm:"type:tinyint;comment:Concurrent"` // 
    Status int `json:"status" gorm:"type:tinyint;comment:Status"` // 
    EntryId int `json:"entryId" gorm:"type:smallint;comment:EntryId"` // 
    CreatedAt time.Time `json:"createdAt" gorm:"type:datetime(3);comment:创建时间"` //创建时间 
    UpdatedAt time.Time `json:"updatedAt" gorm:"type:datetime(3);comment:最后更新时间"` //最后更新时间 
    DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`     //删除时间
    CreateBy int `json:"createBy" gorm:"type:int unsigned;comment:创建者"` //创建者 
    UpdateBy int `json:"updateBy" gorm:"type:int unsigned;comment:更新者"` //更新者 
}

func (SysJob) TableName() string {
    return "sys_job"
}

func NewSysJob() *SysJob{
    return &SysJob{}
}

func (e *SysJob) SetJobId(jobId int) *SysJob {
	e.JobId = jobId
	return e
}
func (e *SysJob) SetJobName(jobName string) *SysJob {
	e.JobName = jobName
	return e
}
func (e *SysJob) SetJobGroup(jobGroup string) *SysJob {
	e.JobGroup = jobGroup
	return e
}
func (e *SysJob) SetJobType(jobType int) *SysJob {
	e.JobType = jobType
	return e
}
func (e *SysJob) SetCronExpression(cronExpression string) *SysJob {
	e.CronExpression = cronExpression
	return e
}
func (e *SysJob) SetInvokeTarget(invokeTarget string) *SysJob {
	e.InvokeTarget = invokeTarget
	return e
}
func (e *SysJob) SetArgs(args string) *SysJob {
	e.Args = args
	return e
}
func (e *SysJob) SetMisfirePolicy(misfirePolicy int) *SysJob {
	e.MisfirePolicy = misfirePolicy
	return e
}
func (e *SysJob) SetConcurrent(concurrent int) *SysJob {
	e.Concurrent = concurrent
	return e
}
func (e *SysJob) SetStatus(status int) *SysJob {
	e.Status = status
	return e
}
func (e *SysJob) SetEntryId(entryId int) *SysJob {
	e.EntryId = entryId
	return e
}
func (e *SysJob) SetCreatedAt(createdAt time.Time) *SysJob {
	e.CreatedAt = createdAt
	return e
}
func (e *SysJob) SetUpdatedAt(updatedAt time.Time) *SysJob {
	e.UpdatedAt = updatedAt
	return e
}
func (e *SysJob) SetCreateBy(createBy int) *SysJob {
	e.CreateBy = createBy
	return e
}
func (e *SysJob) SetUpdateBy(updateBy int) *SysJob {
	e.UpdateBy = updateBy
	return e
}

