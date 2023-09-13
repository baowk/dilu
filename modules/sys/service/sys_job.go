package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type SysJobService struct {
	*base.BaseService
}

var SerSysJob = SysJobService{
	base.NewService("sys"),
}

