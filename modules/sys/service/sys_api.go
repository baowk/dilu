package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type SysApiService struct {
	*base.BaseService
}

var SerSysApi = SysApiService{
	base.NewService("sys"),
}
