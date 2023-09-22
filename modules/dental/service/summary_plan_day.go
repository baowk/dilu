package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type SummaryPlanDayService struct {
	*base.BaseService
}

var SerSummaryPlanDay = SummaryPlanDayService{
	base.NewService("dental"),
}

