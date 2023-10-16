package service

import (
	"dilu/common/consts"

	"github.com/baowk/dilu-core/core/base"
)

type SummaryPlanDayService struct {
	*base.BaseService
}

var SerSummaryPlanDay = SummaryPlanDayService{
	base.NewService(consts.DB_CRM),
}
