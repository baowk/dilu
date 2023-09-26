package service

import (
	"dilu/modules/sys/models"

	"github.com/baowk/dilu-core/core"
	"github.com/baowk/dilu-core/core/base"
)

type SysDeptService struct {
	*base.BaseService
}

var SerSysDept = SysDeptService{
	base.NewService("sys"),
}

func (s *SysDeptService) GetDepts(list *[]models.SysDept) error {
	return core.DB().Find(list).Error
}
