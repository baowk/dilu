package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type SysRoleService struct {
	*base.BaseService
}

var SerSysRole = SysRoleService{
	base.NewService("sys"),
}

