package service

import (
	"dilu/common/consts"
	"dilu/modules/dental/enums"
	"dilu/modules/dental/models"

	"github.com/baowk/dilu-core/core/base"
)

type TargetTaskService struct {
	*base.BaseService
}

var SerTargetTask = TargetTaskService{
	base.NewService(consts.DB_CRM),
}

func (s *TargetTaskService) GetTasks(dayType enums.DayType, day int, teamId int, userId int, deptPath string, list *[]models.TargetTask) error {
	db := s.DB().Where("day_type = ?", dayType).Where("day = ?", day).Where("team_id = ? ", teamId)
	if userId > 0 {
		db.Where("user_id = ?", userId)
	} else if deptPath != "" {
		db.Where("dept_path like ?", deptPath+"%")
	}
	return db.Find(list).Error
}
