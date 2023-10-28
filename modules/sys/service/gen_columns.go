package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type GenColumnsService struct {
	*base.BaseService
}

var SerGenColumns = GenColumnsService{
	base.NewService("sys"),
}

