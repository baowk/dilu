package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type SysApiService struct {
	base.BaseService
}

var SysApiS = SysApiService{
	base.BaseService{
		DbName: "sys",
	},
}

