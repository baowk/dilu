package service

import (
	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core/base"
)

type SysCfgService struct {
	*base.BaseService
}

var SerSysCfg = SysCfgService{
	base.NewService(consts.DB_DEF),
}
