package service

import (
	"dilu/common/consts"
	"dilu/modules/dental/models"

	"github.com/baowk/dilu-core/core/base"
)

type SummaryPlanDayService struct {
	*base.BaseService
}

var SerSummaryPlanDay = SummaryPlanDayService{
	base.NewService(consts.DB_CRM),
}

func (s *SummaryPlanDayService) GetByDay(teamId, userId, day int, res *models.SummaryPlanDay) error {
	return s.GetByWhere(models.SummaryPlanDay{
		TeamId: teamId,
		UserId: userId,
		Day:    day,
	}, res)
}
