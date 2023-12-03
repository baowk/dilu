package service

import (
	"dilu/common/consts"
	"dilu/modules/dental/models"
	"dilu/modules/dental/service/dto"
	senums "dilu/modules/sys/enums"
	smodels "dilu/modules/sys/models"
	"dilu/modules/sys/service"
	"time"

	"github.com/baowk/dilu-core/core/base"
)

type EventDayStService struct {
	*base.BaseService
}

var SerEventDaySt = EventDayStService{
	base.NewService(consts.DB_CRM),
}

func (s *EventDayStService) Create(teamId, userId int, reqId string, data *models.EventDaySt) error {
	if data.UserId == 0 {
		data.UserId = userId
	}
	if teamId != 0 {
		data.TeamId = teamId
	}
	var tu smodels.SysMember
	if err := service.SerSysMember.GetMember(data.TeamId, data.UserId, &tu); err != nil {
		return err
	}
	data.DeptPath = tu.DeptPath
	data.CreateBy = userId
	return s.BaseService.Create(data)
}

func (s *EventDayStService) Update(teamId, userId int, reqId string, data *models.EventDaySt) error {
	if data.UserId == 0 {
		data.UserId = userId
	}
	if teamId != 0 {
		data.TeamId = teamId
	}
	var tu smodels.SysMember
	if err := service.SerSysMember.GetMember(data.TeamId, data.UserId, &tu); err != nil {
		return err
	}
	data.DeptPath = tu.DeptPath
	data.UpdateBy = userId
	return s.BaseService.UpdateById(data)
}

func (s *EventDayStService) GetList(teamId, userId int, deptPath string, begin, end time.Time, list *[]models.EventDaySt) error {
	db := s.DB().Where("team_id = ?", teamId).Where("day >=?", begin).
		Where("day < ?", end)
	if userId > 0 {
		db.Where("user_id = ?", userId)
	} else if deptPath != "" {
		db.Where("dept_path like ?", deptPath+"%")
	}
	return db.Find(list).Error
}

func (s *EventDayStService) Page(teamId, userId int, req dto.EventDayStGetPageReq, list *[]models.EventDaySt, total *int64) error {
	var tu smodels.SysMember
	if err := service.SerSysMember.GetMember(teamId, userId, &tu); err != nil {
		return err
	}
	if tu.PostId == senums.Staff.Id {
		req.UserId = userId
	} else if tu.PostId > senums.Admin.Id {
		req.DeptPath = tu.DeptPath
	}

	db := s.DB().Where("team_id = ?", teamId)
	if !req.Begin.IsZero() {
		db.Where("day >=?", req.Begin)
	}
	if !req.End.IsZero() {
		db.Where("day < ?", req.End)
	}
	if req.UserId > 0 {
		db.Where("user_id = ?", req.UserId)
	} else if req.DeptPath != "" {
		db.Where("dept_path like ?", req.DeptPath+"%")
	}

	return db.Offset(req.GetOffset()).Limit(req.GetSize()).Order("day desc").Order("id desc").Find(list).
		Offset(-1).Limit(-1).Count(total).Error
}
