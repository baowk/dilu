package service

import (
	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core/base"
)

type SysTeamService struct {
	*base.BaseService
}

var SerSysTeam = SysTeamService{
	base.NewService(consts.DB_DEF),
}
