package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type SysDeptService struct {
	*base.BaseService
}

var SerSysDept = SysDeptService{
	base.NewService("sys"),
}

