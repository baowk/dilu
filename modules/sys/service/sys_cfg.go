package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type SysCfgService struct {
	*base.BaseService
}

var SysCfgS = SysCfgService{
	base.NewService("sys"),
}

