package service

import (
	"github.com/baowk/dilu-core/common/consts"
	"github.com/baowk/dilu-core/core/base"
)

type SysOperaLogService struct {
	*base.BaseService
}

var SerSysOperaLog = SysOperaLogService{
	base.NewService(consts.DB_DEF),
}
