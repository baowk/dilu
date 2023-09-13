package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type SysOperaLogService struct {
	*base.BaseService
}

var SerSysOperaLog = SysOperaLogService{
	base.NewService("sys"),
}

