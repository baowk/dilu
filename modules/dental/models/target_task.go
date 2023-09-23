package models

//TargetTask
type TargetTask struct {
	Id             int `json:"id" gorm:"type:int unsigned;primaryKey;autoIncrement;comment:主键"` //主键
	Month          int `json:"month" gorm:"type:bigint;comment:月"`                              //月
	TeamId         int `json:"teamId" gorm:"type:bigint;comment:团队id"`                          //团队id
	UserId         int `json:"userId" gorm:"type:bigint;comment:用户id"`                          //用户id
	NewCustomerCnt int `json:"newCustomerCnt" gorm:"type:bigint;comment:留存任务"`                  //留存任务
	FirstDiagnosis int `json:"firstDiagnosis" gorm:"type:bigint;comment:导诊任务"`                  //导诊任务
	Deal           int `json:"deal" gorm:"type:bigint;comment:成交任务"`                            //成交任务
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
func (e *TargetTask) SetMonth(month int) *TargetTask {
	e.Month = month
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
