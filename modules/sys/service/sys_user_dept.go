package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type SysUserDeptService struct {
	*base.BaseService
}

var SerSysUserDept = SysUserDeptService{
	base.NewService("sys"),
}

