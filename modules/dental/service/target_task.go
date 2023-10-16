package service

import (
	"dilu/common/consts"

	"github.com/baowk/dilu-core/core/base"
)

type TargetTaskService struct {
	*base.BaseService
}

var SerTargetTask = TargetTaskService{
	base.NewService(consts.DB_CRM),
}
