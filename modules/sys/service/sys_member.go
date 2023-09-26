package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type SysMemberService struct {
	*base.BaseService
}

var SerSysMember = SysMemberService{
	base.NewService("sys"),
}
