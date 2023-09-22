package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type TargetTaskService struct {
	*base.BaseService
}

var SerTargetTask = TargetTaskService{
	base.NewService("dental"),
}

