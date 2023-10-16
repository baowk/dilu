package service

import (
	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core/base"
)

type SysJobService struct {
	*base.BaseService
}

var SerSysJob = SysJobService{
	base.NewService(consts.DB_DEF),
}
