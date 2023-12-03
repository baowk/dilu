package service

import (
	"dilu/common/consts"
	"dilu/modules/dental/enums"
	"dilu/modules/dental/models"
	"dilu/modules/dental/service/dto"
	senums "dilu/modules/sys/enums"
	smodels "dilu/modules/sys/models"
	"dilu/modules/sys/service"

	"github.com/baowk/dilu-core/core/base"
)

type TargetTaskService struct {
	*base.BaseService
}

var SerTargetTask = TargetTaskService{
	base.NewService(consts.DB_CRM),
}

func (s *TargetTaskService) Page(req dto.TargetTaskGetPageReq, teamId, userId int, list *[]models.TargetTask, total *int64) error {
	var tu smodels.SysMember
	if err := service.SerSysMember.GetMember(teamId, userId, &tu); err != nil {
		return err
	}
	if tu.PostId == senums.Staff.Id {
		req.UserId = userId
	} else if tu.PostId > senums.Admin.Id {
		req.DeptPath = tu.DeptPath
	}
	db := s.DB().Where("team_id = ?", teamId).Order("id desc").Limit(req.GetSize()).
		Offset(req.GetOffset())
	if req.UserId > 0 {
		db.Where("user_id =?", req.UserId)
	} else if req.DeptPath != "" {
		db.Where("dept_path like?", req.DeptPath+"%")
	}
	return db.Find(list).Offset(-1).Limit(-1).Count(total).Error
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

func (s *TargetTaskService) Create(teamId, userId int, data *models.TargetTask) error {
	data.TeamId = teamId
	if data.UserId < 1 {
		data.UserId = userId
	}
	var tu smodels.SysMember
	if err := service.SerSysMember.GetMember(data.TeamId, data.UserId, &tu); err != nil {
		return err
	}
	data.DeptPath = tu.DeptPath
	data.CreateBy = userId
	return s.BaseService.Create(data)
}

func (s *TargetTaskService) Update(teamId, userId int, data *models.TargetTask) error {
	data.TeamId = teamId
	if data.UserId < 1 {
		data.UserId = userId
	}
	var tu smodels.SysMember
	if err := service.SerSysMember.GetMember(data.TeamId, data.UserId, &tu); err != nil {
		return err
	}
	data.DeptPath = tu.DeptPath
	data.UpdateBy = userId
	return s.BaseService.UpdateById(data)
}
