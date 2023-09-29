package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type SysTeamService struct {
	*base.BaseService
}

var SerSysTeam = SysTeamService{
	base.NewService("sys"),
}

