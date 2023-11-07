package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type SysCfgService struct {
	*base.BaseService
}

var SerSysCfg = SysCfgService{
	base.NewService("sys"),
}
