package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type EventDayStService struct {
	*base.BaseService
}

var SerEventDaySt = EventDayStService{
	base.NewService("dental"),
}

