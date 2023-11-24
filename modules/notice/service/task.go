package service

import (
	"github.com/baowk/dilu-core/core/base"
)

type TaskService struct {
	*base.BaseService
}

var SerTask = TaskService{
	base.NewService("notice"),
}

