package service

import (
	"dilu/common/consts"
	"dilu/modules/dental/models"
	"dilu/modules/sys/service"
	"dilu/modules/sys/service/dto"

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
	var tu dto.TeamMemberResp
	if err := service.SerSysMember.GetTeamUser(data.TeamId, data.UserId, &tu); err != nil {
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
	var tu dto.TeamMemberResp
	if err := service.SerSysMember.GetTeamUser(data.TeamId, data.UserId, &tu); err != nil {
		return err
	}
	data.DeptPath = tu.DeptPath
	data.UpdateBy = userId
	return s.BaseService.UpdateById(data)
}
